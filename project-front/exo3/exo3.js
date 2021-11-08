function addMessage() {
    let text = document.getElementById("text").value;
    fetch('http://localhost:8080/comment', {
        method: 'POST',
        body: JSON.stringify({
            comment: text
        })
    }).then(() => window.location = '.')
}

function getMessage() {
    fetch('http://localhost:8080/comment')
        .then(response => response.json())
        .then(response => {
                for (x in response) {
                    document.getElementById("content").innerHTML += "<p>" + response[x].comment + "</p>";
                }
            }
        )
}

/* Override window.alert */
var originalAlert = window.alert;
window.alert = function(s) {
    parent.postMessage("success", "*");
    setTimeout(function() {
        originalAlert("Congratulations, you executed an alert:\n\n"
            + "\n\nYou can now advance to the next level.");
    }, 50);
}