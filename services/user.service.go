package services

import "github.com/alonelegion/api_golang_mongodb/models"

type UserInterface interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
}
