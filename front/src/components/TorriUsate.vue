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

interface TorriUsate {
  volte: number;
  tipoTorri: string;
}

const $toast = useToast();

const id = ref(0);
const statistica = ref<TorriUsate | undefined>(undefined);

const type = computed(() => {
  switch ($props.op) {
    case 26:
      return "campionato";
    case 27:
      return "evento";
    default:
      return "campionato";
  }
});

async function fetchClassifica() {
  try {
    const res = await fetch(
      `/api/torri_usate/${type.value}/${id.value}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      },
    );

    if (!res.ok) {
      $toast.error("Errore durante la fetch per la torri usate");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    statistica.value = data;
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => { });
</script>

<template>
  <div class="greetings w-full flex flex-col gap-4">
    <h3 class="font-bold text-xl">Statistiche {{ type }}</h3>
    <div class="flex gap-2 items-center justify-evenly">
      <div>ID {{ type }}:</div>
      <input v-model="id" type="number" class="input input-bordered w-24" />
    </div>
    <button @click="fetchClassifica" class="btn btn-primary">
      Visualizza torri
    </button>
    <div v-if="statistica">
      <form>
        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Volte usate </label>
            <input types="text" v-model="statistica.volte" class="input input-bordered" disabled />
          </div>
        </div>

        <div class="form-control">
          <div class="label">
            <label class="lable-text"> Tipo torri </label>
            <input types="text" v-model="statistica.tipoTorri" class="input input-bordered" disabled />
          </div>
        </div>
      </form>
    </div>
    <div v-else class="text-center">Nothings to show</div>
  </div>
</template>

<style></style>
