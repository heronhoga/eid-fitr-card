<script setup>
useHead({
  title: "Scan QR Code",
  meta: [
    {
      name: "description",
      content: "Eid Mubarak",
    },
  ],
});

import QrScanner from "qr-scanner";
import { ref, onMounted, onBeforeUnmount } from "vue";

const videoRef = ref(null);
const fileInput = ref(null);
const result = ref("");
let scanner = null;

function startScanner() {
  if (!videoRef.value) return;

  scanner = new QrScanner(
    videoRef.value,
    (res) => {
      result.value = res.data;
      scanner.stop();
    },
    {
      highlightScanRegion: true,
      highlightCodeOutline: true,
    },
  );

  scanner.start();
}

function restartScan() {
  result.value = "";
  if (scanner) scanner.start();
}

/* upload QR image */
async function handleFileUpload(e) {
  const file = e.target.files?.[0];
  if (!file) return;

  try {
    const res = await QrScanner.scanImage(file);
    result.value = res;
    if (scanner) scanner.stop();
  } catch (err) {
    console.error("QR not detected");
  }

  if (fileInput.value) fileInput.value.value = "";
}

function triggerUpload() {
  fileInput.value?.click();
}

onMounted(() => {
  startScanner();
});

onBeforeUnmount(() => {
  if (scanner) scanner.stop();
});
</script>

<template>
  <div
    class="min-h-screen bg-gradient-to-b from-black to-gray-900 flex flex-col"
  >
    <!-- Header -->
    <header class="pt-8 pb-4 text-center text-white">
      <h1 class="text-2xl font-semibold">Scan QR Code</h1>
      <p class="text-sm text-gray-400 mt-1">Point your camera at the QR code</p>
    </header>

    <!-- Camera Area -->
    <main class="flex-1 flex items-center justify-center px-6">
      <div class="relative w-full max-w-sm">
        <video ref="videoRef" class="w-full rounded-3xl shadow-lg"></video>

        <!-- Scan Frame -->
        <div
          class="absolute inset-0 flex items-center justify-center pointer-events-none"
        >
          <div
            class="w-64 h-64 border-4 border-green-500 rounded-2xl shadow-[0_0_20px_rgba(34,197,94,0.6)]"
          ></div>
        </div>
      </div>
    </main>

    <!-- Result Panel -->
    <div
      v-if="result"
      class="bg-white rounded-3xl p-6 space-y-4 shadow-2xl mx-6 mb-6"
    >
      <div class="text-center">
        <p class="text-sm text-gray-500">Scanned Link</p>
        <p class="text-green-600 break-all font-medium mt-1">
          {{ result }}
        </p>
      </div>

      <NuxtLink
        :to="result"
        target="_blank"
        class="block w-full text-center bg-green-600 hover:bg-green-700 text-white py-3 rounded-2xl font-semibold"
      >
        Open Card
      </NuxtLink>

      <button
        @click="restartScan"
        class="w-full border border-green-600 text-green-600 py-3 rounded-2xl font-semibold hover:bg-green-50"
      >
        Scan Again
      </button>
    </div>

    <!-- Bottom Navigation -->
    <div class="px-6 flex justify-center">
      <button
        @click="triggerUpload"
        class="px-6 py-3 bg-green-600 hover:bg-green-700 text-white rounded-2xl font-semibold"
      >
        Upload QR Image
      </button>

      <input
        ref="fileInput"
        type="file"
        accept="image/*"
        class="hidden"
        @change="handleFileUpload"
      />
    </div>
    <div class="p-6 text-center">
      <NuxtLink
        to="/create"
        class="text-white border border-white px-6 py-2 rounded-xl hover:bg-white hover:text-black transition"
      >
        Create New Card
      </NuxtLink>
    </div>
  </div>
</template>
