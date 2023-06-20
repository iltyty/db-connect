package models

import (
	"errors"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type APIUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func CreateUser(dto dtos.UserRegistrationDTO) (APIUser, error) {
	var user User
	err := db.Where("email = ?", dto.Email).First(&user).Error
	if err == nil {
		return APIUser{user.Email, user.Username}, errors.New("user already exists")
	}
	if err != gorm.ErrRecordNotFound {
		return APIUser{}, err
	}

	user.Email = dto.Email
	user.Password = dto.Password
	res := db.Debug().Model(User{}).Create(&user)
	return APIUser{user.Email, user.Username}, res.Error
}

func GetUser(email string) (APIUser, error) {
	var user User
	res := db.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return APIUser{}, res.Error
	}
	return APIUser{user.Email, user.Username}, nil
}

func AuthenticateUser(email, password string) bool {
	var user User
	res := db.Where("email = ?", email).Find(&user)
	if res.Error != nil {
		return false
	}
	return user.Password == password
}

func ResetPassword(dto dtos.UserResetPasswordDTO) error {
	var user User
	res := db.Where("email = ?", dto.Email).Find(&user)
	if res.Error != nil {
		return res.Error
	}
	if user.Password != dto.OldPassword {
		return errors.New("wrong password")
	}
	res = db.Model(&user).Update("password", dto.NewPassword)
	return res.Error
}
