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
  <div class="greetings">
    <h3 class="font-bold text-xl">Insert giocatore</h3>

    <form class="flex flex-col max-w-sm gap-2" @click.prevent>
      <input type="text" placeholder="Nome" v-model="nome" class="input input-bordered" />
      <input type="text" placeholder="Cognome" v-model="cognome" class="input input-bordered" />
      <input type="date" placeholder="Data di nascita" v-model="dataNascita" class="input input-bordered" />
      <input type="text" placeholder="Luogo di nascita" v-model="luogoNascita" class="input input-bordered" />
      <input type="text" placeholder="Username" v-model="username" class="input input-bordered" />
      <input type="date" placeholder="Dig" v-model="dig" class="input input-bordered" />

      <button @click="insertGiocatore" class="btn btn-primary">Insert user</button>
    </form>
  </div>
</template>

<style></style>
