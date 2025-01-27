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
  <div class="greetings w-full">
    <div class="flex flex-col gap-2">
      <h3 class="font-bold text-xl">Crea team</h3>

      <form class="flex flex-col gap-2" @click.prevent>
        <div>
          <div class="label">
            <span class="label-text">Nome</span>
          </div>
          <input type="text" placeholder="Nome" v-model="nome" class="input input-bordered w-full" />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Logo</span>
          </div>
          <input type="file" placeholder="Logo" disabled class="file-input file-input-bordered w-full" />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Data di fondazione</span>
          </div>
          <input type="date" placeholder="Data di fondazione" v-model="dataFondazione"
            class="input input-bordered w-full" />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Stato geografico</span>
          </div>
          <input type="text" placeholder="Stato geografico" v-model="statoGeografico"
            class="input input-bordered w-full" />
        </div>

        <button @click="creaTeam" class="btn btn-primary mt-2">Crea team</button>
      </form>

      <div class="divider"></div>

      <h3 class="font-bold text-xl">Ingaggia giocatore</h3>

      <form class="flex flex-col gap-2" @click.prevent>
        <div>
          <div class="label">
            <span class="label-text">Nome team</span>
          </div>
          <input type="text" placeholder="Team" v-model="team" class="input input-bordered w-full" />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Username giocatore</span>
          </div>
          <input type="text" placeholder="Giocatore" v-model="giocatore" class="input input-bordered w-full" />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Salario</span>
          </div>
          <input type="number" placeholder="Salario" v-model="salario" class="input input-bordered w-full" />
        </div>

        <button @click="ingaggiaGiocatore" class="btn btn-primary mt-2">Crea team</button>
      </form>
    </div>
  </div>
</template>

<style></style>
