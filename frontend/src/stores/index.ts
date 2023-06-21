import { createPinia } from 'pinia'
import { App } from 'vue'
import persist from 'pinia-plugin-persistedstate'

export function setupStore(app: App<Element>) {
  const pinia = createPinia()
  pinia.use(persist)
  app.use(pinia)
}
