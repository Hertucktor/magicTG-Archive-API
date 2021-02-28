function receiveSelected() {
    let setName = document.getElementById("setSelector").value;

    requestCardsBySetName(setName)
}

function requestCardsBySetName(setName) {
    let apiEndpoint = "localhost:8080/api/cards/set-names/"
    let URL = `${apiEndpoint + setName}`;
    console.log("request")

    loadCards(URL)
}

function loadCards(URL) {
    var xhttp = new XMLHttpRequest();
    console.log(xhttp.readyState)
    xhttp.onreadystatechange = function() {
        if (this.readyState === 4 && this.status === 200) {
            console.log(this.readyState)
            console.log(this.responseText);
        }
    };
    xhttp.open("GET", URL, true);
    xhttp.send();
}