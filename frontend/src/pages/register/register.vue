<template>
  <div class="register">
    <div class="register-container">
      <div class="register-container-title">
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

          <n-form-item path="code">
            <n-input-group>
              <n-input
                v-model:value="formValue.code"
                placeholder="validation code"
              >
              </n-input>
              <n-button ghost strong type="primary" size="large">
                Get
              </n-button>
            </n-input-group>
          </n-form-item>

          <n-form-item>
            <n-button
              block
              strong
              type="primary"
              @click="handleSubmit"
              size="large"
              :loading="loading"
            >
              Sign up
            </n-button>
          </n-form-item>
        </n-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { useMessage, FormRules } from 'naive-ui'
import { login, register } from '@/api/auth'
import { PersonOutline, LockClosedOutline } from '@vicons/ionicons5'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

interface FormParams {
  code: string
  email: string
  password: string
}

const formRef = ref()
const loading = ref<boolean>(false)
const message = useMessage()
const formValue = reactive({
  email: '',
  password: '',
  code: '',
})

const router = useRouter()
const authStore = useAuthStore()

const rules: FormRules = {
  email: {
    required: true,
    type: 'email',
    message: 'invalid email',
    trigger: 'blur',
  },
  password: {
    required: true,
    message: 'password length must be between 6 and 20',
    trigger: 'blur',
    min: 6,
    max: 20,
  },
  code: {
    required: false,
    message: 'code length must be 6',
    trigger: 'blur',
    // len: 6,
  },
}

const handleSubmit = (e: any) => {
  e.preventDefault()
  formRef.value.validate(async (errors: any) => {
    if (errors) {
      message.error('Form format is problematic, please review')
      return
    }
    const { email, password, code } = formValue
    loading.value = true

    const params: FormParams = {
      email,
      password,
      code,
    }

    register(params)
      .then((res: any) => {
        if (res.code != 0) {
          message.error(res.msg)
          return
        }
        message.success('sign up success')
        login(params)
          .then((res: any) => {
            if (res.code != 0) {
              message.error('login failed: ', res.msg)
              return
            }
            authStore.set(res.data.token)
            router.push({
              path: '/home',
            })
          })
          .catch((err: any) => {
            message.error(err)
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
.register {
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: auto;
  align-items: center;

  &-container {
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

  &-form {
    &-code {
      display: flex;
      justify-content: space-between;
    }
  }
}

@media (min-width: 768px) {
  .register {
    background-image: url('../../assets/login.svg');
    background-repeat: no-repeat;
    background-position: 50%;
    background-size: 100%;
  }
}
</style>
