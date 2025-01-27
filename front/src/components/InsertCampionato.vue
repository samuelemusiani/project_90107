<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const nome = ref("");
const luogo = ref("");
const dataInizio = ref("");
const dataFine = ref("");
const tipo = ref("team");
const montepremi = ref(0);

async function insertCampionato() {
  if (new Date(dataInizio.value) > new Date(dataFine.value)) {
    $toast.error("Data di fine deve essere maggiore di data di inizio");
    return;
  }

  try {
    const res = await fetch("/api/campionato", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        luogo: luogo.value,
        dataInizio: new Date(dataInizio.value).toISOString(),
        dataFine: new Date(dataFine.value).toISOString(),
        tipo: tipo.value,
        montepremi: montepremi.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del campionato");
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
    <h3 class="font-bold text-xl">Insert campionato</h3>

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
          <span class="label-text">Data di Inizio</span>
        </div>
        <input
          type="date"
          placeholder="Data di nascita"
          v-model="dataInizio"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Data di Fine</span>
        </div>
        <input
          type="date"
          placeholder="Data di fine"
          v-model="dataFine"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Tipo</span>
        </div>
        <select
          v-model="tipo"
          class="select select-bordered select-primary w-full"
        >
          <option value="team">Teams</option>
          <option value="giocatore">Giocatori</option>
        </select>
      </div>

      <button @click="insertCampionato" class="btn btn-primary mt-4">
        Insert campionato
      </button>
    </form>
  </div>
</template>

<style></style>
