<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const id = ref(0);
const carta = ref("");

async function createMazzo() {
  try {
    const res = await fetch("/api/mazzo", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del mazzo");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    id.value = data.id;
  } catch (err) {
    console.error(err);
  }
}

async function insertFormato() {
  try {
    const res = await fetch("http://130.136.3.143:3000/api/formato", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        mazzo: id.value,
        carta: carta.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento della carta nel mazzo");
      throw new Error("Failed to insert user");
    }
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full">
    <h3 class="font-bold text-xl">Insert mazzo</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div class="flex gap-4 items-center mt-4">
        <input type="text" v-model="id" class="input input-bordered w-full" />

        <button @click="createMazzo" class="btn btn-secondary">
          Crea Mazzo
        </button>
      </div>

      <div>
        <div class="label">
          <span class="label-text">Carta</span>
        </div>
        <input
          type="text"
          placeholder="Nome carta"
          v-model="carta"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="insertFormato" class="btn btn-primary mt-4">
        Insert Carta into Mazzo
      </button>
    </form>
  </div>
</template>

<style></style>
