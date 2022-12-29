package services

import "github.com/alonelegion/api_golang_mongodb/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
	UpdateUserById(id, field, value string) (*models.DBResponse, error)
	UpdateOne(field string, value interface{}) (*models.DBResponse, error)
}
