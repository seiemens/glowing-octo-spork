<template>
  <div class="postit" @click="changePopupState()">
    <div class="header"> {{ title }}</div>
    <div class="content">{{ content }}</div>
    <div class="footer">by {{ author }}</div>
  </div>
</template>


<!-- Only Inline CSS because i had to use reactive properties for the rotation and margin! -->
<style scoped>

.header {
  width: 100%;
  height: 4vh;

  margin: auto 0;
  background-color: var(--top);

  text-align: center;
  font-family: 'Pacifico', cursive;
  font-size: 2vh;
  
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
}

.footer {
  position: absolute;
  bottom: 0;
  right:1vh;
  font-size: 1.5vh;

}

.content {
  width: 90%;
  margin: 0 auto;
  line-height: 2.25vh;

  display: -webkit-box;
  -webkit-line-clamp: 4;
  -webkit-box-orient: vertical; 

  word-wrap: break-word;
  text-overflow: ellipsis;
  overflow: hidden;
}

.postit {
  position: relative;
  width: 20vh;
  height: 20vh;
  overflow: hidden;

  font-size: 1.5vh;
  text-overflow: ellipsis;
  /* border: 0.1vh solid black; */

  background-color: var(--base);
  font-family: 'Gloria Hallelujah', cursive;

  transform: rotateZ(v-bind(rotation));
  margin: v-bind(margin);

  transition: 0.2s;
}

.postit:hover {
  cursor: pointer;
  transform: translate(-1vh, -1vh) rotateZ(v-bind(rotation));
  box-shadow: 1vh 1vh rgba(0, 0, 0, 0.5);
}

</style>


<script setup>
const MAX_ROTATION = 20; // in degree
const MAX_MARGIN = 7; // in vh

const props = defineProps({
  title: String,
  content: String,
  author: String,
});

const emit = defineEmits(['popup']);

const rotation = `${(Math.round(Math.random() * MAX_ROTATION * (Math.random() > 0.5 ? -1 : 1)))}deg`;
const margin = `${Math.round(Math.random() * MAX_MARGIN)}vh`;


function changePopupState() {
  emit('popup');
}

</script>