<script setup lang="ts">
import { ref, onMounted, onUnmounted } from "vue";

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

// Music
const isPlaying = ref(false);
const audio = ref<HTMLAudioElement | null>(null);

function toggleMusic() {
  if (!audio.value) {
    // Replace this src with your actual nasheed/music file path
    audio.value = new Audio("/eid-music.mp3");
    audio.value.loop = true;
    audio.value.volume = 0.4;
  }

  if (isPlaying.value) {
    audio.value.pause();
    isPlaying.value = false;
  } else {
    audio.value.play().catch(() => {});
    isPlaying.value = true;
  }
}

onUnmounted(() => {
  if (audio.value) {
    audio.value.pause();
    audio.value = null;
  }
});

// Particle stars
const stars = Array.from({ length: 36 }, (_, i) => ({
  id: i,
  x: Math.random() * 100,
  y: Math.random() * 100,
  size: Math.random() * 3 + 1,
  delay: Math.random() * 4,
  duration: 2 + Math.random() * 3,
}));

// Lanterns
const lanterns = [
  { left: "8%", delay: "0s", color: "#f59e0b" },
  { left: "22%", delay: "0.6s", color: "#10b981" },
  { left: "75%", delay: "0.3s", color: "#f59e0b" },
  { left: "90%", delay: "1s", color: "#10b981" },
];
</script>

<template>
  <div class="eid-root">
    <!-- Night sky background -->
    <div class="sky">
      <!-- Stars -->
      <div
        v-for="star in stars"
        :key="star.id"
        class="star"
        :style="{
          left: star.x + '%',
          top: star.y + '%',
          width: star.size + 'px',
          height: star.size + 'px',
          animationDelay: star.delay + 's',
          animationDuration: star.duration + 's',
        }"
      />

      <!-- Moon -->
      <div class="moon-wrap">
        <div class="moon" />
        <div class="moon-glow" />
      </div>

      <!-- Lanterns -->
      <div
        v-for="(l, i) in lanterns"
        :key="i"
        class="lantern"
        :style="{ left: l.left, animationDelay: l.delay, '--lc': l.color }"
      >
        <div class="lantern-body" />
        <div class="lantern-glow" />
        <div class="lantern-string" />
      </div>
    </div>

    <!-- Music toggle -->
    <button class="music-btn" @click="toggleMusic" :title="isPlaying ? 'Pause Music' : 'Play Music'">
      <span v-if="isPlaying" class="music-icon playing">
        <span /><span /><span /><span />
      </span>
      <span v-else class="music-icon paused">▶</span>
    </button>

    <!-- Main content -->
    <main class="content-wrap">

      <!-- Loading -->
      <div v-if="pending" class="glass-card loading-card">
        <div class="loader-ring" />
        <p class="loading-text">Loading your Eid card…</p>
      </div>

      <!-- Error -->
      <div v-else-if="error" class="glass-card error-card">
        <div class="error-moon">🌙</div>
        <h2 class="error-title">Card Not Found</h2>
        <p class="error-body">The Eid card you're looking for couldn't be loaded.</p>
      </div>

      <!-- Card -->
      <div v-else-if="data" class="eid-card">

        <!-- Decorative top ornament -->
        <div class="ornament top-ornament">
          <span class="orn-line" /><span class="orn-star">✦</span><span class="orn-line" />
        </div>

        <!-- Header -->
        <div class="card-header">
          <p class="label">Eid Mubarak</p>
          <div class="arabic-text">عيد مبارك</div>
        </div>

        <!-- To -->
        <div class="to-section">
          <span class="to-label">For</span>
          <h1 class="to-name">{{ data.data.to }}</h1>
        </div>

        <!-- Divider -->
        <div class="divider">
          <span class="div-line" /><span class="div-crescent">☽</span><span class="div-line" />
        </div>

        <!-- Message -->
        <div class="message-box">
          <p class="message-text">{{ data.data.description }}</p>
        </div>

        <!-- Divider -->
        <div class="divider">
          <span class="div-line" /><span class="div-star">✦</span><span class="div-line" />
        </div>

        <!-- From -->
        <div class="from-section">
          <span class="from-label">With love from</span>
          <p class="from-name">{{ data.data.from }}</p>
        </div>

        <!-- Bottom ornament -->
        <div class="ornament bottom-ornament">
          <span class="orn-line" /><span class="orn-star">✦</span><span class="orn-line" />
        </div>

        <!-- CTA -->
        <NuxtLink to="/" class="cta-link">
          Create your own Eid card →
        </NuxtLink>

      </div>
    </main>
  </div>
</template>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Cinzel+Decorative:wght@400;700&family=Lora:ital,wght@0,400;0,600;1,400&display=swap');

/* ─── Root ─────────────────────────────────── */
.eid-root {
  min-height: 100vh;
  position: relative;
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem 1rem;
  background: #04060f;
}

/* ─── Sky ───────────────────────────────────── */
.sky {
  position: fixed;
  inset: 0;
  background: radial-gradient(ellipse at 60% 0%, #1a2a4a 0%, #060c1a 55%, #010308 100%);
  z-index: 0;
  overflow: hidden;
}

/* Stars */
.star {
  position: absolute;
  border-radius: 50%;
  background: #fffde7;
  animation: twinkle ease-in-out infinite;
  opacity: 0.7;
}
@keyframes twinkle {
  0%, 100% { opacity: 0.2; transform: scale(0.8); }
  50% { opacity: 1; transform: scale(1.3); }
}

/* Moon */
.moon-wrap {
  position: absolute;
  top: 5%;
  right: 12%;
}
.moon {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  background: radial-gradient(circle at 35% 35%, #fff8e1, #ffd54f 60%, #ffb300);
  box-shadow: 0 0 40px 12px rgba(255, 200, 50, 0.35);
  position: relative;
}
.moon-glow {
  position: absolute;
  inset: -20px;
  border-radius: 50%;
  background: radial-gradient(circle, rgba(255,210,60,0.15) 0%, transparent 70%);
  animation: moonpulse 4s ease-in-out infinite;
}
@keyframes moonpulse {
  0%, 100% { transform: scale(1); opacity: 0.7; }
  50% { transform: scale(1.15); opacity: 1; }
}

/* Lanterns */
.lantern {
  position: absolute;
  top: 0;
  display: flex;
  flex-direction: column;
  align-items: center;
  animation: swing 3.5s ease-in-out infinite;
  transform-origin: top center;
}
.lantern-string {
  width: 1px;
  height: 60px;
  background: rgba(255,220,100,0.4);
  margin-bottom: 0;
}
.lantern-body {
  width: 22px;
  height: 34px;
  background: var(--lc, #f59e0b);
  border-radius: 5px 5px 8px 8px;
  box-shadow: 0 0 18px 6px color-mix(in srgb, var(--lc) 60%, transparent);
  position: relative;
}
.lantern-body::before,
.lantern-body::after {
  content: '';
  position: absolute;
  left: 50%;
  transform: translateX(-50%);
  width: 28px;
  height: 6px;
  background: rgba(255,255,255,0.2);
  border-radius: 3px;
}
.lantern-body::before { top: -4px; }
.lantern-body::after { bottom: -4px; }
.lantern-glow {
  position: absolute;
  bottom: 0;
  width: 50px;
  height: 50px;
  border-radius: 50%;
  background: radial-gradient(circle, color-mix(in srgb, var(--lc) 40%, transparent), transparent 70%);
  transform: translateY(30%);
}
@keyframes swing {
  0%, 100% { transform: rotate(-5deg); }
  50% { transform: rotate(5deg); }
}

/* ─── Music Button ──────────────────────────── */
.music-btn {
  position: fixed;
  bottom: 1.5rem;
  right: 1.5rem;
  z-index: 100;
  width: 48px;
  height: 48px;
  border-radius: 50%;
  border: 1.5px solid rgba(255,210,80,0.4);
  background: rgba(10, 20, 40, 0.7);
  backdrop-filter: blur(10px);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: border-color 0.3s, box-shadow 0.3s;
  color: #ffd54f;
  font-size: 0.85rem;
}
.music-btn:hover {
  border-color: #ffd54f;
  box-shadow: 0 0 16px rgba(255, 213, 79, 0.3);
}
.music-icon.playing {
  display: flex;
  align-items: flex-end;
  gap: 2px;
  height: 18px;
}
.music-icon.playing span {
  display: block;
  width: 3px;
  background: #ffd54f;
  border-radius: 2px;
  animation: bar 0.8s ease-in-out infinite alternate;
}
.music-icon.playing span:nth-child(1) { height: 10px; animation-delay: 0s; }
.music-icon.playing span:nth-child(2) { height: 18px; animation-delay: 0.15s; }
.music-icon.playing span:nth-child(3) { height: 13px; animation-delay: 0.3s; }
.music-icon.playing span:nth-child(4) { height: 7px; animation-delay: 0.45s; }
@keyframes bar {
  from { transform: scaleY(0.4); }
  to { transform: scaleY(1); }
}

/* ─── Content ───────────────────────────────── */
.content-wrap {
  position: relative;
  z-index: 10;
  width: 100%;
  max-width: 480px;
}

/* ─── Glass util card (loading/error) ──────── */
.glass-card {
  background: rgba(255,255,255,0.05);
  border: 1px solid rgba(255,220,80,0.18);
  backdrop-filter: blur(16px);
  border-radius: 24px;
  padding: 2.5rem;
  text-align: center;
  color: #e8d5a0;
}
.loader-ring {
  width: 44px; height: 44px;
  border: 3px solid rgba(255,210,60,0.2);
  border-top-color: #ffd54f;
  border-radius: 50%;
  margin: 0 auto 1.2rem;
  animation: spin 0.9s linear infinite;
}
@keyframes spin { to { transform: rotate(360deg); } }
.loading-text { font-style: italic; opacity: 0.7; }
.error-moon { font-size: 2.5rem; margin-bottom: 0.5rem; }
.error-title { font-family: 'Cinzel Decorative', serif; color: #ef9a9a; font-size: 1.2rem; margin-bottom: 0.5rem; }
.error-body { opacity: 0.7; font-style: italic; }

/* ─── Main Eid Card ─────────────────────────── */
.eid-card {
  background: linear-gradient(160deg,
    rgba(15, 28, 55, 0.92) 0%,
    rgba(8, 18, 36, 0.97) 100%
  );
  border: 1px solid rgba(255, 210, 70, 0.25);
  border-radius: 28px;
  padding: 2.5rem 2rem;
  text-align: center;
  color: #e8d5a0;
  box-shadow:
    0 0 0 1px rgba(255,210,70,0.08),
    0 8px 40px rgba(0,0,0,0.6),
    inset 0 1px 0 rgba(255,220,100,0.1);
  animation: cardrise 0.8s cubic-bezier(0.22, 1, 0.36, 1) both;
  position: relative;
  overflow: hidden;
}
.eid-card::before {
  content: '';
  position: absolute;
  inset: 0;
  background:
    radial-gradient(ellipse at 10% 10%, rgba(16, 185, 129, 0.06) 0%, transparent 50%),
    radial-gradient(ellipse at 90% 90%, rgba(245, 158, 11, 0.08) 0%, transparent 50%);
  pointer-events: none;
}
@keyframes cardrise {
  from { opacity: 0; transform: translateY(30px) scale(0.97); }
  to { opacity: 1; transform: translateY(0) scale(1); }
}

/* Ornaments */
.ornament {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1.2rem;
}
.bottom-ornament { margin-bottom: 0; margin-top: 1.2rem; }
.orn-line {
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(255,210,70,0.4), transparent);
}
.orn-star { color: #ffd54f; font-size: 0.6rem; opacity: 0.8; }

/* Header */
.card-header { margin-bottom: 1.4rem; }
.label {
  font-size: 0.65rem;
  letter-spacing: 0.35em;
  text-transform: uppercase;
  color: #ffd54f;
  opacity: 0.7;
  margin-bottom: 0.5rem;
}
.arabic-text {
  font-size: 2rem;
  color: #ffd54f;
  text-shadow: 0 0 24px rgba(255,210,70,0.5);
  line-height: 1.2;
  font-weight: 700;
}

/* To */
.to-section { margin-bottom: 1.2rem; }
.to-label {
  display: block;
  font-size: 0.65rem;
  letter-spacing: 0.3em;
  text-transform: uppercase;
  color: #10b981;
  opacity: 0.7;
  margin-bottom: 0.3rem;
}
.to-name {
  font-size: clamp(1.3rem, 5vw, 2rem);
  font-weight: 700;
  color: #ffffff;
  text-shadow: 0 0 30px rgba(255,255,255,0.2);
  line-height: 1.2;
  word-break: break-word;
  margin: 0;
}

/* Divider */
.divider {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 1rem 0;
}
.div-line {
  flex: 1;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(16,185,129,0.25), transparent);
}
.div-crescent, .div-star {
  font-size: 0.75rem;
  color: #10b981;
  opacity: 0.7;
}

/* Message */
.message-box {
  background: rgba(16, 185, 129, 0.07);
  border: 1px solid rgba(16, 185, 129, 0.15);
  border-radius: 16px;
  padding: 1.4rem 1.2rem;
  max-height: 200px;
  overflow-y: auto;
  text-align: left;
  scrollbar-width: thin;
  scrollbar-color: rgba(255,210,70,0.3) transparent;
}
.message-text {
  font-size: 0.97rem;
  line-height: 1.8;
  color: #d4c9a8;
  font-style: italic;
  word-break: break-word;
  margin: 0;
}

/* From */
.from-section { margin-top: 1rem; }
.from-label {
  display: block;
  font-size: 0.6rem;
  letter-spacing: 0.3em;
  text-transform: uppercase;
  color: #ffd54f;
  opacity: 0.5;
  margin-bottom: 0.25rem;
}
.from-name {
  font-size: 1rem;
  color: #ffd54f;
  margin: 0;
  word-break: break-word;
}

/* CTA */
.cta-link {
  display: inline-block;
  margin-top: 1.6rem;
  font-size: 0.78rem;
  letter-spacing: 0.1em;
  color: #10b981;
  text-decoration: none;
  opacity: 0.75;
  transition: opacity 0.2s, letter-spacing 0.2s;
}
.cta-link:hover {
  opacity: 1;
  letter-spacing: 0.15em;
}
</style>