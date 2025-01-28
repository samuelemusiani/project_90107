<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

interface Classifica {
  posizione: number;
  teamOGiocatore: number;
}

const $toast = useToast();

const id = ref(0);
const classifica = ref<Classifica[] | undefined>(undefined);

async function fetchClassifica() {
  try {
    const res = await fetch("/api/classifica/" + id.value, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del biglietto");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    classifica.value = data;
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full flex flex-col gap-4">
    <h3 class="font-bold text-xl">Classifica campionato</h3>
    <div class="flex gap-2 items-center justify-evenly">
      <div>ID Evento:</div>
      <input v-model="id" type="number" class="input input-bordered w-24" />
      <button @click="fetchClassifica" class="btn btn-primary">
        Visualizza Classifica
      </button>
    </div>
    <div v-if="classifica">
      <div class="grid grid-cols-2">
        <div class="font-bold text-center">Posizione</div>
        <div class="font-bold text-center">Team o Giocatore</div>
        <template v-for="c in classifica">
          <div class="text-center">
            {{ c.posizione }}
          </div>
          <div class="text-center">
            {{ c.teamOGiocatore }}
          </div>
        </template>
      </div>
    </div>
    <div v-else class="text-center">Nothings to show</div>
  </div>
</template>

<style></style>
