<template>
  <div class="container fc flex-h">
    <Fridge :type="isAdmin ? 'all' : 'private'" :posts="posts" :isAdmin="isAdmin"/>
  </div>
</template>

<script setup>
const isAdmin = ref(false);

const posts = [
  {
    "id": 1,
    "title": "Sunny Days",
    "content": "Just found the perfect spot for my post-it notes on the fridge. Ready for a colorful summer!",
    "author": "SunshineLover",
    "userID": "12345678901234",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 2,
    "title": "Recipe Reminders",
    "content": "My fridge is covered in post-its with all my favorite recipes. No more forgetting ingredients!",
    "author": "ChefAtHeart",
    "userID": "98765432109876",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 3,
    "title": "Artistic Chaos",
    "content": "Turning my fridge into an art gallery with post-it doodles. Embracing the creative mess!",
    "author": "DoodleMaster",
    "userID": "54321098765432",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 4,
    "title": "Workout Motivation",
    "content": "Every morning, a new post-it with a motivational quote on the fridge. Let's crush those fitness goals!",
    "author": "FitLifeChampion",
    "userID": "13579246802468",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 5,
    "title": "Daily Gratitude",
    "content": "Started a gratitude post-it ritual on my fridge. It's amazing how a little note can change your perspective!",
    "author": "GratefulSoul",
    "userID": "24680135791234",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 6,
    "title": "Techie Notes",
    "content": "My fridge is now a tech hub with post-its filled with coding snippets and tech reminders.",
    "author": "CodeGeek",
    "userID": "98765432123456",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 7,
    "title": "Movie Night Picks",
    "content": "Keeping track of must-watch movies with post-its on the fridge. Movie night just got more organized!",
    "author": "Cinephile",
    "userID": "87654321098765",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 8,
    "title": "Travel Bucket List",
    "content": "Dreaming of exotic destinations? Post-it notes on the fridge are turning my travel dreams into plans!",
    "author": "WanderlustDreamer",
    "userID": "34567890123456",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 9,
    "title": "Bookworm's Corner",
    "content": "Transformed my fridge into a book recommendations space. Can't wait to explore new literary worlds!",
    "author": "BookishExplorer",
    "userID": "65432109876543",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  },
  {
    "id": 10,
    "title": "Daily Affirmations",
    "content": "Starting the day right with positive affirmations on the fridge. Today is going to be amazing!",
    "author": "PositivityFan",
    "userID": "12345098765432",
    "status": "visible",
    "comments": [
      { "author": "Commenter1", "comment": "This is such a smart way to manage a family schedule!" },
      { "author": "Commenter2", "comment": "I need to do this too. Family events can get overwhelming!" },
      { "author": "Commenter3", "comment": "Your organizational skills are impressive. Any tips?" },
      { "author": "Commenter4", "comment": "No more missed events for sure. Thanks for the idea!" },
      { "author": "Commenter5", "comment": "I bet your fridge is the family command center now!" },
    ],
  }
];
     
</script>