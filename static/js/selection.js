function searchCardsBySet() {
    let setName = document.getElementById("setSelector").value;
    let apiEndpoint = "localhost:8080/api/cards/set-names/"
    let URL = `${apiEndpoint + setName}`;
    console.log(URL);
    //FIXME: GET request to local URL is not possible


    axios.get("https://httpbin.org/anything")
        .then(response => {
            console.log(`GET response`, response);
        })
        .catch(error => console.error(error));

}