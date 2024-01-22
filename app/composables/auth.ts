/// authentication composable
export default async function() {
  const result = await $fetch('http://localhost:3000/api/user/auth',{credentials:"include"});
  if(result == false) {
      navigateTo('/login');
  } 
}