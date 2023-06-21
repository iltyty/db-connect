import { App } from 'vue'
import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router'

export const RootRoute: RouteRecordRaw = {
  path: '/',
  name: 'Root',
  redirect: '/home',
}

export const LoginRoute: RouteRecordRaw = {
  path: '/login',
  name: 'Login',
  component: () => import('@/pages/login/login.vue'),
}

export const RegisterRoute: RouteRecordRaw = {
  path: '/register',
  name: 'Register',
  component: () => import('@/pages/register/register.vue'),
}

export const ResetRoute: RouteRecordRaw = {
  path: '/reset_password',
  name: 'ResetPassword',
  component: () => import('@/pages/reset/reset.vue'),
}

export const HomeRoute: RouteRecordRaw = {
  path: '/home',
  name: 'Home',
  component: () => import('@/pages/home/home.vue'),
}

const router = createRouter({
  history: createWebHistory(''),
  routes: [RootRoute, LoginRoute, RegisterRoute, ResetRoute, HomeRoute],
})

export function setupRouter(app: App) {
  app.use(router)
}

export default router
