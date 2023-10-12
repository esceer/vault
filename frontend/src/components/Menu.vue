<script setup>
import { RouterLink } from 'vue-router'
</script>

<template>
    <nav>
        <ul>
            <li>
                <RouterLink to="/">Home</RouterLink>
            </li>
            <li>
                <RouterLink to="/list">List</RouterLink>
            </li>
            <li>
                <RouterLink to="/create">New</RouterLink>
            </li>
            <li>
                <RouterLink to="/search">Search</RouterLink>
            </li>
            <li>
                <RouterLink to="/security">Security</RouterLink>
            </li>
        </ul>
    </nav>
</template>

<style>
li {
    list-style-type: none;
}

li::before {
    content: "# "
}
</style>