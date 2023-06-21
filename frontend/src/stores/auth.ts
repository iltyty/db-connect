import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => {
    return {
      jwt: '',
    }
  },

  actions: {
    set(data: string) {
      this.jwt = data
    },
    getTokenHeader() {
      return {
        Authorization: 'token ' + this.jwt,
      }
    },
  },

  persist: true,
})
