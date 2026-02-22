<script setup lang="ts">
const config = useRuntimeConfig()

const form = reactive({
  to: '',
  from: '',
  title: '',
  description: ''
})

const loading = ref(false)
const error = ref('')
const success = ref(false)

async function submitForm() {
  loading.value = true
  error.value = ''

  try {
    await $fetch('/cards', {
      baseURL: config.public.apiBase,
      method: 'POST',
      body: form
    })

    success.value = true
  } catch (err: any) {
    error.value = err?.data?.message || 'Failed to create card'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 px-6 py-8">
    <div class="max-w-md mx-auto">

      <h1 class="text-2xl font-bold text-center mb-6">
        Create Eid Card
      </h1>

      <form @submit.prevent="submitForm" class="space-y-4">

        <input
          v-model="form.to"
          type="text"
          placeholder="To"
          class="w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
          required
        />

        <input
          v-model="form.from"
          type="text"
          placeholder="From"
          class="w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
          required
        />

        <input
          v-model="form.title"
          type="text"
          placeholder="Card Title"
          class="w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
          required
        />

        <textarea
          v-model="form.description"
          placeholder="Write your message..."
          rows="4"
          class="w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
          required
        ></textarea>

        <button
          type="submit"
          class="w-full bg-green-600 text-white py-3 rounded-2xl font-semibold active:scale-95 transition"
          :disabled="loading"
        >
          {{ loading ? 'Creating...' : 'Create Card' }}
        </button>

        <p v-if="error" class="text-red-500 text-sm text-center">
          {{ error }}
        </p>

        <p v-if="success" class="text-green-600 text-sm text-center">
          Card created successfully
        </p>

      </form>

    </div>
  </div>
</template>