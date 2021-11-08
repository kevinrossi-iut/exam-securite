function storeDataInSession() {
    let name = document.getElementById("Username").value;
    let psw  = document.getElementById("Password").value;
    sessionStorage.setItem("Username",name);
    sessionStorage.setItem("Password",psw);
}

function verify() {
    let name = sessionStorage.getItem("Username");
    let psw  = sessionStorage.getItem("Password");

    if (name === "Adm1n" && psw === "GKyMxMNdNYrMFtvXae14") {
        let body = document.getElementById("all")
        body.style.display = "block"
    } else {
        window.location.replace(".");
    }
}
