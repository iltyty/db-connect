package dtos

type UserRegistrationDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
	Code     string `json:"code" form:"code" validate:"required,len=6"`
}

type UserLoginDTO struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=20"`
}

type UserResetPasswordDTO struct {
	Email       string `json:"email" form:"email" validate:"required,email"`
	OldPassword string `json:"old_password" form:"old_password" validate:"required,min=6,max=20"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required,min=6,max=20"`
}

type UserSendEmailCodeDTO struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}
