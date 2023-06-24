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
          <n-form-item path="email" ref="emailRef">
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
              <n-button
                ghost
                strong
                type="primary"
                size="large"
                :disabled="getBtnDisabled"
                @click="getEmailCode"
              >
                {{ getBtnContent }}
                <div v-show="countdownActive">
                  <n-countdown
                    ref="countdownRef"
                    :duration="60000"
                    :active="countdownActive"
                    :render="countdownRender"
                    :on-finish="countdownFinish"
                  />
                </div>
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
              :loading="signupBtnLoading"
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
import { useMessage, FormRules, CountdownProps } from 'naive-ui'
import { login, register } from '@/api/auth'
import { PersonOutline, LockClosedOutline } from '@vicons/ionicons5'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/auth'
import { Result } from '@/utils/axios/types'
import { computed } from 'vue'
import { getCode } from '@/api/auth'

interface FormParams {
  code: string
  email: string
  password: string
}

const formRef = ref()
const emailRef = ref()
const countdownRef = ref()

const getBtnContent = computed(() => {
  return countdownActive.value ? '' : 'Get'
})
const getBtnDisabled = ref(false)

const countdownActive = ref(false)
const countdownRender: CountdownProps['render'] = ({ minutes, seconds }) => {
  const sec = minutes === 1 ? '60' : seconds
  return `resent after ${sec}s`
}
const countdownFinish = () => {
  countdownRef.value.reset()
  getBtnDisabled.value = false
  countdownActive.value = false
}
const signupBtnLoading = ref(false)

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
    trigger: ['blur', 'input'],
  },
  password: {
    required: true,
    message: 'password length must be between 6 and 20',
    trigger: ['blur', 'input'],
    min: 6,
    max: 20,
  },
  code: {
    required: false,
    message: 'code length must be 6',
    trigger: ['blur', 'input'],
    len: 6,
  },
}

const getEmailCode = async () => {
  try {
    await emailRef.value?.validate()
  } catch (err: any) {
    console.log(err)
    return
  }

  getBtnDisabled.value = true
  countdownActive.value = true

  try {
    const res = (await getCode(formValue)) as Result
    if (res.code != 0) {
      message.error(res.msg)
      return
    }
    message.success('email sent success')
  } catch (err: any) {
    message.error((err as Error).message)
  }
}

const handleSubmit = async (e: any) => {
  e.preventDefault()
  formRef.value.validate(async (errors: any) => {
    if (errors) {
      message.error('form format is problematic, please review')
      return
    }
    const { email, password, code } = formValue
    signupBtnLoading.value = true

    const params: FormParams = {
      email,
      password,
      code,
    }

    try {
      let res = (await register(params)) as Result
      if (res.code != 0) {
        message.error(res.msg)
        return
      }
      message.success('sign up success')

      res = (await login(params)) as Result
      if (res.code != 0) {
        message.error('login failed: ' + res.msg)
        return
      }

      authStore.set(res.data.token)
      router.push({
        path: '/home',
      })
    } catch (err: any) {
      message.error((err as Error).message)
    } finally {
      signupBtnLoading.value = false
    }
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
