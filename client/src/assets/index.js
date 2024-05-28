
const loginForm =document.getElementById("login-form");
  loginForm.addEventListener('submit', onFormSubmit);


async function onFormSubmit(event) {
	event.preventDefault();

	const data = new FormData(event.target);


  await fetch("http://localhost:3000/auth/login", {
    method: "POST",
    // headers: {'Content-Type': 'application/json',

    // }, 
    body: JSON.stringify(
      {
        email:data.get('email'),
        password:data.get('password'),

      }
    )
  }).then(res => {
    console.log("Request complete! response:", res);
    window.location.href = "/user.html";
  }).catch(err=>{
    console.log(err)
  });
}


const registerForm =document.getElementById("register-form");
registerForm.addEventListener('submit', onRegisterFormSubmit);


async function onRegisterFormSubmit(event) {
	event.preventDefault();

	const data = new FormData(event.target);


  await fetch("http://localhost:3000/auth/register", {
    method: "POST",
    // headers: {'Content-Type': 'application/json',

    // }, 
    body: JSON.stringify(
      {
        username:data.get('username'),
        email:data.get('email'),
        phone:data.get('phone'),
        password:data.get('password'),
        passwordconfirmation:data.get('passwordconfirmation'),

      }
    )
  }).then(res => {
    console.log("Request complete! response:", res);
    window.location.href = "/login.html";
  }).catch(err=>{
    console.log(err)
  });
}