<template>
  <div class="postit">
    <div class="header flex-v flex-h"> {{ title }}</div>
    {{ content }}
    <div class="footer">by {{ author }}</div>
  </div>
</template>


<!-- Only Inline CSS because i had to use reactive properties for the rotation and margin! -->
<style scoped>

.header {
  width: 100%;
  height: 20%;
  background-color: var(--top);

  font-family: 'Pacifico', cursive;
  font-size: 2vh;
  
  text-overflow: ellipsis;
}

.footer {
  position: absolute;
  bottom: 0;
  right:1vh;
}

.postit {
  position: relative;
  width: 20vh;
  height: 20vh;

  text-overflow: ellipsis;
  /* border: 0.1vh solid black; */

  background-color: var(--base);
  font-family: 'Gloria Hallelujah', cursive;

  transform: rotateZ(v-bind(rotation));
  margin: v-bind(margin);
}
</style>


<script setup>
const MAX_ROTATION = 20; // in degree
const MAX_MARGIN = 7; // in vh

const props = defineProps({
  title: String,
  content: String,
  author: String
});

const rotation = `${(Math.round(Math.random() * MAX_ROTATION * (Math.random() > 0.5 ? -1 : 1)))}deg`;
const margin = `${Math.round(Math.random() * MAX_MARGIN)}vh`;

</script>