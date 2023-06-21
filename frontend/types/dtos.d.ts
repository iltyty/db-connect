export interface UserRegistrationDTO {
  code: string
  email: string
  password: string
}

export interface UserLoginDTO {
  email: string
  password: string
}

export interface AuthResetPasswordDTO {
  email: string
  oldPassword: string
  newPassword: string
}
