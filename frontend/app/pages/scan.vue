<script setup>
import QrScanner from "qr-scanner";
import { ref, onMounted, onBeforeUnmount } from "vue";

const videoRef = ref(null);
const result = ref("");
let scanner = null;

onMounted(() => {
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
});

onBeforeUnmount(() => {
  if (scanner) scanner.stop();
});
</script>

<template>
  <div class="min-h-screen bg-black flex flex-col">
    <!-- Header -->
    <header class="p-4 text-white text-center">
      <h1 class="text-lg font-semibold">Scan QR Code</h1>
    </header>

    <!-- Camera Preview -->
    <div class="flex-1 flex items-center justify-center px-4">
      <div class="w-full max-w-sm relative">
        <video ref="videoRef" class="w-full rounded-2xl"></video>

        <!-- Scan Overlay Box -->
        <div
          class="absolute inset-0 flex items-center justify-center pointer-events-none"
        >
          <div class="w-56 h-56 border-4 border-green-500 rounded-2xl"></div>
        </div>
      </div>
    </div>

    <!-- Result -->
    <div v-if="result" class="bg-white p-4 rounded-t-3xl">
      <p class="text-sm text-gray-500 mb-1">Scanned Result:</p>
      <p class="text-green-600 break-all font-medium">
        {{ result }}
      </p>

      <NuxtLink
        :to="result"
        class="block mt-4 w-full text-center bg-green-600 text-white py-3 rounded-2xl font-semibold"
      >
        Open Link
      </NuxtLink>
    </div>

    <NuxtLink
      to="/create"
      class="m-5 p-2 text-white text-center border border-white rounded-2xl"
      >Or create new card</NuxtLink>
  </div>
</template>
