<template>
  <div class="min-h-screen bg-gradient-to-b from-green-50 to-white flex flex-col">

    <!-- Header -->
    <header class="p-4 text-center">
      <h1 class="text-2xl font-bold text-green-700">
        Eid Mubarak
      </h1>
    </header>

    <!-- Main Content -->
    <main class="flex-1 flex flex-col justify-center px-6 pb-10">

      <div class="text-center space-y-3 mb-10">
        <p class="text-gray-700 text-base">
          Scan with your camera
        </p>
        <p class="text-gray-500 text-sm">
          Or enter the URL manually
        </p>
        <p class="text-gray-500 text-sm">
          You can also create a new card
        </p>
      </div>

      <!-- Action Buttons -->
      <div class="space-y-4">

        <NuxtLink
          to="/scan"
          class="block w-full text-center bg-green-600 text-white py-3 rounded-2xl text-lg font-semibold shadow-md active:scale-95 transition"
        >
          Scan Card
        </NuxtLink>

        <NuxtLink
          to="/create"
          class="block w-full text-center bg-white border border-green-600 text-green-600 py-3 rounded-2xl text-lg font-semibold shadow-sm active:scale-95 transition"
        >
          Create Card
        </NuxtLink>

      </div>

    </main>

  </div>
</template>