<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const nome = ref("");
const luogo = ref("");
const data = ref("");
const postiTotali = ref(0);
const campionato = ref("");

async function insertEvento() {
  try {
    const res = await fetch("/api/evento", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        luogo: luogo.value,
        data: new Date(data.value).toISOString(),
        postiTotali: postiTotali.value,
        campionato: campionato.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento dell'evento");
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
    <h3 class="font-bold text-xl">Insert evento</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Nome</span>
        </div>
        <input
          type="text"
          placeholder="Nome"
          v-model="nome"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Luogo</span>
        </div>
        <input
          type="text"
          placeholder="Luogo"
          v-model="luogo"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Data</span>
        </div>
        <input
          type="date"
          placeholder="Data"
          v-model="data"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Posti totali</span>
        </div>
        <input
          type="number"
          placeholder="Posti totali"
          v-model="postiTotali"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Nome Campionato</span>
        </div>
        <input
          type="text"
          placeholder="Nome campionato"
          v-model="campionato"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="insertEvento" class="btn btn-primary mt-4">
        Insert evento
      </button>
    </form>
  </div>
</template>

<style></style>
