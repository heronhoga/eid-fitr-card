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
  <div class="create-root">
    <!-- Background -->
    <div class="sky"></div>

    <main class="content-wrap">
      <div class="create-card">
        <!-- Header -->
        <div class="header">
          <p class="label">Create Eid Card</p>
          <h1 class="arabic">عيد مبارك</h1>
          <p class="subtitle">Send a digital Eid greeting</p>
        </div>

        <!-- Form -->
        <form @submit.prevent="submitForm" class="form">
          <!-- To -->
          <div class="field">
            <label>To</label>
            <input
              v-model="form.to"
              type="text"
              placeholder="Recipient name"
              maxlength="32"
              required
            />
            <span class="counter">{{ form.to.length }}/32</span>
          </div>

          <!-- From -->
          <div class="field">
            <label>From</label>
            <input
              v-model="form.from"
              type="text"
              placeholder="Your name"
              maxlength="32"
              required
            />
            <span class="counter">{{ form.from.length }}/32</span>
          </div>

          <!-- Message -->
          <div class="field">
            <label>Message</label>
            <textarea
              v-model="form.description"
              rows="4"
              maxlength="256"
              placeholder="Write your Eid wishes..."
              required
            ></textarea>
            <span class="counter">{{ form.description.length }}/256</span>
          </div>

          <!-- Submit -->
          <button class="btn primary" :disabled="loading">
            {{ loading ? "Creating Card..." : "Create Card" }}
          </button>

          <p v-if="error" class="error">{{ error }}</p>
          <p v-if="success" class="success">Card created successfully</p>
        </form>

        <p class="alt-link">
          or
          <NuxtLink to="/scan">scan a QR card</NuxtLink>
        </p>
      </div>
    </main>

    <!-- QR Modal -->
    <div v-if="showQR" class="modal">
      <div class="modal-card">
        <h2 class="modal-title">📷 Scan Your Card</h2>

        <img :src="qrCodeUrl" alt="QR Code" class="qr" />

        <div class="modal-actions">
          <button @click="downloadQR" class="btn outline">
            Download QR Code
          </button>

          <button @click="closeQR" class="btn primary">Close</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.create-root {
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
  max-width: 480px;
  padding: 1rem;
}

.create-card {
  background: linear-gradient(
    160deg,
    rgba(15, 28, 55, 0.92),
    rgba(8, 18, 36, 0.97)
  );
  border: 1px solid rgba(255, 210, 70, 0.25);
  border-radius: 28px;
  padding: 2.5rem 2rem;
  color: #e8d5a0;
  box-shadow:
    0 0 0 1px rgba(255, 210, 70, 0.08),
    0 8px 40px rgba(0, 0, 0, 0.6);
}

.header {
  text-align: center;
  margin-bottom: 1.5rem;
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
  color: #ffd54f;
  margin-top: 4px;
}

.subtitle {
  font-size: 0.85rem;
  opacity: 0.6;
  margin-top: 4px;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.field label {
  font-size: 0.75rem;
  letter-spacing: 0.15em;
  text-transform: uppercase;
  color: #10b981;
}

.field input,
.field textarea {
  width: 100%;
  margin-top: 6px;
  padding: 10px 12px;
  border-radius: 12px;
  border: 1px solid rgba(16, 185, 129, 0.25);
  background: rgba(255, 255, 255, 0.04);
  color: white;
  outline: none;
}

.field textarea {
  resize: none;
}

.counter {
  display: block;
  font-size: 0.7rem;
  opacity: 0.5;
  text-align: right;
}

.btn {
  padding: 12px;
  border-radius: 14px;
  font-weight: 600;
  cursor: pointer;
  border: none;
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
  background: transparent;
}

.error {
  color: #ef9a9a;
  font-size: 0.85rem;
  text-align: center;
}

.success {
  color: #10b981;
  font-size: 0.85rem;
  text-align: center;
}

.alt-link {
  text-align: center;
  margin-top: 1rem;
  font-size: 0.8rem;
}

.alt-link a {
  color: #10b981;
  text-decoration: none;
}

.modal {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 50;
}

.modal-card {
  background: linear-gradient(
    160deg,
    rgba(15, 28, 55, 0.95),
    rgba(8, 18, 36, 0.98)
  );
  border: 1px solid rgba(255, 210, 70, 0.25);
  border-radius: 24px;
  padding: 2rem;
  text-align: center;
  width: 320px;
}

.modal-title {
  color: #ffd54f;
  margin-bottom: 1rem;
}

.qr {
  width: 180px;
  margin: 0 auto 1rem;
}

.modal-actions {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
</style>
