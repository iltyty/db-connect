package services

import (
	"fmt"
	"github.com/iltyty/db_connect/backend_go/dtos"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
)

const (
	emailUsername = "qiuyh0924@163.com"
	emailPassword = "XQYJLVSHVZGNVJWS"
)

func RandomCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%06v", r.Int31n(1000000))
}

func SendEmailCode(dto dtos.UserSendEmailCodeDTO) (code string, err error) {
	code = RandomCode()
	email := gomail.NewMessage()
	email.SetHeader("From", emailUsername)
	email.SetHeader("To", dto.Email)
	email.SetHeader("Subject", "Verification code")
	email.SetBody("text/html", "Your verification code is "+code)

	dialer := gomail.NewDialer("smtp.163.com", 25, emailUsername, emailPassword)
	if err = dialer.DialAndSend(email); err != nil {
		return "", err
	}
	return
}
