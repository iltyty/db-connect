import { AxiosRequestConfig } from 'axios'
import { TsQualityAxios } from './axios'

function createAxios(config?: AxiosRequestConfig) {
  return new TsQualityAxios({
    timeout: 10 * 1000,
    baseURL: 'http://127.0.0.1:8000/api/v1/',
  })
}

export const http = createAxios()
