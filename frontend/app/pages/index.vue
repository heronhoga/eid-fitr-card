<script setup lang="ts">
useHead({
  title: 'Eid Card | By Hoga',
  meta: [
    {
      name: 'description',
      content: 'Eid Mubarak',
    },
  ],
})
</script>

<template>
  <div class="min-h-screen bg-gradient-to-b from-green-100 via-white to-green-50 flex flex-col">

    <!-- Header -->
    <header class="pt-10 pb-4 text-center">
      <h1 class="text-3xl font-bold text-green-700">
        🌙 Eid Mubarak
      </h1>
    </header>

    <!-- Main -->
    <main class="flex-1 flex items-center justify-center px-6">

      <div class="bg-white w-full max-w-md rounded-3xl shadow-xl p-8 space-y-6">

        <!-- Illustration / Icon -->
        <div class="text-center text-5xl">
          🕌✨
        </div>

        <!-- Description -->
        <div class="text-center space-y-2">
          <p class="text-gray-700 text-base">
            Scan a card using your camera
          </p>

          <p class="text-gray-500 text-sm">
            Or open a card directly via URL
          </p>

          <p class="text-gray-500 text-sm">
            You can also create a new Eid card
          </p>
        </div>

        <!-- Buttons -->
        <div class="space-y-4 pt-2">

          <NuxtLink
            to="/scan"
            class="block w-full text-center bg-green-600 hover:bg-green-700 text-white py-3 rounded-2xl text-lg font-semibold shadow-md transition active:scale-95"
          >
            📷 Scan Card
          </NuxtLink>

          <NuxtLink
            to="/create"
            class="block w-full text-center bg-white border-2 border-green-600 text-green-600 py-3 rounded-2xl text-lg font-semibold hover:bg-green-50 transition active:scale-95"
          >
            ✨ Create Card
          </NuxtLink>

        </div>

      </div>

    </main>

    <!-- Footer -->
    <footer class="text-center pb-6 text-sm text-gray-400">
      Made with ❤️ by hoga
    </footer>

  </div>
</template>