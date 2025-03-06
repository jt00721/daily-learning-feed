package repository

import (
	"github.com/jt00721/daily-learning-feed/internal/entity"
	"gorm.io/gorm"
)

type ResourceRepository struct {
	DB *gorm.DB
}

func (repo *ResourceRepository) Create(resource *entity.Resource) error {
	var existing entity.Resource
	err := repo.DB.Where("url = ?", resource.URL).First(&existing).Error

	if err == gorm.ErrRecordNotFound {
		return repo.DB.Create(resource).Error
	}

	if err != nil {
		return err
	}

	return nil
}

func (repo *ResourceRepository) GetAll() ([]entity.Resource, error) {
	var resources []entity.Resource
	err := repo.DB.Find(&resources).Error
	return resources, err
}

func (repo ResourceRepository) GetByID(id uint) (*entity.Resource, error) {
	var resource entity.Resource
	err := repo.DB.First(&resource, id).Error
	return &resource, err
}

func (repo *ResourceRepository) Update(resource *entity.Resource) error {
	return repo.DB.Save(resource).Error
}

func (repo *ResourceRepository) Delete(id uint) error {
	return repo.DB.Delete(&entity.Resource{}, id).Error
}
