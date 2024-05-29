package services

import (
	"go-api/models"
	"go-api/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.FindAllUsers()
}
