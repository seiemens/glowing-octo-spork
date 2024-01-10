<template>
  <div class="container fh flex-v flex-h">
    <label for="mode" class="switcher">
      Login
      <input type="checkbox" name="mode" id="switch" @change="switchMode()">
      Register
    </label>
    <form class="flex-v flex-h col small-fridge" id="minifridge">
      <div class="title">{{ mode }}</div>
      <span class="handle"></span>
      <div class="inputs flex-v flex-h col">
        <input type="text" placeholder="your email">
        <input type="password" name="" id="" placeholder="your password">
        <button type="submit">{{mode}}&#x219D;</button>
      </div>
      <div class="shadow"></div>
    </form>
  </div>
</template>

<style src="~/styles/login.css" scoped></style>

<script setup>
const mode = ref('login');

function switchMode() {
  const fridge = document.getElementById('minifridge');
  const switcher = document.getElementById('switch').checked;

  if (!switcher) {
    fridge.classList.remove('register-mode');
    mode.value = 'login';
  }
  else {
    fridge.classList.add('register-mode');
    mode.value = 'register';
  }

}
</script>