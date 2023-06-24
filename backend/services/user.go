package services

import (
	"errors"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/models"
)

func CreateUser(dto dtos.UserRegistrationDTO) (user models.APIUser, err error) {
	if !CheckEmailCodeValid(dto) {
		return user, errors.New("invalid verification code")
	}
	user, err = models.CreateUser(dto)
	return
}

func ResetPassword(dto dtos.UserResetPasswordDTO) (err error) {
	err = models.ResetPassword(dto)
	return
}
