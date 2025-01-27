<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from 'vue-toast-notification'
import 'vue-toast-notification/dist/theme-sugar.css';

const $toast = useToast()

const nome = ref("");
const logo = ref("");
const dataFondazione = ref("");
const statoGeografico = ref("");

const team = ref("");
const giocatore = ref("");
const salario = ref(0);

async function creaTeam() {
  try {
    const res = await fetch("/api/team", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        logo: logo.value,
        dataFondazione: new Date(dataFondazione.value).toISOString(),
        statoGeografico: statoGeografico.value,
      }),
    });

    if (!res.ok) {
      $toast.error('Errore durante l\'inserimento del giocatore')
      throw new Error("Failed to insert user");
    }
  } catch (err) {
    console.error(err);
  }
}

async function ingaggiaGiocatore() {
  try {
    const res = await fetch("/api/ingaggio", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        team: team.value,
        giocatore: giocatore.value,
        salario: salario.value,
      }),
    });

    if (!res.ok) {
      $toast.error('Errore durante l\'inserimento del giocatore')
      throw new Error("Failed to insert user");
    }
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {
});
</script>

<template>
  <div class="greetings">
    <div class="flex flex-col gap-2">
      <h3 class="font-bold text-xl">Crea team</h3>
      <form class="flex flex-col max-w-sm gap-2" @click.prevent>
        <input type="text" placeholder="Nome" v-model="nome" class="input input-bordered" />
        <input type="file" placeholder="Logo" disabled class="input input-bordered" />
        <input type="date" placeholder="Data di fondazione" v-model="dataFondazione" class="input input-bordered" />
        <input type="text" placeholder="Stato geografico" v-model="statoGeografico" class="input input-bordered" />

        <button @click="creaTeam" class="btn btn-primary">Crea team</button>
      </form>

      <h3 class="font-bold text-xl">Ingaggia giocatore</h3>
      <form class="flex flex-col max-w-sm gap-2" @click.prevent>
        <input type="text" placeholder="Team" v-model="team" class="input input-bordered" />
        <input type="text" placeholder="Giocatore" v-model="giocatore" class="input input-bordered" />
        <input type="number" placeholder="Salario" v-model="salario" class="input input-bordered" />

        <button @click="ingaggiaGiocatore" class="btn btn-primary">Crea team</button>
      </form>
    </div>
  </div>
</template>

<style></style>
