<template>
  <div class="container fc flex-h flex-v col">
    <Fridge :type="isAdmin ? 'all' : 'your'" :posts="posts" :isAdmin="isAdmin"/>
  </div>
</template>

<script setup>
const isAdmin = ref(await checkUserRole());

const posts = await $fetch('http://localhost:8080/api/posts/get',{
  method:'post',
  credentials:'include',
  body: JSON.stringify({
    mode: 'user'
  })
}).then((res)=>{
  return res.answer;
}).catch((err)=>{
  navigateTo('/');
});

async function checkUserRole() {
  return await $fetch('http://localhost:8080/api/user/isadmin', {
    credentials:'include'
  }).then(()=>{
    return true;
  }).catch((err)=>{
    return false;
  });
}
</script>