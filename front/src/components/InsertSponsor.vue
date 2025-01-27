<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const $toast = useToast();

const nome = ref("");
const team = ref("");
const budget = ref("");

async function insertSponsor() {
  try {
    const res = await fetch("/api/sponsor", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        team: team.value,
        budget: budget.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del sponsor");
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
    <h3 class="font-bold text-xl">Insert sponsor</h3>

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
          <span class="label-text">Nome Team</span>
        </div>
        <input
          type="text"
          placeholder="Cognome"
          v-model="team"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Budget</span>
        </div>
        <input
          type="date"
          placeholder="Data di nascita"
          v-model="budget"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="insertSponsor" class="btn btn-primary mt-4">
        Insert sponsor
      </button>
    </form>
  </div>
</template>

<style></style>
