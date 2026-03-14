<script setup lang="ts">
const route = useRoute();
const config = useRuntimeConfig();

interface CardData {
  card_id: string;
  to: string;
  from: string;
  description: string;
}

interface CardResponse {
  data: CardData;
  message: string;
}

const { data, pending, error } = await useFetch<CardResponse>("/cards/get", {
  baseURL: config.public.apiBase,
  method: "POST",
  headers: {
    "Application-Key": config.public.applicationKey,
    "Content-Type": "application/json",
  },
  body: {
    card_id: route.params.cardid as string,
  },
});
</script>

<template>
  <div
    class="min-h-screen bg-gradient-to-b from-green-100 via-white to-green-50 flex items-center justify-center px-4 py-10"
  >
    <!-- Loading -->
    <div v-if="pending" class="bg-white p-8 rounded-3xl shadow-xl text-center">
      <div
        class="animate-spin w-10 h-10 border-4 border-green-500 border-t-transparent rounded-full mx-auto mb-4"
      ></div>
      <p class="text-gray-600">Loading your Eid card...</p>
    </div>

    <!-- Error -->
    <div
      v-else-if="error"
      class="bg-white p-8 rounded-3xl shadow-xl text-center max-w-md w-full"
    >
      <h2 class="text-xl font-semibold text-red-500 mb-2">Card Not Found</h2>
      <p class="text-gray-500">
        The Eid card you are trying to view could not be loaded.
      </p>
    </div>

    <!-- Card -->
    <div
      v-else-if="data"
      class="bg-white max-w-md w-full rounded-3xl shadow-xl p-8 text-center space-y-6"
    >
      <!-- To -->
      <div class="space-y-1">
        <p class="text-sm text-gray-500">Eid Greeting For</p>
        <h1
          class="text-2xl sm:text-3xl font-bold text-green-700 break-words"
        >
          {{ data.data.to }}
        </h1>
      </div>

      <!-- Message -->
      <div
        class="bg-green-50 rounded-2xl p-5 max-h-64 overflow-y-auto"
      >
        <p class="text-lg font-medium text-gray-700 mb-2">
          🌙 Eid Mubarak
        </p>

        <p
          class="text-gray-600 leading-relaxed break-words text-left"
        >
          {{ data.data.description }}
        </p>
      </div>

      <!-- From -->
      <div class="pt-2">
        <p class="text-sm text-gray-500">From</p>
        <p
          class="text-lg font-semibold text-gray-700 break-words"
        >
          {{ data.data.from }}
        </p>
      </div>

      <NuxtLink
        to="/"
        class="block mt-4 text-green-600 font-medium hover:underline"
      >
        Create your own card →
      </NuxtLink>
    </div>
  </div>
</template>
