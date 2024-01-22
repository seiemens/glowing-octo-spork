<template>
  <div class="container fc flex-h">
    <Fridge type="public" :posts="posts"/>
  </div>
</template>

<script setup>

const posts = await $fetch('http://localhost:8080/api/posts/get',{
  method:'post',
  body: JSON.stringify({
    mode: 'public'
  })
}).then((res)=>{
  return res.answer;
});
</script>