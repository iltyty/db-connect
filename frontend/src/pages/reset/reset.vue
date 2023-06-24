<template>
  <div class="reset">
    <div class="reset-container">
      <div class="reset-container-title">
        <h1>DB Connect</h1>
      </div>
      <div class="reset-form">
        <n-form
          ref="formRef"
          size="large"
          label-placement="left"
          :model="formValue"
          :rules="rules"
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

          <n-form-item path="oldPassword">
            <n-input
              v-model:value="formValue.oldPassword"
              type="password"
              showPasswordOn="click"
              placeholder="old password"
            >
              <template #prefix>
                <n-icon size="18" color="#808695">
                  <LockClosedOutline />
                </n-icon>
              </template>
            </n-input>
          </n-form-item>

          <n-form-item path="newPassword">
            <n-input
              v-model:value="formValue.newPassword"
              type="password"
              showPasswordOn="click"
              placeholder="new password"
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
              Confirm
            </n-button>
          </n-form-item>

          <n-form-item>
            <div style="margin-right: auto">
              <router-link to="/reset_password"></router-link>
            </div>

            <div style="margin-left: auto">
              <router-link to="/register"></router-link>
            </div>
          </n-form-item>
        </n-form>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { FormRules, useMessage } from 'naive-ui'
import { PersonOutline, LockClosedOutline } from '@vicons/ionicons5'
import { resetPassword } from '@/api/auth'
import { useRouter } from 'vue-router'

const formRef = ref()
const loading = ref(false)
const message = useMessage()
const formValue = reactive({
  email: '',
  oldPassword: '',
  newPassword: '',
})

const rules: FormRules = {
  email: {
    required: true,
    type: 'email',
    message: 'invalid email',
    trigger: ['blur', 'input'],
  },
  oldPassword: {
    required: true,
    message: 'password length must be between 6 and 20',
    trigger: ['blur', 'input'],
    min: 6,
    max: 20,
  },
  newPassword: {
    required: true,
    message: 'password length must be between 6 and 20',
    trigger: ['blur', 'input'],
    min: 6,
    max: 20,
  },
}

const router = useRouter()
const handleSubmit = (e: any) => {
  e.preventDefault()
  formRef.value.validate(async (errors: any) => {
    if (errors) {
      message.error('form format is problematic, please review')
      return
    }

    resetPassword(formValue)
      .then((res: any) => {
        if (res.code != 0) {
          message.error(res.msg)
          return
        }
        message.success('password reset success')
        router.push({
          path: '/login',
        })
      })
      .catch((err: Error) => {
        message.error(err.message)
      })
  })
}
</script>

<style lang="less" scoped>
.reset {
  display: flex;
  flex-direction: column;
  height: 100vh;
  overflow: auto;
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
  .reset {
    background-image: url('../../assets/login.svg');
    background-repeat: no-repeat;
    background-position: 50%;
    background-size: 100%;
  }
}
</style>
