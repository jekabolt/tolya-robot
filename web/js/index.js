let form = document.querySelector('form');

form.addEventListener('submit', handleSubmit);

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

        checked[fieldName].push(field.value);
    });

    console.log(checked);
}