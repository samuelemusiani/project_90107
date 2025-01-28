<script setup lang="ts">
import { formatDate } from "@/utils";
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const response = ref(false);
const evento = ref("");

const numero = ref(0);

async function viewBigliettoiEvento() {
  try {
    const res = await fetch(`/api//biglietti/evento/${evento.value}`, {
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
      Visualizza il numero di bilgietti venduti per un evento
    </h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Nome evento</span>
        </div>
        <input
          type="text"
          placeholder="Nome"
          v-model="evento"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="viewBigliettoiEvento" class="btn btn-primary mt-4">
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
