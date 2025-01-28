<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

interface Carta {
  id: number;
  nome: number;
  elisir: string;
  danni: number;
}

const $toast = useToast();

const id = ref(0);
const carta = ref<Carta | undefined>(undefined);

async function fetchClassifica() {
  try {
    const res = await fetch("/api/carta_piu_usata/partita/" + id.value, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      $toast.error("Errore durante la fetch per la carta");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    carta.value = data;
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full flex flex-col gap-4">
    <h3 class="font-bold text-xl">Carta più usata partita</h3>
    <div class="flex gap-2 items-center justify-evenly">
      <div>ID partita:</div>
      <input v-model="id" type="number" class="input input-bordered w-24" />
      <button @click="fetchClassifica" class="btn btn-primary">
        Visualizza Carta
      </button>
    </div>
    <div v-if="carta">
      <form>
        <div class="form-control">
          <div class="label">
            <label class="lable-text">
              Numero di volte in cui è stata usata
            </label>
            <input
              types="text"
              v-model="carta.id"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Nome </label>
            <input
              types="text"
              v-model="carta.nome"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Elisir </label>
            <input
              types="text"
              v-model="carta.elisir"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Danni </label>
            <input
              types="text"
              v-model="carta.danni"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>
      </form>
    </div>
    <div v-else class="text-center">Niente da vedere</div>
  </div>
</template>

<style></style>
