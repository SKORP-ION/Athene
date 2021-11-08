var warehousesMap = {}
var statesMap = {}
var PropertyId = 0

async function getWarehouses(url) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        }
    })
    return await response.json();
}

async function getStates(url) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        }
    })
    return await response.json();
}

async function getPropertyInfo(url) {
    let urlParams = new URLSearchParams(window.location.search)

    if (urlParams.get("id") === null) {
        return null;
    }
    let id = urlParams.get("id")

    const response = await fetch(url + `?id=${id}`, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        }
    })
    return await response.json()
}

async function getPropertyHistory(url) {
    let urlParams = new URLSearchParams(window.location.search)

    if (urlParams.get("id") === null) {
        return null;
    }
    let id = urlParams.get("id")

    const response = await fetch(url + `?id=${id}`, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        }
    })
    return await response.json()
}

function TransformRow() {
    if (!this.classList.contains("selected")) {
        this.style.height = "200px";
        this.style.verticalAlign = "top";
        this.classList.add("selected")
    } else {
        this.style.height = "48px";
        this.style.verticalAlign = "";
        this.classList.remove("selected")
    }
}

function DoAction() {
    let data = [{
        "id": PropertyId,
        "serial": document.querySelector("#Serial").innerText,
        "inventory": document.querySelector("#Inventory").innerText,
        "name": document.querySelector("#Name").innerText
    }]

    document.cookie = `properties=${JSON.stringify({"props": data})}; path=/view/action; domain=${domain};`
    window.location.href = "/view/action"
}

getWarehouses(requestURL + "/private/property/getWarehouses")
.then((data)=>{
    //Запись в глобальную мапу
    for (i = 0; i < data.length; i++) {
        warehousesMap[data[i]["Id"]] = data[i]["Name"]
    }
    let warehouseField = document.querySelector("#Warehouse")
    let id = parseInt(warehouseField.innerHTML)
    warehouseField.innerHTML = warehousesMap[id]
})
.then(()=>{
    getStates(requestURL + "/private/property/getActions")
        .then((data)=>{
            for (i = 0; i < data.length; i++) {
                statesMap[data[i]["Id"]] = data[i]["Name"]
            }
        })
        .then(()=>{
        getPropertyHistory(requestURL + "/private/history/getHistory")
            .then((data)=>{
                let tbody = document.querySelector("#HistoryTable > tbody")
                for (let row of data) {
                    let tr = document.createElement("tr")

                    let action = document.createElement("td")
                    let img = document.createElement("img")
                    img.src = `/static/img/states/${row["Action"]}.png`
                    img.classList.add("statusIcon")
                    action.append(img)
                    action.append(statesMap[row["Action"]])
                    //action.innerHTML = img + statesMap[row["Action"]]
                    action.title = statesMap[row["Action"]]
                    action.classList.add("Action")
                    tr.append(action)

                    let note = document.createElement("td")
                    note.innerHTML = row["Note"] || ""
                    note.title = row["Note"] || ""
                    note.classList.add("Note")
                    tr.append(note)

                    let date = document.createElement("td")
                    date.innerHTML = FormatDate(new Date(row["Date"]["seconds"] * 1000))
                    date.title = FormatDate(new Date(row["Date"]["seconds"] * 1000))
                    date.classList.add("Date")
                    tr.append(date)

                    let user = document.createElement("td")
                    user.innerHTML = row["User"]
                    user.title = row["User"]
                    user.classList.add("User")
                    tr.append(user)
                    tr.addEventListener("click", TransformRow)

                    tbody.append(tr)
                }
                let tableCont = document.querySelector(".TableBody")
                tableCont.scrollBy(0, tableCont.querySelector("table").scrollHeight)
            })
    })
        .then(()=>{
        getPropertyInfo(requestURL + "/private/property/getOneProperty")
            .then((data)=>{
                PropertyId = data["Id"]
                document.querySelector("#Name").innerHTML = data["Name"] || ""
                document.querySelector("#Inventory").innerHTML = data["Inventory"] || ""
                document.querySelector("#Serial").innerHTML = data["Serial"] || ""
                document.querySelector("#StateImg").src = `/static/img/states/${data["State"]}big.png`
                document.querySelector("#StateImg").alt = statesMap[data["State"]]
                document.querySelector("#StateImg").title = statesMap[data["State"]]
                let date = new Date(data["Created_at"]["seconds"] * 1000)
                document.querySelector("#Created_at").innerHTML = FormatDate(date) || ""
                date = date = new Date(data["Updated_at"]["seconds"] * 1000)
                document.querySelector("#Updated_at").innerHTML = FormatDate(date) || ""
                document.querySelector("#Warehouse").innerHTML = warehousesMap[data["Warehouse"]] || ""
            })
    })
})

WriteBackUrl()
document.querySelector("#doAction").addEventListener("click", DoAction)