<template>
  <div class="flex-h flex-v row">
    <div class="note">
      <div class="popup-header flex-h flex-v">{{ title }}</div>
      <div class="text flex-v">{{ x.content }}</div>
      <div class="author">written by {{ author }}</div>
      <div class="arrow"></div>
      <div class="select">
        <select name="visibility" v-model="postVisibility" v-if="isOwner">
          <option value="1">published</option>
          <option value="2">hidden</option>
          <option value="3">deleted</option>
        </select>
      </div>
    </div>
    <div class="comment-section flex-v col">
      <div class="input-field flex-h">
        <input type="text" placeholder="your comment" class="comment">
        <button @click="uploadComment()" class="comment">post</button>
      </div>
      <div class="comments" :id="id">
        <div class="comment flex-v" v-for="c in x.comments">
        <div class="author">from {{ c.author }}</div>
        <div class="text">{{ c.comment }}</div>
      </div>
    </div>
    </div>
  </div>
</template>

<style src="~/styles/detailed.css"></style>

<script setup>
const x = {
  "id": 1,
  "title": "Sunny Days",
  "content": "Just found the perfect spot for my post-it notes on the fridge. Ready for a colorful summer!",
  "author": "SunshineLover",
  "userID": "12345678901234",
  "status": 1,
  "comments": [
    { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
    { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
    { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
    { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
    { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
    { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
    { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
    { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
    { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
    { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
    { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
    { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
    { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
  ],
};


// const post = await $fetch('/api/post/',{
//   method: 'post',
//   credentials: 'include',
//   body: JSON.stringify({
//     id
//   })
// }).then(async (res)=>{
//   await res.json();
//   return res;
// });


</script>