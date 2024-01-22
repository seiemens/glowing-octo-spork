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
      <div class="small-section flex-h">
        <div class="title">{{ mode }}</div>
      </div>
      <span class="handle"></span>
      <div class="small-section">
        <div class="inputs flex-v flex-h col">
          <input type="text" v-model="username" placeholder="your username" required v-if="mode == 'register' | mode == 'login'">
          <input type="text" v-model="mail" placeholder="your email" required v-if="mode == 'register'">
          <input type="password" v-model="pw" placeholder="your password" required v-if="mode == 'register' | mode == 'login'"
          pattern="(?=.*\d)(?=.*[\W_])(?=.*[A-Z]).{8,}" title="Minimum of 8 characters. Should have at least one special, one numeric and one capital character."
          >
          <input type="text" placeholder="sms code" required v-if="mode == 'sms'">
          <button type="submit">{{ mode }}&#x219D;</button>
        </div>
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
const username = ref('');

function switchMode() {
  const fridge = document.getElementById('minifridge');
  const switcher = document.getElementById('switch').checked;

  fridge.classList.remove('sms-mode');

  if (!switcher) {
    fridge.classList.remove('register-mode');
    mode.value = 'login';
  }
  else {
    fridge.classList.add('register-mode');
    mode.value = 'register';
  }
}

async function submitForm() {
  let data = {
    username: username.value,
    password: pw.value,
    email: mail.value
  };

  console.log(data);
  await $fetch(`http://localhost:8080/api/user/${mode.value}`, {
    method: 'POST',
    credentials: "include",
    body: JSON.stringify(data)
  }).then((res)=> {
    console.log(res);
  });
  document.getElementById('minifridge').classList.add('sms-mode');
  document.getElementById('minifridge').classList.remove('register-mode');
  mode.value = 'sms';
}
</script>