<template>
  <div id="main" @mousemove="move()">
    <div id="blob"></div>
    <div class="loading flex-h flex-v col" v-if="loading">
      <span class="spinheader">loading...</span>
      <div class="spinfridge">
        <span class="spinhandle"></span>
      </div>
    </div>
    <Header/>
    <NuxtPage/>
  </div>
</template>

<!-- Styling of the Blob and Hover animation-->
<style>
#blob {
  position:fixed;
  top: 50%;
  left: 50%;
  translate: -50% -50%;
  height:1.5vh;
  width:1.5vh;
  background-color:white;
  border-radius:50%;
  
  transition: 0.25s;
  pointer-events: none;
  z-index:1;

  mix-blend-mode:exclusion;

}

</style>

<script setup>
const nuxtApp = useNuxtApp();
const loading = ref(true);
nuxtApp.hook("page:start", () => {
  loading.value = true;
});
nuxtApp.hook("page:finish", () => {
  setTimeout(() => {
    loading.value = false;
  }, 1800);
});

function move() {  
  document.getElementById('main').addEventListener('mouseover', (e)=>{
    let tracer = document.getElementById("blob");
    if (['A','INPUT','BUTTON', 'LABEL'].includes(e.target.tagName)) {
      tracer.style.scale = 0;
    }
    else {
      tracer.style.scale = 1;
    }
  });

  let tracer = document.getElementById("blob");
  let {clientX, clientY} = window.event;

  tracer.animate({
    left: `${clientX}px`,
    top: `${clientY}px`,
  },{duration:250, fill:"forwards"});
}

/* CLICKJACKING SHIT */
if (top != window) {
  top.location = window.location;
}

</script>
