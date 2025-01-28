<script setup lang="ts">
import { computed } from "@vue/reactivity";
import { ref, onMounted, defineProps } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $props = defineProps({
  op: {
    type: Number,
    required: true,
  },
});

interface Statistica {
  elisirUsato: number;
  elisirSprecato: number;
  danniFatti: number;
  danniSubiti: number;
}

const $toast = useToast();

const idGiocatore = ref(0);
const idTorneoEventoPartita = ref(0);
const statistica = ref<Statistica | undefined>(undefined);

const type = computed(() => {
  switch ($props.op) {
    case 19:
      return "campionato";
    case 20:
      return "evento";
    case 21:
      return "partita";
    default:
      return "campionato";
  }
});

async function fetchClassifica() {
  try {
    const res = await fetch(
      `/api/statistiche/${type.value}/${idGiocatore.value}/${idTorneoEventoPartita.value}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      },
    );

    if (!res.ok) {
      $toast.error("Errore durante la fetch per la carte usate");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    statistica.value = data;
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full flex flex-col gap-4">
    <h3 class="font-bold text-xl">Statistiche {{ type }}</h3>
    <div class="flex gap-2 items-center justify-evenly">
      <div>ID giocatore:</div>
      <input
        v-model="idGiocatore"
        type="number"
        class="input input-bordered w-24"
      />
    </div>
    <div class="flex gap-2 items-center justify-evenly">
      <div>ID {{ type }}:</div>
      <input
        v-model="idTorneoEventoPartita"
        type="number"
        class="input input-bordered w-24"
      />
    </div>
    <button @click="fetchClassifica" class="btn btn-primary">
      Visualizza Carta
    </button>
    <div v-if="statistica">
      <form>
        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Elisir usato </label>
            <input
              types="text"
              v-model="statistica.elisirUsato"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Elisir sprecato </label>
            <input
              types="text"
              v-model="statistica.elisirSprecato"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Danni fatti </label>
            <input
              types="text"
              v-model="statistica.danniFatti"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Danni subiti </label>
            <input
              types="text"
              v-model="statistica.danniSubiti"
              class="input input-bordered"
              disabled
            />
          </div>
        </div>
      </form>
    </div>
    <div v-else class="text-center">Nothings to show</div>
  </div>
</template>

<style></style>
