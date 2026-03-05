<script setup lang="ts">
useHead({
  title: "Create New Card",
  meta: [
    {
      name: "description",
      content: "Eid Mubarak",
    },
  ],
});

import { useToast } from "vue-toastification";
import QRCode from "qrcode";

// variables
const config = useRuntimeConfig();
const toast = useToast();
const form = reactive({
  to: "",
  from: "",
  title: "",
  description: "",
  type: "send",
});

// qr code
const cardId = ref("");
const showQR = ref(false);
const qrCodeUrl = ref("");
function resetForm() {
  form.to = "";
  form.from = "";
  form.title = "";
  form.description = "";
  form.type = "send";

  cardId.value = "";
  success.value = false;
  error.value = "";
  qrCodeUrl.value = "";
}

function closeQR() {
  showQR.value = false;
  resetForm();
}
// qr code

const loading = ref(false);
const error = ref("");
const success = ref(false);

//interfaces
interface CreateCardResponse {
  card_id: number;
  message?: string;
}

async function submitForm() {
  loading.value = true;
  error.value = "";

  try {
    const response: CreateCardResponse = await $fetch("/cards", {
      baseURL: config.public.apiBase,
      method: "POST",
      headers: {
        "Application-Key": config.public.applicationKey,
        "Content-Type": "application/json",
      },
      body: form,
    });

    // If your API returns: { card_id: 123 }
    console.log(response.card_id);

    cardId.value = String(response.card_id);

    // generate QR (example: link to card page)
    const qrData = `${window.location.origin}/card/${cardId.value}`;
    qrCodeUrl.value = await QRCode.toDataURL(qrData);

    showQR.value = true;
    success.value = true;

    toast.success("Card created successfully");
  } catch (err: any) {
    error.value = err?.data?.message || "Failed to create card";
    toast.error(error.value);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 px-6 py-8">
    <div class="max-w-md mx-auto">
      <h1 class="text-2xl font-bold text-center mb-6">Create Eid Card</h1>

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
          {{ loading ? "Creating..." : "Create Card" }}
        </button>

        <p v-if="error" class="text-red-500 text-sm text-center">
          {{ error }}
        </p>

        <p v-if="success" class="text-green-600 text-sm text-center">
          Card created successfully
        </p>
      </form>

      <!-- QR Modal -->
      <div
        v-if="showQR"
        class="fixed inset-0 bg-black/50 flex items-center justify-center z-50"
      >
        <div class="bg-white p-6 rounded-2xl w-80 text-center space-y-4">
          <h2 class="text-lg font-semibold">Scan QR Code</h2>

          <img :src="qrCodeUrl" alt="QR Code" class="mx-auto w-48 h-48" />

          <button
            @click="closeQR"
            class="w-full bg-green-600 text-white py-2 rounded-xl"
          >
            Close
          </button>
        </div>
      </div>

      <p class="text-center mt-4">
        <NuxtLink to="/scan" class="text-green-600 hover:underline">
          Or scan QR Code
        </NuxtLink>
      </p>
    </div>
  </div>
</template>
