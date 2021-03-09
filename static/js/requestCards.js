function searchCardsBySet() {
    let setName = document.getElementById("setSelector").value;
    let apiEndpoint = "http://localhost:8080/api/cards/set-names/"
    let URL = `${apiEndpoint + setName}`;

    axios.get(URL)
        .then(function (response) {
            let cards = response.data;
            sortCards(cards);
        })
        .catch(function (error) {
            if (error.response) {
                // The request was made and the server responded with a status code
                // that falls out of the range of 2xx
                console.log(error.response.data);
                console.log(error.response.status);
                console.log(error.response.headers);
            } else if (error.request) {
                // The request was made but no response was received
                // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
                // http.ClientRequest in node.js
                console.log(error.request);
            } else {
                // Something happened in setting up the request that triggered an Error
                console.log('Error', error.message);
            }
            console.log(error.config);
        });
}

function sortCards(cards) {
    console.log(cards.length)
    cards.forEach(function(card) {
        //object colors with index 0 = string(color)
        console.log(card['colors']);


        /*switch (card['colors']){
            case "White":
                console.log("White!");
                break;
            case "Blue":
                console.log("Blue!");
                break;
            case "Black":
                console.log("Black!");
                break;
            case "Green":
                console.log("Green!");
                break;
            case "Red":
                console.log("Red!");
                break;
            case card['colors'].length > 1:
                console.log("multi-color!");
                break;
            default:
                alert("default");
        }*/
    });
}

/*
function updateDisplay(card) {

}*/
