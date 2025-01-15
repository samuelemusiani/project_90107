<script setup lang="ts">
import { ref, onMounted } from 'vue'

const msg = ref('')
const users = ref<string[]>([])

async function fetchApiStatus() {
  const res = await fetch('/api')
  const data = await res.json()
  msg.value = data.message
}

async function fetchUsers() {
  const res = await fetch('/api/users')
  const data = await res.json()
  users.value = data.users
}

const newUser = ref('')
async function insertUser() {
  try {
    const res = await fetch('/api/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ user: newUser.value }),
    })

    if (!res.ok) {
      throw new Error('Failed to insert user')
    }

    await fetchUsers()
  } catch (err) {
    console.error(err)
  }
}

onMounted(() => {
  fetchApiStatus()
  fetchUsers()
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

    <div>
      <input type="text" placeholder="Enter username" v-model="newUser" />
      <button @click="insertUser">Insert user</button>
    </div>
    <h2>
      Users:
    </h2>
    <ul>
      <li v-for="user in users" :key="user">
        {{ user }}
      </li>
    </ul>
  </div>
</template>

<style scoped></style>
