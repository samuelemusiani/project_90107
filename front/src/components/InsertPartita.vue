<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";
import { formatTime, rand } from "@/utils";

const tipoTorri = ["base", "cannone", "balestra"];

const $toast = useToast();

const vincitore = ref(false);
const orario = ref(formatTime(new Date()));
const tempo = ref(rand(40, 580));
const evento = ref("");
const giocatore1 = ref("");
const giocatore2 = ref("");
const mazzo1 = ref(0);
const mazzo2 = ref(0);
const elisirUsato1 = ref(rand(0, tempo.value));
const elisirUsato2 = ref(rand(0, tempo.value));
const elisirSprecato1 = ref(tempo.value - elisirUsato1.value);
const elisirSprecato2 = ref(tempo.value - elisirUsato2.value);
const danniFatti1 = ref(rand(0, 10000));
const danniFatti2 = ref(rand(0, 10000));
const tipoTorre1 = ref(tipoTorri[rand(0, 2)]);
const tipoTorre2 = ref(tipoTorri[rand(0, 2)]);

async function insertPartita() {
  try {
    const res = await fetch("/api/partita", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        vincitore: vincitore.value,
        orario: orario.value,
        tempo: tempo.value,
        evento: evento.value,
        giocatore1: giocatore1.value,
        giocatore2: giocatore2.value,
        mazzo1: mazzo1.value,
        mazzo2: mazzo2.value,
        elisirUsato1: elisirUsato1.value,
        elisirUsato2: elisirUsato2.value,
        elisirSprecato1: elisirSprecato1.value,
        elisirSprecato2: elisirSprecato2.value,
        danniFatti1: danniFatti1.value,
        danniFatti2: danniFatti2.value,
        tipoTorri1: tipoTorre1.value,
        tipoTorri2: tipoTorre2.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento dell'partita");
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
    <h3 class="font-bold text-xl">Insert partita</h3>

    <form class="flex flex-col w-full gap-2" @submit.prevent="">
      <div>
        <div class="label">
          <span class="label-text">Vincitore</span>
        </div>
        <label class="label cursor-pointer">
          <span class="label-text">{{
            vincitore ? "Giocatore 1" : "Giocatore 2"
          }}</span>
          <input type="checkbox" v-model="vincitore" class="checkbox" />
        </label>
        <div></div>
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Orario</span>
        </div>
        <input
          type="time"
          placeholder="Luogo"
          v-model="orario"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Tempo (secondi)</span>
        </div>
        <input
          type="number"
          placeholder="125"
          v-model.number="tempo"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Nome Evento</span>
        </div>
        <input
          type="text"
          placeholder="Nome evento"
          v-model="evento"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Username giocatore 1</span>
        </div>
        <input
          type="text"
          placeholder="Username"
          v-model="giocatore1"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Username giocatore 2</span>
        </div>
        <input
          type="text"
          placeholder="Username"
          v-model="giocatore2"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Mazzo giocatore 1</span>
        </div>
        <input
          type="number"
          v-model="mazzo1"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Mazzo giocatore 2</span>
        </div>
        <input
          type="number"
          v-model="mazzo2"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Elisir usato giocatore 1</span>
        </div>
        <input
          type="number"
          v-model="elisirUsato1"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Elisir usato giocatore 2</span>
        </div>
        <input
          type="number"
          v-model="elisirUsato2"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Elisir precato giocatore 1</span>
        </div>
        <input
          type="number"
          v-model="elisirSprecato1"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Elisir precato giocatore 2</span>
        </div>
        <input
          type="number"
          v-model="elisirSprecato2"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Danni fatti giocatore 1</span>
        </div>
        <input
          type="number"
          v-model="danniFatti1"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Danni fatti giocatore 2</span>
        </div>
        <input
          type="number"
          v-model="danniFatti2"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Tipo torri giocatore 1</span>
        </div>
        <select
          v-model="tipoTorre1"
          class="select select-bordered select-primary w-full"
        >
          <option v-for="tipo in tipoTorri" :key="tipo" :value="tipo">
            {{ tipo }}
          </option>
        </select>
      </div>

      <div class="form-control">
        <div class="label">
          <span class="label-text">Tipo torri giocatore 2</span>
        </div>
        <select
          v-model="tipoTorre2"
          class="select select-bordered select-primary w-full"
        >
          <option v-for="tipo in tipoTorri" :key="tipo" :value="tipo">
            {{ tipo }}
          </option>
        </select>
      </div>

      <button @click="insertPartita" class="btn btn-primary mt-4">
        Insert partita
      </button>
    </form>
  </div>
</template>

<style></style>
