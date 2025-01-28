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

const $toast = useToast();

const response = ref(false);
const nome = ref("");

const numero = ref(0);

const type = computed(() => {
  switch ($props.op) {
    case 28:
      return "campionato";
    case 29:
      return "evento";
    default:
      return "campionato";
  }
});

async function viewBiglietti() {
  try {
    const res = await fetch(`/api/biglietti/${type.value}/${nome.value}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      response.value = false;
      $toast.error("Errore durante la visualizzazione del team");
      throw new Error("Failed to insert user");
    } else {
      const data = await res.json();
      numero.value = data.numero;
      response.value = true;
    }
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full">
    <h3 class="font-bold text-xl">
      Visualizza il numero di bilgietti venduti per un {{ type }}
    </h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Nome {{ type }}</span>
        </div>
        <input
          type="text"
          placeholder="Nome"
          v-model="nome"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="viewBiglietti" class="btn btn-primary mt-4">
        Visualizza
      </button>
    </form>

    <div v-if="response">
      <h3 class="font-bold text-xl">Numero di biglietti venduti</h3>
      <p>{{ numero }}</p>
    </div>
  </div>
</template>

<style></style>
