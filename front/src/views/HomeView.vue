<script setup lang="ts">
import OP01 from "@/components/InsertGiocatore.vue";
import OP02 from "@/components/FormareTeam.vue";
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
          <option value="1">Inserisci Giocatore</option>
          <option value="2">Formare un team</option>
        </select>
      </div>
      <div class="shadow-lg bg-base-200 flex justify-center p-4 rounded-lg">
        <component :is="comp" class="max-w-md" />
      </div>
    </main>
  </div>
</template>
