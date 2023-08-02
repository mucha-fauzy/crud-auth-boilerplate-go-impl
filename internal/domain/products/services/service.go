package services

import (
	"github.com/evermos/boilerplate-go/internal/domain/products/models"

	"github.com/evermos/boilerplate-go/internal/domain/products/repository"
)

type Service interface {
	ReadUser() ([]models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateProduct(product *models.Product) (*models.Product, error)
	ListProducts(filter models.ProductFilter, sortBy string, page, size int) ([]models.ProductView, error)
	GetProductByID(productID int64) (*models.Product, error)
	UpdateVariant(variantID int64, variant *models.UpdateVariant) (*models.UpdateVariant, error)
	GetVariantByID(variantID int64) (*models.Variant, error)
	SoftDeleteProduct(productID int64, deletedBy string) error
	HardDeleteProduct(productID int64) error
}

type ServiceImpl struct {
	Repo repository.Repository
}

func ProvideService(r repository.Repository) *ServiceImpl {
	return &ServiceImpl{
		Repo: r,
	}
}

func (s *ServiceImpl) ReadUser() ([]models.User, error) {
	return s.Repo.ReadUser()
}

func (s *ServiceImpl) GetUserByUsername(username string) (*models.User, error) {
	return s.Repo.GetUserByUsername(username)
}

func (s *ServiceImpl) CreateProduct(product *models.Product) (*models.Product, error) {
	return s.Repo.CreateProduct(product)
}

func (s *ServiceImpl) ListProducts(filter models.ProductFilter, sortBy string, page, size int) ([]models.ProductView, error) {
	return s.Repo.ListProducts(filter, sortBy, page, size)
}

func (s *ServiceImpl) GetProductByID(productID int64) (*models.Product, error) {
	return s.Repo.GetProductByID(productID)
}

func (s *ServiceImpl) UpdateVariant(variantID int64, variant *models.UpdateVariant) (*models.UpdateVariant, error) {
	return s.Repo.UpdateVariant(variantID, variant)
}

func (s *ServiceImpl) GetVariantByID(variantID int64) (*models.Variant, error) {
	return s.Repo.GetVariantByID(variantID)
}

func (s *ServiceImpl) SoftDeleteProduct(productID int64, deletedBy string) error {
	return s.Repo.SoftDeleteProduct(productID, deletedBy)
}

func (s *ServiceImpl) HardDeleteProduct(productID int64) error {
	return s.Repo.HardDeleteProduct(productID)
}
