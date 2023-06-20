package services

import (
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/models"
)

func CreateUser(userDTO dtos.UserRegistrationDTO) (user models.APIUser, err error) {
	user, err = models.CreateUser(userDTO)
	return
}
