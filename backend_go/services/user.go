package services

import (
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/models"
)

func CreateUser(dto dtos.UserRegistrationDTO) (user models.APIUser, err error) {
	user, err = models.CreateUser(dto)
	return
}

func ResetPassword(dto dtos.UserResetPasswordDTO) (err error) {
	err = models.ResetPassword(dto)
	return
}
