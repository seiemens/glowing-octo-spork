<template>
  <div class="head flex-v">
    <a class="logo" href="/">FRIDGE IT!</a>
    <div class="nav flex-v">
      <a href="/login" v-if="!authed">login</a>
      <div v-if="authed">
        <a href="/dashboard">dashboard</a>
        <a href="/profile">profile</a>
        <a href="#" @click="endSession()">logout</a>
      </div>
    </div>
  </div>
</template>

<style src="~/styles/header.css" scoped></style>

<script setup>
const authed = ref(auth());

async function endSession() {
  await $fetch('http://localhost:8080/api/user/logout', {credentials:'include'}).then((res)=>{
    navigateTo('/login');
  });
}
</script>