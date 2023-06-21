import axios, { AxiosInstance } from 'axios'
import { AxiosResponse } from 'axios'
import { AxiosRequestConfig } from 'axios'
import { Result } from './types'

export class TsQualityAxios {
  private instance: AxiosInstance
  private config: AxiosRequestConfig

  constructor(config: AxiosRequestConfig) {
    this.config = config
    this.instance = axios.create(config)
    this.setupInterceptors()
  }

  private setupInterceptors() {
    this.instance.interceptors.response.use(
      (res: AxiosResponse) => {
        // TODO: handle according to the data returned by the backend
        return res
      },
      (err: any) => {
        if (err.response) {
          return err.response
        }
        let message = ''
        if (err && err.response) {
          // TODO: handle according to the response status
          switch (err.response.status) {
            default:
              message = 'error'
          }
        }
        return Promise.reject(message)
      }
    )
  }

  public getAxiosInstance(): AxiosInstance {
    return this.instance
  }

  public configAxios(config: AxiosRequestConfig) {
    if (!this.instance) {
      return
    }
    this.config = config
    this.instance = axios.create(config)
  }

  public request(config: AxiosRequestConfig) {
    return new Promise((resolve, reject) => {
      this.instance
        .request<Result, AxiosResponse<Result>>(config)
        .then((res: AxiosResponse<Result>) => {
          resolve(res.data)
        })
        .catch((e: Error) => {
          reject(e)
        })
    })
  }
}
