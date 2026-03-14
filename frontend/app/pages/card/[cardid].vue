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

const { data, pending, error } = await useFetch<CardResponse>("/cards", {
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
  <div v-if="pending">Loading...</div>

  <div v-else-if="error">
    Error loading
  </div>

  <div v-else-if="data">
    <h1 class="text-3xl font-bold mb-4">{{ data.data.to }}</h1>
    <p>{{ data.data.from }}</p>
    <p class="text-lg mb-2">{{ data.data.description }}</p>
  </div>
</template>