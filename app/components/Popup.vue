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
      
      <a class="open-btn" @click="changeCommentState()">↫</a>
    </div>
    <div class="comments" :id="id">
      <div class="comment flex-v" v-for="c in comments">
      <div class="author">from {{ c.author }}</div>
      <div class="text">{{ c.comment }}</div>
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
  left: 50%;
  font-weight:bold;
  font-size: 4vh;
  transform: rotateZ(v-bind(buttonRotation));
  transition: 0.25s;
}

.open-btn:hover {
  cursor: pointer;
  scale: 1.2;
  color: var(--body);
}
</style>

<script setup>
const props = defineProps({
  id: Number,
  title: String,
  content: String,
  author: String,
  userID: String,
  comments: Array,
  status: Number,

  isOwner: Boolean
});
const postVisibility = ref(props.status);

const commentsVisible = ref("-31vh");
const buttonRotation = ref("-90deg");

const emit = defineEmits(['popup']);
function changePopupState() {
  emit('popup');
}

function changeCommentState() {
  commentsVisible.value = commentsVisible.value == "1vh" ? "-31vh" : "1vh";
  buttonRotation.value = buttonRotation.value == "-90deg" ? "90deg" : "-90deg";
}
</script>