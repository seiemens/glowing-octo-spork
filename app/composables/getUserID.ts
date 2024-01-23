export const getUserID = async () => {
  return await $fetch('http://localhost:8080/api/user/info', {
    credentials:'include'
  }).then((res:any)=>{
    return res.id;
  }).catch(()=>{
    return false;
  })
}