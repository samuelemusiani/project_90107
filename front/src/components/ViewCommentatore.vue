<script setup lang="ts">
import { formatDate } from "@/utils";
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const response = ref(false);
const partita = ref(0);

const nome = ref("");
const cognome = ref("");
const dataNascita = ref("");
const luogoNascita = ref("");

async function viewCommentatore() {
  try {
    const res = await fetch(`/api/commentatore/${partita.value}`, {
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
      nome.value = data.nome;
      cognome.value = data.cognome;
      dataNascita.value = formatDate(new Date(data.dataNascita));
      luogoNascita.value = data.luogoNascita;
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
    <h3 class="font-bold text-xl">Visualizza il commentatore di una partita</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Id partita</span>
        </div>
        <input
          type="number"
          placeholder="Nome"
          v-model="partita"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="viewCommentatore" class="btn btn-primary mt-4">
        Insert coach
      </button>
    </form>

    <div v-if="response">
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
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Cognome</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="cognome"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Data di nascita</span>
          </div>
          <input
            type="date"
            placeholder="Nome"
            v-model="dataNascita"
            class="input input-bordered w-full"
            disabled
          />
        </div>
        <div>
          <div class="label">
            <span class="label-text">Luogo di nascita</span>
          </div>
          <input
            type="text"
            placeholder="Nome"
            v-model="luogoNascita"
            class="input input-bordered w-full"
            disabled
          />
        </div>
      </form>
    </div>
  </div>
</template>

<style></style>
