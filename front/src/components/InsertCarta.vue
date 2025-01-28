<script setup lang="ts">
import { ref, onMounted } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";
import { rand } from "@/utils";

const $toast = useToast();

const nome = ref("");
const elisir = ref(rand(1, 9));
const danni = ref(rand(100, 250) * elisir.value);

async function insertCarta() {
  try {
    const res = await fetch("/api/carta", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        nome: nome.value,
        elisir: elisir.value,
        danni: danni.value,
      }),
    });

    if (!res.ok) {
      $toast.error("Errore durante l'inserimento del carta");
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
    <h3 class="font-bold text-xl">Insert carta</h3>

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
          <span class="label-text">Elisir</span>
        </div>
        <input
          type="number"
          v-model="elisir"
          class="input input-bordered w-full"
        />
      </div>

      <div>
        <div class="label">
          <span class="label-text">Danni</span>
        </div>
        <input
          type="number"
          v-model="danni"
          class="input input-bordered w-full"
        />
      </div>

      <button @click="insertCarta" class="btn btn-primary mt-4">
        Insert carta
      </button>
    </form>
  </div>
</template>

<style></style>
