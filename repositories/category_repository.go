package repositories

import (
	"errors"
	"kasir-api/models"
)

type CategoryRepository struct {
	categories []models.Category
}

func NewCategoryRepository(categories []models.Category) *CategoryRepository {
	return &CategoryRepository{categories: categories}
}

func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	return repo.categories, nil
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	// masukkin data ke dalam variable categories di repository ini
	category.ID = len(repo.categories) + 1
	repo.categories = append(repo.categories, *category)

	return nil
}

func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	for _, c := range repo.categories {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, errors.New("Category not found")
}

func (repo *CategoryRepository) Update(category *models.Category) error {
	// loop categories, cari id, ganti sesuai data dari request
	for i := range repo.categories {
		if repo.categories[i].ID == category.ID {
			repo.categories[i] = *category

			return nil
		}
	}

	return errors.New("Category not found")
}

func (repo *CategoryRepository) Delete(id int) error {
	// loop categories cari ID, dapet index yang mau dihapus
	for i, c := range repo.categories {
		if c.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			repo.categories = append(repo.categories[:i], repo.categories[i+1:]...)

			return nil
		}
	}

	return errors.New("Category not found")
}
