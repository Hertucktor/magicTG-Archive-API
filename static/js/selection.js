function searchCardsBySet() {
    let setName = document.getElementById("setSelector").value;
    let apiEndpoint = "http://localhost:8080/api/cards/set-names/"
    let URL = `${apiEndpoint + setName}`;

    axios.get(URL)
        .then(response => {
            console.log(`GET response`, response);
        })
        .catch(error => console.error(error));
}