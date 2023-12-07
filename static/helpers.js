function createHTMLTable(jsonData) {
    let tableHTML = '<table>';

    // Create table header
    tableHTML += '<thead><tr>';
    for (let key in jsonData[0]) {
        tableHTML += '<th>' + key + '</th>';
    }
    tableHTML += '</tr></thead>';

    // Create table body
    tableHTML += '<tbody>';
    for (let i = 0; i < jsonData.length; i++) {
        tableHTML += '<tr>';
        for (let key in jsonData[i]) {
            tableHTML += '<td>' + jsonData[i][key] + '</td>';
        }
        tableHTML += '</tr>';
    }
    tableHTML += '</tbody>';

    tableHTML += '</table>';

    return tableHTML;
}

function generalizedGet(listenerElement, url, displayElement, idKey) {
    listenerElement.addEventListener("submit", function (event) {
        event.preventDefault();
        const formData = new FormData(listenerElement);
        const codeData = Object.fromEntries(formData.entries());
        fetch(idKey?
            `${url}/${codeData[idKey]}`:
            `${url}`
        ).then(response => response.json())
            .then(data => {
                displayElement.innerText = createHTMLTable(data);
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

function generalizedDelete(listenerElement, url, idKey) {
    listenerElement.addEventListener("submit", function (event) {
        event.preventDefault();
        const formData = new FormData(listenerElement);
        const codeData = Object.fromEntries(formData.entries());
        fetch(`${url}/${codeData[idKey]}`, {
            method: 'DELETE'
        })
            .then(response => response.text())
            .then(data => {
                console.log("Inspector Code deleted:", data);
                // Handle success or display a success message
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}

function generalizedUpdateCreate(formElement, url, switchKey, idKey) {
    formElement.addEventListener("submit", function (event) {
        event.preventDefault();
        const formData = new FormData(formElement);
        const codeData = Object.fromEntries(formData.entries());
        fetch(codeData[switchKey] ?
            `${url}/${codeData[idKey]}` :
            `${url}`, {
            method: codeData[switchKey] ? 'PUT' : 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(codeData)
        })
            .then(response => response.text())
            .then(data => {
                console.log("Item created/updated:", data);
                // Handle success or display a success message
            })
            .catch(error => {
                console.error("Error:", error);
            });
    });
}