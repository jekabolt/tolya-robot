let form = document.querySelector('form');

form.addEventListener('submit', handleSubmit);
path = window.location.pathname.split('/')
id = path[path.length - 1]

function handleSubmit() {
    event.preventDefault();
    let checked = {};
    let checkFields = document.querySelectorAll('.checkbox__input:checked');

    checkFields.forEach(field => {
        let fieldName = field.name;
        if (fieldName.slice(-2) === '[]') { // Remove the "[]" in the end
            fieldName = fieldName.slice(0, -2);
        }

        if (!checked.hasOwnProperty(fieldName)) {
            checked[fieldName] = [];
        }

        checked[fieldName].push(parseInt(field.value));
    });

    url = "http://localhost:8080/api/v1.0/submit/" + id
    console.log(url);
    console.log(JSON.stringify(checked));

    createRequest(url, JSON.stringify(checked))
}

createRequest = function(url, postData) {

    var method = "POST";
    var shouldBeAsync = true;
    var request = new XMLHttpRequest();
    request.onload = function() {
        var status = request.status; // HTTP response status, e.g., 200 for "200 OK"
        var data = request.responseText; // Returned data, e.g., an HTML document.
        console.log("status:", status);
        console.log("data:", data);
    }

    request.open(method, url, shouldBeAsync);

    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

    request.send(postData);
}