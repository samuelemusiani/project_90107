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
import OP011 from "@/components/InsertCommentatore.vue";
import OP12 from "@/components/ClassificaCampionato.vue";
import OP13 from "@/components/CartaUsataCampionato.vue";
import OP14 from "@/components/CartaUsataEvento.vue";
import OP15 from "@/components/CartaUsataPartita.vue";
import OP16_18 from "@/components/CarteUsate.vue";
import OP19_21 from "@/components/Statistiche.vue";
import OP22_25 from "@/components/MazzoUsato.vue";
import OP26_27 from "@/components/TorriUsate.vue";
import OP28_29 from "@/components/ViewBiglietti.vue";
import OP30 from "@/components/ViewCommentatore.vue";
import OP31 from "@/components/ViewTeam.vue";
import OP32 from "@/components/UpdateIngaggio.vue";
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
    case 11:
      return OP011;
    case 12:
      return OP12;
    case 13:
      return OP13;
    case 14:
      return OP14;
    case 15:
      return OP15;
    case 16:
      return OP16_18;
    case 17:
      return OP16_18;
    case 18:
      return OP16_18;
    case 19:
      return OP19_21;
    case 20:
      return OP19_21;
    case 21:
      return OP19_21;
    case 22:
      return OP22_25;
    case 23:
      return OP22_25;
    case 24:
      return OP22_25;
    case 25:
      return OP22_25;
    case 26:
      return OP26_27;
    case 27:
      return OP26_27;
    case 28:
      return OP28_29;
    case 29:
      return OP28_29;
    case 30:
      return OP30;
    case 31:
      return OP31;
    case 32:
      return OP32;
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
          <option value="11">Inserire Commentatore</option>
          <option value="12">Visualizza classifica campionato</option>
          <option value="13">Visualizza carta più usata campionato</option>
          <option value="14">Visualizza carta più usata evento</option>
          <option value="15">Visualizza carta più usata partita</option>
          <option value="16">Visualizza carte più usate campionato</option>
          <option value="17">Visualizza carte più usate evento</option>
          <option value="18">Visualizza carte più usate partita</option>
          <option value="19">Visualizza statistiche campionato</option>
          <option value="20">Visualizza statistiche evento</option>
          <option value="21">Visualizza statistiche partita</option>
          <option value="22">Visualizza mazzo più usato campionato</option>
          <option value="23">Visualizza mazzo più usato evento</option>
          <option value="24">Visualizza mazzo più efficace campionato</option>
          <option value="25">Visualizza mazzo più efficace evento</option>
          <option value="26">Visualizza torri più usate campionato</option>
          <option value="27">Visualizza torri più usate evento</option>
          <option value="28">
            Visualizza biglietti venduti per un campionato
          </option>
          <option value="29">Visualizza biglietti venduti per un evento</option>
          <option value="30">Visualizza commentatore di una partita</option>
          <option value="31">Visualizza team</option>
          <option value="32">Aggiorna ingaggio</option>
        </select>
      </div>
      <div class="shadow-lg bg-base-200 flex justify-center p-4 rounded-lg">
        <component :is="comp" class="max-w-md" :op="op" />
      </div>
    </main>
  </div>
</template>
