function verify() {
    fetch('http://localhost:8080/user-agent')
        .then(response => {
            if (response.ok) {
                return response.text()
            } else {
                return new Promise((resolve) => {
                    setTimeout(() => {
                        resolve('wrong user-agent');
                    }, 300);
                });
            }
        }).then(response => {
            alert(response)
        }
    )
}
