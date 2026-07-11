<template>
  <div class="flex min-h-screen items-center justify-center bg-gray-50 px-4 dark:bg-dark-950">
    <div class="w-full max-w-md rounded-2xl border border-gray-200 bg-white p-6 text-center shadow-sm dark:border-dark-800 dark:bg-dark-900">
      <div v-if="!errorMessage" class="mx-auto mb-4 h-8 w-8 animate-spin rounded-full border-4 border-primary-500 border-t-transparent"></div>
      <h1 class="text-lg font-semibold text-gray-900 dark:text-white">
        {{ errorMessage ? t('common.error') : 'Opening secure checkout' }}
      </h1>
      <p class="mt-2 text-sm text-gray-500 dark:text-gray-400">
        {{ errorMessage || 'Please wait while we prepare your Simulator billing session.' }}
      </p>
      <button v-if="errorMessage" class="btn btn-primary mt-5 w-full" @click="router.replace('/login')">
        {{ t('auth.login') }}
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRoute, useRouter } from 'vue-router'
import { authAPI } from '@/api/auth'
import { useAuthStore } from '@/stores/auth'
import { extractApiErrorMessage } from '@/utils/apiError'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const { t } = useI18n()
const errorMessage = ref('')

function singleQueryValue(value: unknown): string {
  if (Array.isArray(value)) return typeof value[0] === 'string' ? value[0] : ''
  return typeof value === 'string' ? value : ''
}

onMounted(async () => {
  const token = singleQueryValue(route.query.token)
  if (!token) {
    errorMessage.value = 'Billing session is missing or expired.'
    return
  }

  try {
    const redeemed = await authAPI.redeemBillingSession(token)
    await authStore.setToken(redeemed.access_token)

    const query: Record<string, string> = {
      tab: 'subscription',
    }
    const planId = singleQueryValue(route.query.plan_id) || (redeemed.plan_id ? String(redeemed.plan_id) : '')
    const paymentType = singleQueryValue(route.query.payment_type) || redeemed.payment_type || ''
    const orderType = singleQueryValue(route.query.order_type) || 'subscription'
    const returnTo = singleQueryValue(route.query.return_to) || redeemed.return_url || ''

    if (planId) query.plan_id = planId
    if (paymentType) query.payment_type = paymentType
    if (orderType) query.order_type = orderType
    if (returnTo) query.return_to = returnTo

    await router.replace({ path: '/purchase', query })
  } catch (err) {
    errorMessage.value = extractApiErrorMessage(err, 'Could not open billing session.')
  }
})
</script>
