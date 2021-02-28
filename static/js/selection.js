function receiveSelected() {
    let setName = document.getElementById("setSelector").value;

    requestCardsBySetName(setName)
}

function requestCardsBySetName(setName) {
    let apiEndpoint = "localhost:8080/api/cards/set-names/"
    let URL = `${apiEndpoint + setName}`;
    alert(URL);
    /*$(document).ready(function() {
        $.get(URL, function(data, status){
            //alert("Data: " + data + "\nStatus: " + status);
            console.log(status)
        });
    });*/
}