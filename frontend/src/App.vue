<script setup>
import { RouterView } from 'vue-router'
import Menu from './components/Menu.vue';
</script>

<template>
  <aside>
    <header>
      <img class="logo" src="@/assets/logo.svg" width="50" height="50" />
      Safe
    </header>
    <Menu></Menu>
    <footer>
        <a href="https://github.com/esceer/vault">Github</a>
    </footer>
  </aside>

  <main>
    <RouterView />
  </main>
</template>

<style scoped>
aside {
  background-color: var(--color-background);
  color: var(--color-text);
  flex: 1;
}

header {
  border-bottom: 2px solid var(--color-border);
  font-size: 1.5rem;
  margin: 20px;
  padding-bottom: 20px;

  display: flex;
  justify-content: center;
  align-items: center;
  gap: 10px;
}

footer {
  margin-top: 25px;
  padding-left: 40px;
}

main {
  background-color: #cfe2df;
  color: var(--vt-c-black-soft);
  flex: 3.5;
}
</style>