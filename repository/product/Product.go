package product

import "PriceMonitoringService/models"

type ProductRepository interface {
	Save(*models.Product) error
	Update(string, *models.Product) error
	Delete(string) error
	FindByID(string) (*models.Product, error)
	FindProductsByUserId(string) (*models.Product, error)
	FindAll() (models.Products, error)
}
