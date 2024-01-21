/// authentication composable
export default async function() {
  const result = await $fetch('/api/auth',{credentials:"include"});
  if(result == false) {
      navigateTo('/login');
  } 
}