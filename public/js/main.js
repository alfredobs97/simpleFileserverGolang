//endpoints
const uploadEndpoint = "/upload";

//selectors
const inputSelector  = 'form input';
const buttonSelector = 'button';
const textSelector   = 'form p';

let input = document.querySelector(inputSelector);

input.onchange = () => {
    document.querySelector(textSelector).innerHTML = input.files.length + ' files added';
}

document.querySelector(buttonSelector).onclick = () => {
    let data = new FormData();

    if(input.files.length < 1) {
        alert("You have not selected any files")
        return
    }

    data.append('file', input.files[0]);

    upload(data);
}

const upload = (file) => {
    const request = {
        method: 'POST',
        body: file
    }

    fetch(uploadEndpoint, request);
} 
// https://github.com/pabloos