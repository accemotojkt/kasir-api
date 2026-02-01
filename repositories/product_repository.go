package repositories

import (
	"errors"
	"kasir-api/models"
)

type ProductRepository struct {
	products []models.Product
}

func NewProductRepository(products []models.Product) *ProductRepository {
	return &ProductRepository{products: products}
}

func (repo *ProductRepository) GetAll() ([]models.Product, error) {
	return repo.products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	// masukkin data ke dalam variable products di repository ini
	product.ID = len(repo.products) + 1
	repo.products = append(repo.products, *product)

	return nil
}

func (repo *ProductRepository) GetByID(id int) (*models.Product, error) {
	for _, p := range repo.products {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("Product not found")
}

func (repo *ProductRepository) Update(product *models.Product) error {
	// loop products, cari id, ganti sesuai data dari request
	for i := range repo.products {
		if repo.products[i].ID == product.ID {
			repo.products[i] = *product

			return nil
		}
	}

	return errors.New("Product not found")
}

func (repo *ProductRepository) Delete(id int) error {
	// loop products cari ID, dapet index yang mau dihapus
	for i, p := range repo.products {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			repo.products = append(repo.products[:i], repo.products[i+1:]...)

			return nil
		}
	}

	return errors.New("Product not found")
}
