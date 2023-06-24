package services

import (
	"context"
	"fmt"
	"github.com/iltyty/db_connect/backend_go/conf"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"github.com/iltyty/db_connect/backend_go/utils"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

func RandomCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", r.Int31n(1000000))
}

func SendEmailCode(dto dtos.UserSendEmailCodeDTO) (code string, err error) {
	code = RandomCode()
	err = utils.RDB.Set(context.Background(), dto.Email+"_code", code, time.Minute*5).Err()
	utils.CheckError(err)

	email := gomail.NewMessage()
	email.SetHeader("From", conf.EmailConf.Username)
	email.SetHeader("To", dto.Email)
	email.SetHeader("Subject", "Verification code")
	email.SetBody("text/html", "Your verification code is "+code)

	dialer := gomail.NewDialer(
		"smtp.163.com", 25, conf.EmailConf.Username, conf.EmailConf.Password,
	)
	if err = dialer.DialAndSend(email); err != nil {
		return "", err
	}
	return
}

func CheckEmailCodeValid(dto dtos.UserRegistrationDTO) bool {
	code, err := utils.RDB.Get(context.Background(), dto.Email+"_code").Result()
	return err == nil && code == dto.Code
}
