package repository

import (
	"time"

	"github.com/evermos/boilerplate-go/internal/domain/products/models"

	"github.com/rs/zerolog/log"
)

type VariantRepository interface {
	UpdateVariant(variantID int64, variant *models.UpdateVariant) (*models.UpdateVariant, error)
	GetVariantByID(variantID int64) (*models.Variant, error)
}

func (r *RepositoryImpl) UpdateVariant(variantID int64, variant *models.UpdateVariant) (*models.UpdateVariant, error) {
	// Implement the logic to update a variant in the database using r.DB.Write

	query := `
		UPDATE variants
		SET name=?, price=?, stock=?, status=?, updated_at=?, updated_by=?
		WHERE id=?
	`
	variant.ID = variantID
	variant.UpdatedAt = time.Now()

	_, err := r.DB.Write.Exec(
		query,
		variant.Name,
		variant.Price,
		variant.Stock,
		variant.Status,
		variant.UpdatedAt,
		variant.UpdatedBy,
		variant.ID,
	)
	if err != nil {
		return nil, err
	}

	return variant, nil
}

func (r *RepositoryImpl) GetVariantByID(variantID int64) (*models.Variant, error) {
	query := "SELECT * FROM variants WHERE id = ?"

	var variant models.Variant
	err := r.DB.Read.Get(&variant, query, variantID)
	if err != nil {
		log.Error().Err(err).Msg("Something went wrong")
		return nil, err
	}
	return &variant, nil
}
