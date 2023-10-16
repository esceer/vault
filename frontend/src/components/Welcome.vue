<template>
    <div class="container">
        <img src="../assets/logo.svg">
        <div class="text">
            <h2>Welcome to Vault, your personal secret assistant!</h2>
            <p>You can securely store passwords, pins, notes or all sorts of secrets you want to hide from the public.</p>
            <p>Vault keeps everything <span>isolated</span> for you.</p>
        </div>
    </div>
</template>

<style scoped>
.container {
    align-items: center;
    display: flex;
}

img {
    width: 300px;
}

h2, span{
    color: var(--dusky-green);
}
</style>