<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const nome = ref("");
const cognome = ref("");
const dataNascita = ref("");
const luogoNascita = ref("");
const team = ref("");
const salario = ref("");

async function insertCoach() {
  try {
    const res = await fetch("/api/coach", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        cognome: cognome.value,
        dataNascita: new Date(dataNascita.value).toISOString(),
        luogoNascita: luogoNascita.value,
        team: team.value,
        salario: salario.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del coach");
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
    <h3 class="font-bold text-xl">Insert coach</h3>

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
          <span class="label-text">Cognome</span>
        </div>
        <input
          type="text"
          placeholder="Cognome"
          v-model="cognome"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Data di nascita</span>
        </div>
        <input
          type="date"
          placeholder="Data di nascita"
          v-model="dataNascita"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Luogo di nascita</span>
        </div>
        <input
          type="text"
          placeholder="Luogo di nascita"
          v-model="luogoNascita"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Nome Team</span>
        </div>
        <input
          type="text"
          placeholder="Nome Team"
          v-model="team"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Salario</span>
        </div>
        <input
          type="number"
          placeholder="Dig"
          v-model="salario"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="insertCoach" class="btn btn-primary mt-4">
        Insert coach
      </button>
    </form>
  </div>
</template>

<style></style>
