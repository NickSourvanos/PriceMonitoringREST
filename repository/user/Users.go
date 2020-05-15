package user

import (
	"PriceMonitoringService/models"
)

type UserRepository interface {
	Save(*models.User) error
	Update(string, *models.User) error
	Delete(string) error
	FindByID(string) (*models.ResponseUser, error)
	FindUserByUsername(string) (*models.User, error)
	FindAll() (models.ResponseUsers, error)
}

