package repository

import "github.com/evermos/boilerplate-go/infras"

type Repository interface {
	UserRepository
	ProductRepository
	VariantRepository
}

type RepositoryImpl struct {
	DB *infras.MySQLConn
}

func ProvideRepo(db *infras.MySQLConn) *RepositoryImpl {
	return &RepositoryImpl{
		DB: db,
	}
}
