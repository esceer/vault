<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
    <nav>
        <ul>
            <li>
                <img class="logo" src="@/assets/home.svg" />
                <RouterLink to="/">Home</RouterLink>
            </li>
            <li>
                <img class="logo" src="@/assets/list.png" />
                <RouterLink to="/list">List</RouterLink>
            </li>
            <li>
                <img class="logo" src="@/assets/plussign.png" />
                <RouterLink to="/create">New</RouterLink>
            </li>
            <li>
                <img class="logo" src="@/assets/search.svg" />
                <RouterLink to="/search">Search</RouterLink>
            </li>
            <li>
                <img class="logo" src="@/assets/shield.png" />
                <RouterLink to="/security">Security</RouterLink>
            </li>
        </ul>
    </nav>
    <footer>
        <img class="logo" src="@/assets/info.svg" />
        <a href="https://github.com/esceer/vault">Github</a>
    </footer>
</template>

<style scoped>
li {
    align-items: center;
    display: flex;
    line-height: 1.7rem;
    list-style-type: none;
}

.logo {
    height: 20px;
    margin-right: 7px;
    padding: 1px;
    width: 20px;
}

footer {
    align-items: center;
    display: flex;
    margin-top: 25px;
    padding-left: 40px;
}
</style>