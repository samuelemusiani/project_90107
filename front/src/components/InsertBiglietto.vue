<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";
import { rand } from "@/utils";

const $toast = useToast();

const id = ref(0);
const prezzo = ref(rand(10, 20));
const posto = ref(rand(1, 200));
const persona = ref(0);
const evento = ref(0);

async function createBiglietto() {
  try {
    const res = await fetch("/api/biglietto", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        prezzo: prezzo.value,
        posto: posto.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del biglietto");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    id.value = data.id;
  } catch (err) {
    console.error(err);
  }
}

async function createAssiste() {
  try {
    const res = await fetch("/api/assiste", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        prezzo: prezzo.value,
        posto: posto.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del biglietto");
      throw new Error("Failed to insert user");
    }

    const data = await res.json();
    id.value = data.id;
  } catch (err) {
    console.error(err);
  }
}

onMounted(() => {});
</script>

<template>
  <div class="greetings w-full">
    <h3 class="font-bold text-xl">Create biglietto</h3>

    <form class="flex flex-col w-full gap-2" @click.prevent>
      <div class="form-control">
        <label class="label">
          <span class="label-text">Prezzo</span>
        </label>
        <input
          type="number"
          v-model="prezzo"
          class="input input-bordered w-full"
        />
      </div>

      <div class="form-control">
        <label class="label">
          <span class="label-text">Posto</span>
        </label>
        <input
          type="number"
          v-model="posto"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="createBiglietto" class="btn btn-secondary mt-4">
        Crea Biglietto
      </button>
    </form>

    <h3 class="font-bold text-xl mt-8">Create Assiste</h3>
    <form class="flex flex-col w-full gap-2" @click.prevent="">
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
          <span class="label-text">Evento (ID)</span>
        </div>
        <input
          type="number"
          v-model.number="evento"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Biglietto (ID)</span>
        </div>
        <input
          type="number"
          v-model.number="id"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="createAssiste" class="btn btn-primary mt-4">
        Eroga biglietto per persona ed evento
      </button>
    </form>
  </div>
</template>

<style></style>
