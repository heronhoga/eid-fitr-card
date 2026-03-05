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

const loading = ref(false);
const error = ref("");
const success = ref(false);

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

function downloadQR() {
  if (!qrCodeUrl.value) return;

  const link = document.createElement("a");
  link.href = qrCodeUrl.value;
  link.download = `eid-card-${cardId.value}.png`;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
}

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

    cardId.value = String(response.card_id);

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
  <div
    class="min-h-screen bg-gradient-to-b from-green-100 via-white to-green-50 flex flex-col"
  >
    <!-- Header -->
    <header class="text-center pt-10 pb-4">
      <h1 class="text-3xl font-bold text-green-700">✨ Create Eid Card</h1>
      <p class="text-gray-500 text-sm mt-1">Send a digital Eid greeting</p>
    </header>

    <!-- Form Container -->
    <main class="flex-1 flex items-center justify-center px-6 pb-12">
      <div class="bg-white w-full max-w-md rounded-3xl shadow-xl p-8">
        <form @submit.prevent="submitForm" class="space-y-5">
          <div>
            <label class="text-sm text-gray-600">To</label>
            <input
              v-model="form.to"
              type="text"
              placeholder="Recipient name"
              class="mt-1 w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
              required
            />
          </div>

          <div>
            <label class="text-sm text-gray-600">From</label>
            <input
              v-model="form.from"
              type="text"
              placeholder="Your name"
              class="mt-1 w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
              required
            />
          </div>

          <div>
            <label class="text-sm text-gray-600">Card Title</label>
            <input
              v-model="form.title"
              type="text"
              placeholder="Eid Mubarak"
              class="mt-1 w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
              required
            />
          </div>

          <div>
            <label class="text-sm text-gray-600">Message</label>
            <textarea
              v-model="form.description"
              rows="4"
              placeholder="Write your Eid wishes..."
              class="mt-1 w-full p-3 rounded-xl border border-gray-300 focus:ring-2 focus:ring-green-500"
              required
            ></textarea>
          </div>

          <button
            type="submit"
            class="w-full bg-green-600 hover:bg-green-700 text-white py-3 rounded-2xl font-semibold shadow-md transition active:scale-95"
            :disabled="loading"
          >
            {{ loading ? "Creating Card..." : "Create Card" }}
          </button>

          <p v-if="error" class="text-red-500 text-sm text-center">
            {{ error }}
          </p>

          <p v-if="success" class="text-green-600 text-sm text-center">
            Card created successfully
          </p>
        </form>

        <p class="text-center mt-6 text-sm text-gray-500">
          or
          <NuxtLink
            to="/scan"
            class="text-green-600 font-medium hover:underline"
          >
            scan a QR card
          </NuxtLink>
        </p>
      </div>
    </main>

    <!-- QR Modal -->
    <div
      v-if="showQR"
      class="fixed inset-0 bg-black/60 flex items-center justify-center z-50 px-4"
    >
      <div
        class="bg-white rounded-3xl p-6 w-full max-w-sm text-center shadow-xl space-y-4"
      >
        <h2 class="text-xl font-semibold text-green-700">📷 Scan Your Card</h2>

        <img :src="qrCodeUrl" alt="QR Code" class="mx-auto w-48 h-48" />

        <div class="space-y-2">
          <button
            @click="downloadQR"
            class="w-full bg-white border border-green-600 text-green-600 py-2 rounded-xl font-medium hover:bg-green-50"
          >
            Download QR Code
          </button>

          <button
            @click="closeQR"
            class="w-full bg-green-600 text-white py-2 rounded-xl"
          >
            Close
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
