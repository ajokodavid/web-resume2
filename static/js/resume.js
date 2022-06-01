function welcomeUser() {
    alert("Welcome to my resume site")
 }
 
 welcomeUser();

 function displayNumberOfProjects() {
     let opt = document.querySelector("output");
     for(let i = 0; i < 20; i++) {
         opt.innerHTML = i;
     }
 }

 displayNumberOfProjects();