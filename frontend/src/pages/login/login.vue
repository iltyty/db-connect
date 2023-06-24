<template>
  <div class="login">
    <div class="login-container">
      <div class="login-container-title">
        <h1>DB Connect</h1>
      </div>
      <div class="login-form">
        <n-form
          ref="formRef"
          size="large"
          label-placement="left"
          :rules="rules"
          :model="formValue"
        >
          <n-form-item path="email">
            <n-input v-model:value="formValue.email" placeholder="email">
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <PersonOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>

          <n-form-item path="password">
            <n-input
              v-model:value="formValue.password"
              type="password"
              showPasswordOn="click"
              placeholder="password"
            >
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <LockClosedOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>

          <n-form-item>
            <n-button
              type="primary"
              @click="handleSubmit"
              size="large"
              :loading="loading"
              block
            >
              Sign in
            </n-button>
          </n-form-item>

          <n-form-item>
            <!-- <n-checkbox>Remember me</n-checkbox> -->
            <div style="margin-right: auto">
              <router-link to="reset_password">Forgot password?</router-link>
            </div>

            <div style="margin-left: auto">
              <router-link to="register">Sign up</router-link>
            </div>
          </n-form-item>
        </n-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { PersonOutline, LockClosedOutline } from '@vicons/ionicons5'
import { FormRules, useMessage } from 'naive-ui'
import { login } from '@/api/auth'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

interface FormParams {
  email: string
  password: string
}

const formRef = ref()
const loading = ref(false)
const message = useMessage()
const formValue = reactive({
  email: '',
  password: '',
})

const router = useRouter()

const rules: FormRules = {
  email: {
    required: true,
    type: 'email',
    message: 'invalid email',
    trigger: ['blur', 'input'],
  },
  password: {
    required: true,
    message: 'password length must be between 6 and 20',
    trigger: ['blur', 'input'],
    min: 6,
    max: 20,
  },
}

const authStore = useAuthStore()

const handleSubmit = (e: any) => {
  e.preventDefault()
  formRef.value.validate(async (errors: any) => {
    if (errors) {
      message.error('form format is problematic, please review')
      return
    }
    const { email, password } = formValue
    loading.value = true

    const params: FormParams = {
      email,
      password,
    }

    login(params)
      .then((res: any) => {
        if (res.code != 0) {
          message.error(res.msg)
          return
        }
        message.success('login success')
        authStore.set(res.data.token)
        router.push({
          path: '/home',
        })
      })
      .catch((err: Error) => {
        message.error(err.message)
      })
      .finally(() => {
        loading.value = false
      })
  })
}
</script>

<style lang="less" scoped>
.login {
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: auto;
  // justify-content: center;
  align-items: center;

  &-container {
    // flex: 1;
    padding: 32px 12px;
    max-width: 384px;
    min-width: 320px;
    margin-left: auto;
    margin-right: auto;
    margin-top: 20vh;

    &-title {
      text-align: center;
    }
  }
}

@media (min-width: 768px) {
  .login {
    background-image: url('../../assets/login.svg');
    background-repeat: no-repeat;
    background-position: 50%;
    background-size: 100%;
  }
}
</style>
