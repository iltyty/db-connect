import { http } from '@/utils/axios'
import { AuthResetPasswordDTO, UserLoginDTO, UserRegistrationDTO } from '#/dtos'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

export async function register(dto: UserRegistrationDTO) {
  try {
    const res = await http.request({
      url: 'user',
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

export function logout() {
  authStore.set('')
}

export async function checkAuth() {
  if (authStore.jwt == '') {
    return false
  }
  try {
    return await http
      .request({
        url: '/auth/check',
        method: 'get',
        headers: authStore.getTokenHeader(),
      })
      .then((res: any) => {
        return res.code === 0
      })
      .catch(() => {
        return false
      })
  } catch (err) {
    return false
  }
}

export async function resetPassword(dto: AuthResetPasswordDTO) {
  try {
    return await http.request({
      url: 'auth/reset_password',
      method: 'post',
      data: {
        email: dto.email,
        old_password: dto.oldPassword,
        new_password: dto.newPassword,
      },
    })
  } catch (err) {
    throw err
  }
}
