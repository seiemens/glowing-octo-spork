<template>
  <div class="container flex-h">
    <Fridge type="public" :posts="posts"/>
  </div>
</template>

<style src="~/styles/global.css"></style>

<script setup>
const posts = [
  {
    "id": 1,
    "title": "Sunny Days",
    "content": "Just found the perfect spot for my post-it notes on the fridge. Ready for a colorful summer!",
    "author": "SunshineLover",
    "userID": 12345678901234,
    "status": "enabled"
  },
  {
    "id": 2,
    "title": "Recipe Reminders",
    "content": "My fridge is covered in post-its with all my favorite recipes. No more forgetting ingredients!",
    "author": "ChefAtHeart",
    "userID": 98765432109876,
    "status": "enabled"
  },
  {
    "id": 3,
    "title": "Artistic Chaos",
    "content": "Turning my fridge into an art gallery with post-it doodles. Embracing the creative mess!",
    "author": "DoodleMaster",
    "userID": 54321098765432,
    "status": "enabled"
  },
  {
    "id": 4,
    "title": "Workout Motivation",
    "content": "Every morning, a new post-it with a motivational quote on the fridge. Let's crush those fitness goals!",
    "author": "FitLifeChampion",
    "userID": 13579246802468,
    "status": "enabled"
  },
  {
    "id": 5,
    "title": "Daily Gratitude",
    "content": "Started a gratitude post-it ritual on my fridge. It's amazing how a little note can change your perspective!",
    "author": "GratefulSoul",
    "userID": 24680135791234,
    "status": "enabled"
  },
  {
    "id": 6,
    "title": "Techie Notes",
    "content": "My fridge is now a tech hub with post-its filled with coding snippets and tech reminders.",
    "author": "CodeGeek",
    "userID": 98765432123456,
    "status": "enabled"
  },
  {
    "id": 7,
    "title": "Movie Night Picks",
    "content": "Keeping track of must-watch movies with post-its on the fridge. Movie night just got more organized!",
    "author": "Cinephile",
    "userID": 87654321098765,
    "status": "enabled"
  },
  {
    "id": 8,
    "title": "Travel Bucket List",
    "content": "Dreaming of exotic destinations? Post-it notes on the fridge are turning my travel dreams into plans!",
    "author": "WanderlustDreamer",
    "userID": 34567890123456,
    "status": "enabled"
  },
  {
    "id": 9,
    "title": "Bookworm's Corner",
    "content": "Transformed my fridge into a book recommendations space. Can't wait to explore new literary worlds!",
    "author": "BookishExplorer",
    "userID": 65432109876543,
    "status": "enabled"
  },
  {
    "id": 10,
    "title": "Daily Affirmations",
    "content": "Starting the day right with positive affirmations on the fridge. Today is going to be amazing!",
    "author": "PositivityFan",
    "userID": 12345098765432,
    "status": "enabled"
  },
  {
    "id": 11,
    "title": "Garden Goals",
    "content": "Planning my garden on the fridge with post-it notes. Excited to see it bloom!",
    "author": "GreenThumb",
    "userID": 98765123456789,
    "status": "enabled"
  },
  {
    "id": 12,
    "title": "Language Learning",
    "content": "Learning a new language? Post-it vocabulary on the fridge is my daily dose of language practice.",
    "author": "PolyglotExplorer",
    "userID": 54321234567890,
    "status": "enabled"
  },
  {
    "id": 13,
    "title": "Music Madness",
    "content": "Each post-it represents a song suggestion. Turning my fridge into a music discovery zone!",
    "author": "MusicLover",
    "userID": 13579135790246,
    "status": "enabled"
  },
  {
    "id": 14,
    "title": "Home DIY Ideas",
    "content": "Fridge post-its are now filled with DIY project ideas. Let the home improvement journey begin!",
    "author": "DIYEnthusiast",
    "userID": 98765432123456,
    "status": "enabled"
  },
  {
    "id": 15,
    "title": "Productivity Hacks",
    "content": "Boosting productivity with post-it notes on the fridge. A simple system for a more organized day!",
    "author": "ProductivityGuru",
    "userID": 65432109876543,
    "status": "enabled"
  },
  {
    "id": 16,
    "title": "Fitness Challenges",
    "content": "Setting fitness challenges with post-it goals on the fridge. Let the healthy competition begin!",
    "author": "FitnessFreak",
    "userID": 24680135792468,
    "status": "enabled"
  },
  {
    "id": 17,
    "title": "Mindfulness Moments",
    "content": "Fridge post-its now remind me to practice mindfulness daily. Embracing the present moment!",
    "author": "MindfulSoul",
    "userID": 13579246801357,
    "status": "enabled"
  },
  {
    "id": 18,
    "title": "Tech Innovation",
    "content": "Exploring the latest tech innovations with post-it notes on the fridge. The future is exciting!",
    "author": "TechEnthusiast",
    "userID": 78901234567890,
    "status": "enabled"
  },
  {
    "id": 19,
    "title": "Family Calendar",
    "content": "Keeping the family schedule organized with a fridge covered in post-it notes. No more missed events!",
    "author": "FamilyOrganizer",
    "userID": 98765432101234,
    "status": "enabled"
  },
];

</script>