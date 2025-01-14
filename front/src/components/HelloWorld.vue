<script setup lang="ts">
import { ref, onMounted } from 'vue'

const msg = ref('')

async function fetchApiStatus() {
  const res = await fetch('/api')
  const data = await res.json()
  msg.value = data.message
}

onMounted(() => {
  fetchApiStatus()
})
</script>

<template>
  <div class="greetings">
    <h3>
      You’ve successfully created a project with
    </h3>
    <p>
      Message from server: {{ msg }}
    </p>
  </div>
</template>

<style scoped></style>
