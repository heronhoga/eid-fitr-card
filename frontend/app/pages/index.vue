<script setup lang="ts">
useHead({
  title: "Eid Card | By Hoga",
  meta: [
    {
      name: "description",
      content: "Eid Mubarak",
    },
  ],
});
</script>

<template>
  <div class="home-root">
    <!-- Night sky background -->
    <div class="sky"></div>

    <main class="content-wrap">
      <div class="home-card">
        <!-- Header -->
        <div class="header">
          <p class="label">Eid Mubarak</p>
          <h1 class="arabic">عيد مبارك</h1>
        </div>

        <!-- Icon -->
        <div class="icon">🕌✨</div>

        <!-- Description -->
        <div class="desc">
          <p>Scan a card using your camera</p>
          <p class="sub">Or open a card directly via URL</p>
          <p class="sub">You can also create a new Eid card</p>
        </div>

        <!-- Buttons -->
        <div class="actions">
          <NuxtLink to="/scan" class="btn primary"> 📷 Scan Card </NuxtLink>

          <NuxtLink to="/create" class="btn outline"> ✨ Create Card </NuxtLink>
        </div>

        <p class="footer">Made with ❤️ by hoga</p>
      </div>
    </main>
  </div>
</template>

<style scoped>
.home-root {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #04060f;
  overflow: hidden;
}

.sky {
  position: fixed;
  inset: 0;
  background: radial-gradient(
    ellipse at 60% 0%,
    #1a2a4a 0%,
    #060c1a 55%,
    #010308 100%
  );
  z-index: 0;
}

.content-wrap {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 420px;
  padding: 1rem;
}

.home-card {
  background: linear-gradient(
    160deg,
    rgba(15, 28, 55, 0.92) 0%,
    rgba(8, 18, 36, 0.97) 100%
  );
  border: 1px solid rgba(255, 210, 70, 0.25);
  border-radius: 28px;
  padding: 2.5rem 2rem;
  text-align: center;
  color: #e8d5a0;
  box-shadow:
    0 0 0 1px rgba(255, 210, 70, 0.08),
    0 8px 40px rgba(0, 0, 0, 0.6);
}

.label {
  font-size: 0.65rem;
  letter-spacing: 0.35em;
  text-transform: uppercase;
  color: #ffd54f;
  opacity: 0.7;
}

.arabic {
  font-size: 1.8rem;
  margin-top: 6px;
  color: #ffd54f;
  text-shadow: 0 0 24px rgba(255, 210, 70, 0.5);
}

.icon {
  font-size: 2.5rem;
  margin: 1.5rem 0;
}

.desc p {
  font-size: 0.95rem;
  margin: 4px 0;
}

.desc .sub {
  opacity: 0.65;
}

.actions {
  margin-top: 1.6rem;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.btn {
  display: block;
  padding: 12px;
  border-radius: 16px;
  font-weight: 600;
  text-decoration: none;
  transition: 0.25s;
}

.btn.primary {
  background: #10b981;
  color: white;
}

.btn.primary:hover {
  background: #059669;
}

.btn.outline {
  border: 1px solid #10b981;
  color: #10b981;
}

.btn.outline:hover {
  background: rgba(16, 185, 129, 0.1);
}

.footer {
  margin-top: 1.6rem;
  font-size: 0.75rem;
  opacity: 0.5;
}
</style>
