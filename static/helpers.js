function createHTMLTable(jsonData, appendTo) {
    let table = document.createElement("table");
    table.className = "table table-striped";
    let thead = document.createElement("thead");
    let tbody = document.createElement("tbody");
    let tr = document.createElement("tr");
    for (let key in jsonData[0]) {
        let th = document.createElement("th");
        th.innerText = JSON.stringify(key);
        tr.appendChild(th);
    }
    thead.appendChild(tr);
    table.appendChild(thead);
    for (let i = 0; i < jsonData.length; i++) {
        let tr = document.createElement("tr");
        for (let key in jsonData[i]) {
            let td = document.createElement("td");
            const value = jsonData[i][key];
            console.log(value);
            if(typeof value === "object"){
                if(value["String"]) td.innerText = JSON.stringify(value["String"])
                else if(value["Int32"]) td.innerText = JSON.stringify(value["Int32"])
                else td.innerText = JSON.stringify(value)
            }
            else td.innerText = value;
            tr.appendChild(td);
        }
        tbody.appendChild(tr);
    }
    table.appendChild(tbody);
    if (appendTo) appendTo.appendChild(table);
    return table.outerHTML;
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
                 createHTMLTable(data, displayElement);
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