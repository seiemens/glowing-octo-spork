<template>
  <div class="container fh flex-v flex-h">
    <div class="switch-container flex-v flex-h row">
      <span class="switch-text">
        log in
      </span>
      <label class="switch">
        <input type="checkbox" name="mode" id="switch" @change="switchMode()">
        <span class="slider"></span>
      </label>
      <span class="switch-text">
        register
      </span>
    </div>
    <form class="flex-v flex-h col small-fridge" id="minifridge" @submit.prevent="submitForm()">
      <div class="title">{{ mode }}</div>
      <span class="handle"></span>
      <div class="inputs flex-v flex-h col">
        <input type="text" v-model="mail" placeholder="your email" required>
        <input type="password" v-model="pw" placeholder="your password" required>
        <button type="submit">{{ mode }}&#x219D;</button>
      </div>
      <div class="shadow"></div>
    </form>
  </div>
</template>

<style src="~/styles/login.css" scoped></style>

<script setup>
const mode = ref('login');
const mail = ref('');
const pw = ref('');

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

function submitForm() {
  
}

</script>