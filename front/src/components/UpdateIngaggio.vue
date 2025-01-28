<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const team = ref("");
const giocatore = ref("");
const salario = ref(0);

async function updateIngaggio() {
  try {
    const res = await fetch("/api/ingaggio", {
      method: "PUT",
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
      $toast.error("Errore durante l'update dell'ingaggio");
      throw new Error("Failed to update ingaggio");
    }
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full">
    <div class="flex flex-col gap-2">
      <h3 class="font-bold text-xl">Aggiorna ingaggio</h3>

      <form class="flex flex-col gap-2" @click.prevent>
        <div>
          <div class="label">
            <span class="label-text">Nome nuovo team</span>
          </div>
          <input
            type="text"
            placeholder="Team"
            v-model="team"
            class="input input-bordered w-full"
          />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Username giocatore</span>
          </div>
          <input
            type="text"
            placeholder="Giocatore"
            v-model="giocatore"
            class="input input-bordered w-full"
          />
        </div>

        <div>
          <div class="label">
            <span class="label-text">Nuovo salario</span>
          </div>
          <input
            type="number"
            placeholder="Salario"
            v-model="salario"
            class="input input-bordered w-full"
          />
        </div>

        <button @click="updateIngaggio" class="btn btn-primary mt-2">
          Aggiorna ingaggio
        </button>
      </form>
    </div>
  </div>
</template>

<style></style>
