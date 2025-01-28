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

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full">
    <h3 class="font-bold text-xl">Insert mazzo</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div class="flex gap-4 items-center">
        <button @click="createMazzo" class="btn btn-secondary mt-4">
          Crea Mazzo
        </button>
        <div class="flex items-center">ID: {{ id }}</div>
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

      <button @click="" class="btn btn-primary mt-4">
        Insert Carta into Mazzo
      </button>
    </form>
  </div>
</template>

<style></style>
