export const getUserID = () => {
  const cookieName = "user";
  let cookies = document.cookie.split("; ");
  let val;

  cookies.forEach((cookie)=>{
    let [cname, cval] = cookie.split("=");

    if (cname === cookieName) {
      val = cval;
    }
  });

  return val || false;
}