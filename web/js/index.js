let form = document.querySelector('form');

form.addEventListener('submit', handleSubmit);
path = window.location.pathname.split('/')
id = path[path.length - 1]


// 2) брюки, штаны, юбки и т.д. option63
// 3) футболки, худи, блузки, платья и т.д. option62
// 4) куртки, пальто, шубы и т.д. option61
// 5) аксессуары (сумки, перчатки, очки, нижнее белье и т.д.) option65

// 2) штаны, шорты, брюки и т.д. option63
// 3) футболки, худи, байки и т.д. option62
// 4) куртки, пальто и т.д. option61
// 5) аксессуары (рюкзаки, перчатки, очки, нижнее белье и т.д.) option65


function onlyOneFemale() {

    var bottoms = document.getElementById("option63for");
    var tops = document.getElementById("option62for");
    var out = document.getElementById("option61for");
    var accs = document.getElementById("option65for");


    var casual = document.getElementById("option21for");
    var street = document.getElementById("option22for");
    var classic = document.getElementById("option23for");
    var avantgarde = document.getElementById("option24for");


    var checkBoxFemale = document.getElementById("option12");
    if (checkBoxFemale.checked == true) {
        checkBoxMale = document.getElementById("option11");
        checkBoxMale.checked = false

        casual.src = "https://sun9-53.userapi.com/c206520/v206520576/7de3b/vI9vP4U6GKk.jpg"
        street.src = "https://sun9-46.userapi.com/c206520/v206520576/7de4f/tLzLkozupXE.jpg"
        classic.src = "https://sun9-13.userapi.com/c206520/v206520576/7de27/VQYb_eUpnDM.jpg"
        avantgarde.src = "https://sun9-70.userapi.com/c206520/v206520576/7de13/nYDNlMybdj4.jpg"

        bottoms.innerText = "брюки, штаны, юбки и т.д."
        tops.innerText = "футболки, худи, блузки, платья и т.д."
        out.innerText = "куртки, пальто, шубы и т.д."
        accs.innerText = "аксессуары (сумки, перчатки, очки, нижнее белье и т.д.)"
    }
}


function handleGender() {
    var checkBox = document.getElementById("option11");
    var bottoms = document.getElementById("option63for");
    var tops = document.getElementById("option62for");
    var out = document.getElementById("option61for");
    var accs = document.getElementById("option65for");

    var casual = document.getElementById("option21for");
    var street = document.getElementById("option22for");
    var classic = document.getElementById("option23for");
    var avantgarde = document.getElementById("option24for");

    if (checkBox.checked == true) {

        casual.src = "https://sun9-43.userapi.com/c206520/v206520576/7de45/IXT8Ikf3Qu4.jpg"
        street.src = "https://sun9-7.userapi.com/c206520/v206520576/7de59/YD3CGKz40-w.jpg"
        classic.src = "https://sun9-4.userapi.com/c206520/v206520576/7de31/MPT5YDMrre4.jpg"
        avantgarde.src = "https://sun9-41.userapi.com/c206520/v206520576/7de1d/NP3tYhhpjgk.jpg"

        checkBoxFemale = document.getElementById("option12");
        checkBoxFemale.checked = false
        bottoms.innerText = "штаны, шорты, брюки и т.д."
        tops.innerText = "футболки, худи, байки и т.д."
        out.innerText = "куртки, пальто и т.д."
        accs.innerText = "аксессуары (рюкзаки, перчатки, очки, нижнее белье и т.д.)"
    }
}

handleGender()

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

    if (checkFields.length < 6) {
        alert("Выбери как минимум одно поле из каждой категории")
        return
    }


    // const url = "http://example.com";
    url = "http://localhost:8080/api/v1.0/submit/" + id
    fetch(url, {
        method: "POST",
        body: JSON.stringify(checked),
    }).then(
        response => response.text() // .json(), etc.
    ).then(
        html => console.log(html)
    );


    // url = "http://localhost:8080/api/v1.0/submit/" + id
    // console.log(url);
    // console.log(JSON.stringify(checked));

    // status = createRequest(url, JSON.stringify(checked))
    // window.location.href = "https://telegram.me/tolyahobot"
}



createRequest = function(url, postData) {
    var method = "POST";
    var shouldBeAsync = true;
    var request = new XMLHttpRequest();
    request.onload = function() {
        var status = request.status; // HTTP response status, e.g., 200 for "200 OK"
        var data = request.responseText; // Returned data, e.g., an HTML document.
    }

    request.open(method, url, shouldBeAsync);

    request.setRequestHeader("Content-Type", "application/json;charset=UTF-8");

    request.send(postData);
}