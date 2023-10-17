<script setup>
import { RouterView } from 'vue-router'
import Menu from './components/Menu.vue';
</script>

<template>
  <aside>
    <header>
      <img class="logo" src="@/assets/img/logo.svg" />
      Safe
    </header>
    <Menu></Menu>
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

main {
  background-color: var(--color-background-secondary);
  color: var(--vt-c-black-soft);
  flex: 3.5;
}

.logo {
  width: 75px;
  height: 75px;
}
</style>