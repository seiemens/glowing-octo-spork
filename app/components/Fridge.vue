<!-- TODO: IMPLEMENT POPUP   -->
<template>
  <div id="fridge">
    <div class="fridge-head flex-h flex-v col">
      <span class="title">{{ type }} post-its</span>

      <div class="buttons">
        <input type="text" placeholder="search a post-it!" id="filter" @input="filter()">
        <button @click="newPost = !newPost" class="add-btn" v-if="type != 'public'">new post-it</button>
      </div>
      <span class="handle"></span>
    </div>
    <div class="divider"></div>
    <div class="section flex-h">
      <span class="handle-2"></span>
      <PostIt v-for="p in displayedPosts"
        :number="p.id"
        :title="p.title"
        :content="p.content"
        :author="p.author"
        @popup="popup = p"
      />
    </div>
    <div class="bottom"></div>
    <Popup v-if="popup != undefined" @popup="popup = undefined"
      :id="popup.id"
      :title="popup.title"
      :content="popup.content"
      :author="popup.author"
      :comments="popup.comments"
      :userid="popup.userid"
      :status="popup.status"
      :isOwner="checkOwner(popup.userid)"
    />
    <NewPost v-if="newPost" @popup="newPost = !newPost"/>
  </div>
</template>

<style src="~/styles/fridge.css"></style>

<script setup>
const props = defineProps({
  type: {
    type: String,
    required: true
  },
  posts: {
    type: Array,
    required: true
  },
  isAdmin: Boolean
});

const displayedPosts = ref(props.posts);
const popup = ref(undefined);
const newPost = ref(false);
const uid = await getUserID();


function filter() {
  const search = document.getElementById('filter').value.toLowerCase();
  displayedPosts.value = [];

  props.posts.forEach(element => {
    if (element.title.toLowerCase().includes(search)) {
      displayedPosts.value.push(element);
    }
  });
}

function checkOwner(id) {
  console.log(uid + " | " + id);
  // only show the visibility checkbox if its not on the public page
  if (((props.isAdmin) | (id == uid)) && props.type != "public") {
    return true;
  }
  else {
    return false;
  }
}
</script>