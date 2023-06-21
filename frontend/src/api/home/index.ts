import { useAuthStore } from '@/stores/auth'
import { http } from '@/utils/axios'

const authStore = useAuthStore()

export async function getUserEmail() {
  try {
    return await http.request({
      url: '/user',
      headers: authStore.getTokenHeader(),
    })
  } catch (err) {
    throw err
  }
}

export async function getTestData() {
  try {
    const res = await http.request({
      url: '/test',
      headers: authStore.getTokenHeader(),
    })
    return res
  } catch (err) {
    throw err
  }
}
