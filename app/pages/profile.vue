<template>
  <div class="flex-h flex-v col fh">
    <div class="flex-h row">
      <div class="box">
        <div class="header">Details</div>
        <div class="infos flex-h col">
          <span><span class="small-header">Username: </span>{{ userData.username }}</span>
          <span><span class="small-header">Email: </span>{{ userData.email }}</span>
        </div>
      </div>
      <div class="box">
        <div class="header">Phone #</div>
        <form class="phone flex-h col" @submit.prevent="changeTelNumber()">
          <div class="inputs">
            <input id="tel-input" type="text" pattern="[0-9]{12}" title="please enter a valid phone number (12 digits)" disabled v-model="tel" required/>
            <button @click="changeInputState()">✎</button>
          </div>
          <button type="submit">Change Phone Number</button>
        </form>
      </div>
    </div>
    <div class="flex-h row">
      <div class="box api flex-v flex-h col">
        <span class="header">Your API key:</span>
        <code>{{ userData.apiKey }}</code>
      </div>
      <div class="box flex-h flex-v">
        <img :src="`https://chart.googleapis.com/chart?cht=qr&chs=400x400&chl=${generateQR()}`" alt="TOTP QR" crossorigin>
      </div>
    </div>
  </div>
</template>

<style src="~/styles/profile.css"></style>

<script setup>
const userData = await $fetch('http://localhost:8080/api/user/info',{
  method:'get',
  credentials:'include'
}).then((res)=>{
  return res;
}).catch((err)=>{
  navigateTo('/');
});

const tel = ref(userData.phone);

function generateQR() {
    // a lot of extra stuff that isn't necessary but google needs it: https://github.com/google/google-authenticator/wiki/Key-Uri-Format
    return encodeURI(`otpauth://totp/M183:TOTP?secret=${userData.apiKey}&issuer=OOGA BOOGA INC&algorithm=SHA1`);
}

function changeInputState(optional) {
  const telInput = document.getElementById('tel-input');
  telInput.disabled = optional || !telInput.disabled;
}

async function changeTelNumber() {
  if (userData.phone != tel.value) {
    // do api things
    await $fetch('http://localhost:8080/api/user/telefon', {
      method:'post',
      credentials:'include',
      body: JSON.stringify({phone: tel.value})
    }).then((res)=>{
      console.log(res);
      alert('phone number changed!');
      reloadNuxtApp();
    }).catch((e)=>console.log(err));
    // lock the field again
    changeInputState(false);
  }
}
</script>