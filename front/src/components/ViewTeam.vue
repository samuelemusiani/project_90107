<script setup lang="ts">
import { formatDate } from "@/utils";
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const response = ref(false);
const nome = ref("");

const nomeTeam = ref("");
const dataFondazione = ref();
const statoGeografico = ref("");
const gpg = ref("");
const gpv = ref("");
const etaMedia = ref(0.0);
const coach = ref("");
const sponsor = ref("");

async function viewTeam() {
  try {
    const res = await fetch(`/api/team/${nome.value}`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      response.value = false;
      $toast.error("Errore durante l'inserimento del coach");
      throw new Error("Failed to insert user");
    } else {
      const data = await res.json();
      nomeTeam.value = data.nome;
      dataFondazione.value = formatDate(new Date(data.dataFondazione));
      statoGeografico.value = data.statoGeografico;
      gpg.value = data.gpg;
      gpv.value = data.gpv;
      etaMedia.value = data.etaMedia;
      coach.value = data.coach;
      sponsor.value = data.sponsor;
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
    <h3 class="font-bold text-xl">Visualizza team</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Nome team</span>
        </div>
        <input
          type="text"
          placeholder="Nome"
          v-model="nome"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="viewTeam" class="btn btn-primary mt-4">
        Insert coach
      </button>
    </form>

    <div v-if="response">
      <form class="flex flex-col w-full gap-2" @click.prevent>
        <div>
          <div class="label">
            <span class="label-text">Nome team</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="nomeTeam"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Data di fondazione</span>
          </div>
          <input
            type="date"
            placeholder="Nome"
            v-model="dataFondazione"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Stato geografico</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="statoGeografico"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Giocatore più giovane</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="gpg"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Giocatore più vecchio</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="gpv"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Età media giocatori</span>
          </div>
          <input
            type="number"
            placeholder="Nome"
            v-model="etaMedia"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Coach</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="coach"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Sponsor</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="sponsor"
            class="input input-bordered w-full"
            disabled
          />
        </div>
      </form>
    </div>
  </div>
</template>

<style></style>
