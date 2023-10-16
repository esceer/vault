<script setup>
import Welcome from '../components/Welcome.vue'
</script>

<template>
  <main>
    <div class="welcome">
      <Welcome />
    </div>
    <h2>Let's get started!</h2>
    <p>Save your secrets. You will find it here whenever you need.</p>
    <div>
      <button class="filled">Add secret</button>
      <button class="framed">Search</button>
    </div>
  </main>
</template>

<style scoped>
main {
  align-items: center;
  display: flex;
  flex-direction: column;
}

.welcome {
  margin: 50px;
}

h2 {
  font-weight: bold;
}

p {
  color: var(--vt-c-black-mute);
}

button {
  border: 2px solid var(--dusky-green);
  border-radius: 5px;
  font-weight: bold;
  margin: 5px;
  min-width: 150px;
  padding: 8px 10px;
}

.filled {
  background-color: var(--dusky-green);
  color: var(--soft-white-green-accent);
}

.framed {
  background-color: var(--soft-white-green-accent);
  color: var(--dusky-green);
}
</style>