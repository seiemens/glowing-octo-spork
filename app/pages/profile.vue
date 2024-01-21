<template>
  <div class="flex-h flex-v col fh">
    <div class="flex-h row">
      <div class="box">
        <div class="header">Details</div>
        <div class="infos flex-h col">
          <span><span class="small-header">ID: </span>{{ ud.id }}</span>
          <span><span class="small-header">Username: </span>{{ ud.username }}</span>
          <span><span class="small-header">Email: </span>{{ ud.email }}</span>
        </div>
      </div>
      <div class="box">
        <div class="header">Phone #</div>
        <div class="phone flex-h col">
          <div class="inputs">
            <input id="tel-input" type="tel" disabled v-model="tel"/>
            <button @click="changeInputState()">✎</button>
          </div>
          <button @click="changeTelNumber()">Change Phone Number</button>
        </div>
      </div>
    </div>
    <div class="flex-h row">
      <div class="box api flex-v flex-h col">
        <span class="header">Your API key:</span>
        <code>{{ ud.api_key }}</code>
      </div>
      <div class="box flex-h flex-v">
        <img src="~/public/vibin2cropped.gif" alt="">
      </div>
    </div>
  </div>
</template>

<style src="~/styles/profile.css"></style>

<script setup>
const ud = {
  id:"3290428244",
  username:"Kevin",
  email:"kevin@kevin.ch",
  phone: "435787654",
  api_key: "OWidajodiwajdioajd0329448720"
};
const tel = ref(ud.phone);

function changeInputState(optional) {
  const telInput = document.getElementById('tel-input');
  telInput.disabled = optional || !telInput.disabled;
}

function changeTelNumber() {
  console.log(ud.phone + " | " + tel.value);
  if (ud.phone != tel.value) {
    // do api things

    // lock the field again
    changeInputState(false);
  }
}
</script>