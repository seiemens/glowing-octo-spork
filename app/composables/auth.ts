/// authentication composable
export default async function() {
  return await $fetch('http://localhost:8080/api/user/auth',{
    credentials:"include"
  }).then((res)=>{
    return true;
  }).catch((err)=> {
    return false;
  });
}