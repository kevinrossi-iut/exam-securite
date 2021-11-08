function verify() {
    let name = document.getElementById("Username").value;
    let psw  = document.getElementById("Password").value;

    fetch('http://localhost:8080/connect/'+name+'/'+psw)
        .then(response => {
            if (response.ok) {
                return response.text()
            } else {
                return new Promise((resolve) => {
                    setTimeout(() => {
                        resolve('wrong user/password');
                    }, 300);
                });
            }
        }).then(response => {
            alert(response)
        }
    )
}
