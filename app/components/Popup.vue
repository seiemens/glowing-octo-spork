<template>
  <div class="popup flex-h flex-v col">
    <div class="bg" @click="changePopupState()"></div>
    <div class="contents">
      <div class="popup-header flex-h flex-v">{{ title }}</div>
      <div class="text flex-v">{{ content }}</div>
      <div class="author">written by {{ author }}</div>

      <div class="select">
        <select name="visibility" v-model="postVisibility" v-if="isOwner">
          <option value="1">published</option>
          <option value="2">hidden</option>
          <option value="3">deleted</option>
        </select>
      </div>
      
      <div class="open-btn flex-v flex-h col">
        <span>Comments</span>
        <a @click="changeCommentState()">↫</a>
      </div>
    </div>
    <div class="comments" :id="id">
      <form class="comment flex-v" @submit.prevent="submitForm()">
        <input class="popup-input ci" placeholder="your comment..." v-model="comment"/>
        <button type="submit" class="ci">➤</button>
      </form>
      <div class="comment flex-v" v-for="c in comments">
      <div class="author">from {{ c.author }}</div>
      <div class="text">{{ c.content }}</div>
      </div>
    </div>
  </div>
</template>

<style src="~/styles/popup.css"></style>

<style scoped>
.comments {
  position: relative;
  width: 80vh;
  margin-top: v-bind(commentsVisible);
  max-height: 30vh;
  z-index: 15;
  overflow-y: scroll;
  transition: 0.25s;
}

.open-btn {
  position: absolute;
  bottom: 0;
  width: 100%;
  font-family: 'Pacifico', cursive;
}
.open-btn a {
  font-size: 4vh;
  transform: rotateZ(v-bind(buttonRotation));
  transition: 0.25s;
}
.open-btn a:hover {
  cursor: pointer;
  scale: 1.2;
  color: var(--body);
}

.open-btn span {
  font-size: 2.5vh;
  margin-bottom: -2vh;
}
</style>

<script setup>
const props = defineProps({
  id: String,
  title: String,
  content: String,
  author: String,
  userid: String,
  comments: Array,
  status: Number,

  isOwner: Boolean
});
const postVisibility = ref(props.status);
const commentsVisible = ref("-31vh");
const buttonRotation = ref("-90deg");

const comment = ref('');


const emit = defineEmits(['popup']);
function changePopupState() {
  emit('popup');
}

function changeCommentState() {
  commentsVisible.value = commentsVisible.value == "1vh" ? "-31vh" : "1vh";
  buttonRotation.value = buttonRotation.value == "-90deg" ? "90deg" : "-90deg";
}

async function submitForm() {
  if (await auth() == false) {
    alert('you have to sign in to write a comment.');
    navigateTo('/login');
  }

  const data = {
    postid:props.id,
    content: comment.value.substring(0,200)
  };

  await $fetch('http://localhost:8080/api/posts/comment', {
    method:'post',
    credentials:'include',
    body: JSON.stringify(data)
  }).then((res)=>{
    
  }).catch((err)=>{
  });
}

watch(postVisibility, async ()=>{
  const data = {
    id: props.id,
    userid: props.userid,
    status: Number.parseInt(postVisibility.value)
  };

  await $fetch('http://localhost:8080/api/posts/visibility', {
    method:'post',
    credentials:'include',
    body: JSON.stringify(data)
  }).then((res)=>{
    reloadNuxtApp();
  }).catch((err)=>{
    console.log(err);
  });
});
</script>