package repository

import (
	"github.com/jt00721/daily-learning-feed/internal/domain"
	"gorm.io/gorm"
)

type ResourceRepository struct {
	DB *gorm.DB
}

func (repo *ResourceRepository) Create(resource *domain.Resource) error {
	var existing domain.Resource
	err := repo.DB.Where("url = ?", resource.URL).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return repo.DB.Create(resource).Error
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *ResourceRepository) GetAll() ([]domain.Resource, error) {
	var resources []domain.Resource
	err := repo.DB.Find(&resources).Error
	return resources, err
}

func (repo ResourceRepository) GetByID(id uint) (*domain.Resource, error) {
	var resource domain.Resource
	err := repo.DB.First(&resource, id).Error
	return &resource, err
}

func (repo *ResourceRepository) Update(resource *domain.Resource) error {
	return repo.DB.Save(resource).Error
}

func (repo *ResourceRepository) Delete(id uint) error {
	return repo.DB.Delete(&domain.Resource{}, id).Error
}
