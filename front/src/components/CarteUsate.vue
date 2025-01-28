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

interface Carta {
  id: number;
  nome: number;
  elisir: string;
  danni: number;
}

const $toast = useToast();

const id = ref(0);
const carte = ref<Carta[] | undefined>(undefined);

const type = computed(() => {
  switch ($props.op) {
    case 16:
      return "campionato";
    case 17:
      return "evento";
    case 18:
      return "partita";
    default:
      return "campionato";
  }
});

async function fetchClassifica() {
  try {
    const res = await fetch(`/api/carte_usate/${type.value}/${id.value}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      $toast.error("Errore durante la fetch per la carte usate");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    carte.value = data;
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full flex flex-col gap-4">
    <h3 class="font-bold text-xl">Carta più usate {{ type }}</h3>
    <div class="flex gap-2 items-center justify-evenly">
      <div>ID {{ type }}:</div>
      <input v-model="id" type="number" class="input input-bordered w-24" />
      <button @click="fetchClassifica" class="btn btn-primary">
        Visualizza Carta
      </button>
    </div>
    <div v-if="carte">
      <div class="grid grid-cols-2">
        <div class="font-bold text-center">Nome</div>
        <div class="font-bold text-center">Volte giocata</div>
        <template v-for="c in carte">
          <div class="text-center">
            {{ c.nome }}
          </div>
          <div class="text-center">
            {{ c.id }}
          </div>
        </template>
      </div>
    </div>
    <div v-else class="text-center">Nothings to show</div>
  </div>
</template>

<style></style>
