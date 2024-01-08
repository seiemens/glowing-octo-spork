<!-- TODO: IMPLEMENT POPUP   -->

<template>
  <div id="fridge">
    <div class="fridge-head flex-h flex-v">
      <span class="title">{{ type }} post-its</span>
      <span class="handle"></span>
    </div>
    <div class="divider"></div>
    <div class="section flex-h">
      <span class="handle-2"></span>
      <PostIt v-for="p in posts" :key="p.id"
        :id="p.id"
        :title="p.title"
        :content="p.content"
        :author="p.author"
        @popup="openPopup"
      />
      />
    </div>
    <div class="bottom"></div>
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
  }
});
</script>