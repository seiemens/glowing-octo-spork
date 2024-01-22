<template>
  <div class="container fc flex-h flex-v col">
    <Fridge :type="isAdmin ? 'all' : 'private'" :posts="posts" :isAdmin="isAdmin"/>
  </div>
</template>

<script setup>
const isAdmin = ref(false);

const posts = await $fetch('http://localhost:8080/api/posts/get',{
  method:'post',
  credentials:'include',
  body: JSON.stringify({
    mode: isAdmin.value ? 'admin' : 'user'
  })
}).then((res)=>{
  return res.answer;
}).catch((err)=>{
  console.log(err);
});
</script>