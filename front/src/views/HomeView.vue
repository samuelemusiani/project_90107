<script setup lang="ts">
import OP01 from "@/components/InsertGiocatore.vue"
import OP02 from "@/components/FormareTeam.vue"
import NoAction from "@/components/NoAction.vue"
import { ref, computed } from "vue";

const op = ref(0);

const comp = computed(() => {
  console.log(op.value)
  switch (op.value) {
    case 1:
      return OP01;
    case 2:
      return OP02;
    default:
      return NoAction;
  }
});
</script>

<template>
  <div class="p-2 flex flex-col items-center gap-2">
    <header>
      <h1 class="text-center font-bold text-2xl">Progetto BASI</h1>
    </header>
    <main class="flex flex-col gap-4">
      <div>
        <select v-model.number="op" class="select select-bordered w-full">
          <option value="0">Seleziona operazione</option>
          <option value="1">Inserisci Giocatore</option>
          <option value="2">Formare un team</option>
        </select>
      </div>
      <div>
        <component :is="comp" />
      </div>
    </main>
  </div>
</template>
