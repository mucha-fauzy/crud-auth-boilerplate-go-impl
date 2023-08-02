package repository

import (
	"time"

	"github.com/evermos/boilerplate-go/internal/domain/products/models"

	"github.com/rs/zerolog/log"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) (*models.Product, error) // using * as it create directly to memory address
	ListProducts(filter models.ProductFilter, sortBy string, page, size int) ([]models.ProductView, error)
	GetProductByID(productID int64) (*models.Product, error)
	SoftDeleteProduct(productID int64, deletedBy string) error
	HardDeleteProduct(productID int64) error
}

func (r *RepositoryImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	// Implement the logic to create a product in the database using r.DB.Write
	query := `
		INSERT INTO products (user_id, brand_id, warehouse_id, name, created_at, created_by, updated_at, updated_by)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	result, err := r.DB.Write.Exec(
		query,
		product.UserID,
		product.BrandID,
		product.WarehouseID,
		product.Name,
		product.CreatedAt,
		product.CreatedBy,
		product.UpdatedAt,
		product.UpdatedBy,
	)
	if err != nil {
		return nil, err
	}

	productID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	product.ID = productID
	return product, nil
}

func (r *RepositoryImpl) ListProducts(filter models.ProductFilter, sortBy string, page, size int) ([]models.ProductView, error) {
	// Implement the logic to fetch the list of products from the database using r.DB.Read
	query := `
		SELECT p.id, u.username AS user_name, b.name AS brand_name, v.name AS variant_name, i.image_url, v.price, v.stock, v.status, p.created_at, p.updated_by
		FROM products p
		JOIN users u ON p.user_id = u.id
		JOIN brands b ON p.brand_id = b.id
		JOIN variants v ON b.variant_id = v.id
		JOIN images i ON b.image_id = i.id
		WHERE p.deleted_at IS NULL
	`

	// Add filters
	args := []interface{}{}
	if filter.BrandName != "" {
		query += " AND b.name LIKE ?"
		args = append(args, "%"+filter.BrandName+"%")
	}

	if filter.ProductName != "" {
		query += " AND p.name LIKE ?"
		args = append(args, "%"+filter.ProductName+"%")
	}

	if filter.VariantName != "" {
		query += " AND v.name LIKE ?"
		args = append(args, "%"+filter.VariantName+"%")
	}

	if filter.Status != "" {
		query += " AND v.status LIKE ?"
		args = append(args, "%"+filter.Status+"%")
	}

	// Add sorting
	if sortBy != "" {
		query += " ORDER BY " + sortBy
	}

	// Add pagination
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}
	query += " LIMIT ? OFFSET ?"
	offset := (page - 1) * size
	args = append(args, size, offset)

	var products []models.ProductView
	err := r.DB.Read.Select(&products, query, args...)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *RepositoryImpl) SoftDeleteProduct(productID int64, deletedBy string) error {
	query := `
		UPDATE products 
		SET deleted_at=?, deleted_by=?
		WHERE id=?
	`
	_, err := r.DB.Write.Exec(
		query,
		time.Now(),
		deletedBy,
		productID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) HardDeleteProduct(productID int64) error {
	query := `
		DELETE FROM products 
		WHERE id=?
	`
	_, err := r.DB.Write.Exec(query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (r *RepositoryImpl) GetProductByID(productID int64) (*models.Product, error) {
	query := "SELECT * FROM products WHERE id = ?"

	var product models.Product
	err := r.DB.Read.Get(&product, query, productID)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong")
		return nil, err
	}
	return &product, nil
}
