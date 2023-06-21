import { http } from '@/utils/axios'
import { UserLoginDTO, UserRegistrationDTO } from '#/dtos'

export async function register(dto: UserRegistrationDTO) {
  try {
    const res = await http.request({
      url: 'users',
      method: 'post',
      data: dto,
    })
    return res
  } catch (err) {
    throw err
  }
}

export async function login(dto: UserLoginDTO) {
  try {
    return await http.request({
      url: 'auth/login',
      method: 'post',
      data: dto,
    })
  } catch (err) {
    throw err
  }
}
