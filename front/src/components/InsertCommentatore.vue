<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const persona = ref(0);
const partita = ref(0);
const lingua = ref("Italiano");

async function insertCommenta() {
  try {
    const res = await fetch("/api/commenta", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        persona: persona.value,
        partita: partita.value,
        lingua: lingua.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del giocatore");
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
    <h3 class="font-bold text-xl">Insert giocatore</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div>
        <div class="label">
          <span class="label-text">Persona (ID)</span>
        </div>
        <input
          type="number"
          v-model.number="persona"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Partita (ID)</span>
        </div>
        <input
          type="number"
          v-model.number="partita"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Lingua</span>
        </div>
        <input
          type="text"
          v-model="lingua"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="insertCommenta" class="btn btn-primary mt-4">
        Insert commenta
      </button>
    </form>
  </div>
</template>

<style></style>
