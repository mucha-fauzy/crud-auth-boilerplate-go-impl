package models

import "time"

type ProductView struct {
	ID          int64     `db:"id"`
	UserName    string    `db:"user_name"`
	BrandName   string    `db:"brand_name"`
	VariantName string    `db:"variant_name"`
	ImageURL    string    `db:"image_url"`
	Price       float64   `db:"price"`
	Stock       int64     `db:"stock"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedBy   string    `db:"updated_by"`
}

type ProductFilter struct {
	BrandName   string `json:"brand_name"`
	ProductName string `json:"product_name"`
	VariantName string `json:"variant_name"`
	Status      string `json:"status"`
}

type User struct {
	ID        int64      `db:"id"`
	UserName  string     `db:"username"`
	UserType  string     `db:"user_type"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt time.Time  `db:"updated_at"`
	UpdatedBy string     `db:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at"` // using * to handle NULL value as it refer to memory address
	DeletedBy *string    `db:"deleted_by"`
}

type Product struct {
	ID          int64      `db:"id"`
	UserID      int64      `db:"user_id"`
	BrandID     int64      `db:"brand_id"`
	WarehouseID int64      `db:"warehouse_id"`
	Name        string     `db:"name"`
	CreatedAt   time.Time  `db:"created_at"`
	CreatedBy   string     `db:"created_by"`
	UpdatedAt   time.Time  `db:"updated_at"`
	UpdatedBy   string     `db:"updated_by"`
	DeletedAt   *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy   *string    `db:"deleted_by" json:"deleted_by"`
}

type Brand struct {
	ID        int64      `db:"id"`
	VariantID int64      `db:"variant_id"`
	ImageID   int64      `db:"image_id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt time.Time  `db:"updated_at"`
	UpdatedBy string     `db:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at"`
	DeletedBy *string    `db:"deleted_by"`
}

type Variant struct {
	ID        int64      `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Price     float64    `db:"price" json:"price"`
	Stock     int        `db:"stock" json:"stock"`
	Status    string     `db:"status" json:"status"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	CreatedBy string     `db:"created_by" json:"created_by"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	UpdatedBy string     `db:"updated_by" json:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
	DeletedBy *string    `db:"deleted_by" json:"deleted_by"`
}

type UpdateVariant struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Price     float64   `db:"price" json:"price"`
	Stock     int       `db:"stock" json:"stock"`
	Status    string    `db:"status" json:"status"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
}

type Warehouse struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	City      string     `db:"city"`
	Province  string     `db:"province"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt time.Time  `db:"updated_at"`
	UpdatedBy string     `db:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at"`
	DeletedBy *string    `db:"deleted_by"`
}

type Image struct {
	ID        int64      `db:"id"`
	ImageURL  string     `db:"image_url"`
	CreatedAt time.Time  `db:"created_at"`
	CreatedBy string     `db:"created_by"`
	UpdatedAt time.Time  `db:"updated_at"`
	UpdatedBy string     `db:"updated_by"`
	DeletedAt *time.Time `db:"deleted_at"`
	DeletedBy *string    `db:"deleted_by"`
}
