<script setup lang="ts">
import OP01 from "@/components/InsertGiocatore.vue";
import OP02 from "@/components/FormareTeam.vue";
import OP03 from "@/components/InsertCoach.vue";
import OP04 from "@/components/InsertSponsor.vue";
import OP05 from "@/components/InsertCampionato.vue";
import OP06 from "@/components/InsertEvento.vue";
import OP07 from "@/components/InsertPartita.vue";
import OP08 from "@/components/InsertMazzo.vue";
import OP09 from "@/components/InsertCarta.vue";
import OP010 from "@/components/InsertBiglietto.vue";
import NoAction from "@/components/NoAction.vue";
import { ref, computed, watch, onMounted } from "vue";

const op = ref(0);

const comp = computed(() => {
  console.log(op.value);
  switch (op.value) {
    case 1:
      return OP01;
    case 2:
      return OP02;
    case 3:
      return OP03;
    case 4:
      return OP04;
    case 5:
      return OP05;
    case 6:
      return OP06;
    case 7:
      return OP07;
    case 8:
      return OP08;
    case 9:
      return OP09;
    case 10:
      return OP010;
    default:
      return NoAction;
  }
});

watch(op, (newVal) => {
  localStorage.setItem("op", newVal.toString());
});

onMounted(() => {
  const savedOp = localStorage.getItem("op");
  if (savedOp) {
    op.value = parseInt(savedOp);
  }
});
</script>

<template>
  <div class="p-2 flex flex-col items-center gap-8 w-full">
    <header class="w-full bg-base-300 rounded-lg">
      <RouterLink :to="'/'">
        <h1 class="text-center font-bold text-2xl my-4">Torneo Clash Royale</h1>
      </RouterLink>
    </header>
    <main class="flex flex-col gap-4 w-full max-w-lg">
      <div class="">
        <select
          v-model.number="op"
          class="select select-bordered select-primary w-full"
        >
          <option value="0">Seleziona operazione</option>
          <option value="1">Inserire Giocatore</option>
          <option value="2">Fomare un team</option>
          <option value="3">Inserire Coach</option>
          <option value="4">Inserire Sponsor</option>
          <option value="5">Inserire Campionato</option>
          <option value="6">Inserire Evento</option>
          <option value="7">Inserire Partita</option>
          <option value="8">Inserire Mazzo</option>
          <option value="9">Inserire Carta</option>
          <option value="10">Erogare Biglietto</option>
        </select>
      </div>
      <div class="shadow-lg bg-base-200 flex justify-center p-4 rounded-lg">
        <component :is="comp" class="max-w-md" />
      </div>
    </main>
  </div>
</template>
