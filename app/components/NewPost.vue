<template>
  <div class="popup flex-h flex-v col">
    <div class="bg" @click="changePopupState()"></div>
    <div class="contents flex-v col">
      <div class="popup-header flex-h flex-v">New Post</div>
      <input type="text" placeholder="title" class="popup-input">
      <textarea class="text flex-v" v-model="content" placeholder="contents of post-it"></textarea>
      <button type="submit">post-it!</button>
    </div>
  </div>
</template>

<style src="~/styles/popup.css"></style>

<script setup>
const emit = defineEmits(['popup']);
function changePopupState() {
  emit('popup');
}
</script>

