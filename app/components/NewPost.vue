<template>
  <div class="popup flex-h flex-v col">
    <div class="bg" @click="changePopupState()"></div>
    <form class="contents flex-v col" @submit.prevent="submitForm()">
      <div class="popup-header flex-h flex-v">New Post</div>
      <input type="text" placeholder="title" class="popup-input" v-model="title">
      <textarea class="text flex-v" v-model="content" placeholder="contents of post-it"></textarea>
      <button type="submit">post-it!</button>
    </form>
  </div>
</template>

<style src="~/styles/popup.css"></style>

<script setup>
const title = ref('');
const content = ref('');

const emit = defineEmits(['popup']);
function changePopupState() {
  emit('popup');
}

async function submitForm() {
  const data = {
    title: title.value,
    content: content.value
  };

  await $fetch('http://localhost:8080/api/posts/create',{
    method:'post',
    credentials:"include",
    body: JSON.stringify(data)
  }).then((res)=>{
    console.log(res);
  });
}
</script>

