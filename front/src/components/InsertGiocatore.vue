<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from 'vue-toast-notification'
import 'vue-toast-notification/dist/theme-sugar.css';

const $toast = useToast()

const nome = ref("");
const cognome = ref("");
const dataNascita = ref("");
const luogoNascita = ref("");
const username = ref("");
const dig = ref("");

async function insertGiocatore() {
  try {
    const res = await fetch("/api/giocatore", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        cognome: cognome.value,
        dataNascita: new Date(dataNascita.value).toISOString(),
        luogoNascita: luogoNascita.value,
        username: username.value,
        dig: new Date(dig.value).toISOString(),
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
    <h3 class="font-bold text-xl">Insert giocatore</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Nome</span>
        </div>
        <input type="text" placeholder="Nome" v-model="nome" class="input input-bordered w-full" />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Cognome</span>
        </div>
        <input type="text" placeholder="Cognome" v-model="cognome" class="input input-bordered w-full" />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Data di nascita</span>
        </div>
        <input type="date" placeholder="Data di nascita" v-model="dataNascita" class="input input-bordered w-full" />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Luogo di nascita</span>
        </div>
        <input type="text" placeholder="Luogo di nascita" v-model="luogoNascita" class="input input-bordered w-full" />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Username</span>
        </div>
        <input type="text" placeholder="Username" v-model="username" class="input input-bordered w-full" />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Data di inzio gioco</span>
        </div>
        <input type="date" placeholder="Dig" v-model="dig" class="input input-bordered w-full" />
      </div>

      <button @click="insertGiocatore" class="btn btn-primary">Insert user</button>
    </form>
  </div>
</template>

<style></style>
