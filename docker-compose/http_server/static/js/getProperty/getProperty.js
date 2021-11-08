const properties = document.querySelector("#PropertyTable")
const warehouses = document.querySelector("#warehouses")
const states = document.querySelector("#states")
var warehousesMap = {}
var statesMap = {}

let needSearch = false
let timeout = new Date()
let symbolNumber = 0

async function SearchGroups(names) {
    const datalist = document.querySelector("#groupsList")
    while (true) {
        if (!needSearch) {
            return
        }
        let delta = new Date() - timeout

        if (delta > 500 && needSearch && symbolNumber >= 1) {
            needSearch = false
            GET(requestURL + `/private/groups/GetGroups?name=${names[names.length - 1]}`)
                .then((data)=>{
                    if (data === null) {
                        console.log("Empty response")
                        return
                    }
                    datalist.innerHTML = ""
                    for (let group of data) {
                        let option = document.createElement("option")
                        option.value = group["Name"]
                        option.addEventListener("click", WriteGroup)
                        datalist.append(option)
                    }
                })
        } else if (symbolNumber < 3) {
            datalist.innerHTML = ""
        }
        await new Promise(r => setTimeout(r, 200))
    }
}

function getSearchProperties() {
    var params = {
        "Inventory": "",
        "Serial": "",
        "Action": 0,
        "Warehouse": 0,
        "Offset": 0,
        "Limit": 30,
        "Order": "",
        "Groups": [],
    }

    let urlParams = new URLSearchParams(window.location.search)

    if (urlParams.get("limit") !== null) {
        params["Limit"] = parseInt(urlParams.get("limit"))
    }

    if (urlParams.get("page") !== null) {
        params["Offset"] = ((urlParams.get("page") - 1) * params["Limit"])
    }

    if (urlParams.get("inventory") !== null) {
        params["Inventory"] = urlParams.get("inventory")
    }

    if (urlParams.get("serial") !== null) {
        params["Serial"] = urlParams.get("serial")
    }

    if (urlParams.get("name") !== null) {
        params["Name"] = urlParams.get("name")
    }

    if (urlParams.get("groups") !== null) {
        params["Groups"] = urlParams.get("groups").split(";")
    }

    if (urlParams.get("action") !== null) {
        params["Action"] = parseInt(urlParams.get("action"))
    }

    if (urlParams.get("warehouse") !== null) {
        params["Warehouse"] = parseInt(urlParams.get("warehouse"))
    }

    if (urlParams.get("order") !== null) {
        params["Order"] = urlParams.get("order")
    }


    return params
}

function WriteFormFromParams() {
    let params = new URLSearchParams(window.location.search)

    if (params.get("inventory") !== null) {
        document.querySelector("#inventory").value = params.get("inventory")
    }

    if (params.get("serial") !== "") {
        document.querySelector("#serial").value = params.get("serial")
    }

    if (params.get("name") !== "") {
        document.querySelector("#name").value = params.get("name")
    }

    if (params.get("groups") !== "") {
        document.querySelector("#groups").value = params.get("groups")
    }

    if (params.get("action") !== null) {
        document.querySelector("#states").value = parseInt(params.get("action")) || ""
    }

    if (params.get("warehouse") !== "" && params.get("warehouse") != null) {
        document.querySelector("#warehouses").value = parseInt(params.get("warehouse"))
    }

    if (params.get("order") !== null) {
        switch (params.get("order")) {
            case "inventory0": document.querySelector("#orderInventory").classList.add("ASC"); break;
            case "inventory1": document.querySelector("#orderInventory").classList.add("DESC");break;
            case "serial0": document.querySelector("#orderSerial").classList.add("ASC");break;
            case "serial1": document.querySelector("#orderSerial").classList.add("DESC");break;
            case "name0": document.querySelector("#orderName").classList.add("ASC");break;
            case "name1": document.querySelector("#orderName").classList.add("DESC");break;
            case "warehouse0": document.querySelector("#orderWarehouse").classList.add("ASC");break;
            case "warehouse1": document.querySelector("#orderWarehouse").classList.add("DESC");break;
            case "created_at0": document.querySelector("#orderCreated_at").classList.add("ASC");break;
            case "created_at1": document.querySelector("#orderCreated_at").classList.add("DESC");break;
            case "updated_at0": document.querySelector("#orderUpdated_at").classList.add("ASC");break;
            case "updated_at1": document.querySelector("#orderUpdated_at").classList.add("DESC");break;
        }
    }
}

function selectPage(count) {
    let params = new URLSearchParams(window.location.search)

    let pageId = params.get("page") || "1"

    let page = document.querySelector(`#page${pageId}`)

    page.classList.add("selected")

    let limit = parseInt(getSearchProperties()["Limit"]) || 30

    let offset = parseInt(pageId)

    let pagesCount = Math.trunc(count / limit) + 1

    let result = ""

    if (offset < pagesCount) {
        result = `${limit}/${count}`
    } else {
        result = `${count % limit}/${count}`
    }

    let span = document.createElement("span")
    span.innerText = result

    let pages = document.querySelector("#pagesDiv")
    pages.append(span)

}

function SetOrder() {
    let orderBy = ""

    switch (this.id) {
        case "orderInventory": orderBy = "inventory";break;
        case "orderSerial": orderBy = "serial";break;
        case "orderName": orderBy = "name";break;
        case "orderCreated_at": orderBy = "created_at";break;
        case "orderUpdated_at": orderBy = "updated_at";break;
        case "orderWarehouse": orderBy = "warehouse";break;
    }

    if (this.classList.contains("ASC")) {
        orderBy += "1"
    } else {
        orderBy += "0"
    }

    let url =  window.location.href
    let params = new URLSearchParams(new URL(url).search)
    if (Array.from(params).length === 0) {
        url += "?order=" + orderBy
    }
    else if (url.includes("order=")) {
        url = url.replace(/(order=)\D*(0|1)/, "order=" + orderBy )
    } else {
        url += "&order=" + orderBy
    }

    window.location.href = url
}

function BindOrderButtons() {
    let orders = document.querySelectorAll(".Ordered")

    for (let order of orders) {
        order.addEventListener("click", SetOrder)
    }
}

function DoAction() {
    let props = []

    for (let row of selectedRows) {
        props.push({
            "id": row.propertyId,
            "serial": row.querySelector(`#${row.id}serial`).innerHTML || "",
            "inventory": row.querySelector(`#${row.id}inventory`).innerHTML || "",
            "name": row.querySelector(`#${row.id}name`).innerHTML || ""
        })
    }
    let data = {
        "props": props
    }
    document.cookie = `properties=${JSON.stringify(data)}; path=/view/action; domain=${domain};`
    window.location.href = "/view/action"
}

function GetGroups() {
    let name = document.querySelector("#groups").value
    symbolNumber = name.length
    if (!needSearch) {
        needSearch = true
        timeout = new Date()
        SearchGroups(name.split(","))
    }

}

function WriteGroup(event) {
    event.preventDefault()
    document.documentElement.style.setProperty("--content", this.value)
}

GET(requestURL + "/private/property/getWarehouses")//GET WAREHOUSES
    .then((data) => {
        console.log(data)

        //Запись в глобальную мапу
        for (i = 0; i < data.length; i++) {
            warehousesMap[data[i]["Id"]] = data[i]["Name"]
        }

        warehouses.innerHTML = ""
        let opt = document.createElement("option")
        opt.value = ""
        opt.innerText = "Все площадки"
        warehouses.append(opt)
        for (let i = 0; i < data.length; i++) {
            let opt = document.createElement("option")
            opt.value = data[i]["Id"]
            opt.innerText = data[i]["Name"]
            warehouses.append(opt)
        }
        return
    })
    .then(()=>{
        GET(requestURL + "/private/property/getActions")//GET ACTIONS
            .then((data)=>{
                states.innerHTML = ""
                let opt = document.createElement("option")
                opt.value = ""
                opt.innerText = "Любой"
                states.append(opt)
                for (let i = 0; i < data.length; i++) {
                    if ([12, 13, 15, 16].includes(data[i]["Id"])) {
                        continue
                    }
                    statesMap[data[i]["Id"]] = data[i]["Name"]
                    let opt = document.createElement("option")
                    opt.value = data[i]["Id"]
                    opt.innerText = data[i]["Name"]
                    states.append(opt)
                }
            })
            .then(WriteFormFromParams)
    })
    .then(() => {
        POST(requestURL + "/private/property/getCount", getSearchProperties())//GET COUNT
            .then((data) => {
                let pages = document.querySelector("#pagesDiv")
                console.log(data)
                pages.innerHTML = ""
                let limit = parseInt(getSearchProperties()["Limit"]) || 30
                for (i = 1; i <= (data["Number"] / limit) + 1; i++) {
                    let page = document.createElement("a")
                    page.href = function () {
                        let url =  window.location.href
                        let params = new URLSearchParams(new URL(url).search)
                        if (Array.from(params).length === 0) {
                            url += "?page=" + i
                        }
                        else if (url.includes("page=")) {
                            url = url.replace(/(page=)\d*/, "page=" + i )
                        } else {
                            url += "&page=" + i
                        }
                        return url
                    } ()
                    page.rel = "keep-params"
                    page.innerText = (i).toString()
                    page.id = `page${i}`
                    pages.append(page)
                }
                selectPage(data["Number"])
            })
            .then(() => {
            POST(requestURL + "/private/property/getProperty", getSearchProperties()) //GET PROPERTY
                .then((data) => {
                    let rows = document.querySelectorAll(".TableRow")
                    for (i = 0; i < rows.length; i++) {
                        rows[i].remove()
                    }

                    let propData = data["Properties"]

                    for (i = 0; i < propData.length; i++) {
                        let row = document.createElement("tr")
                        row.id = `row${i}`
                        row.propertyId = propData[i]["Id"]

                        let status = document.createElement("td")
                        let icon = document.createElement("img")
                        let state = propData[i]["State"]
                        icon.src = `/static/img/states/${state}.png`
                        icon.classList.add("statusIcon")
                        icon.title = statesMap[state]
                        status.append(icon)

                        row.append(status)

                        let inventory = document.createElement("td")
                        inventory.innerText = propData[i]["Inventory"] || ""
                        inventory.id = `row${i}inventory`
                        row.append(inventory)

                        let serial = document.createElement("td")
                        serial.innerText = propData[i]["Serial"] || ""
                        serial.id = `row${i}serial`
                        row.append(serial)

                        let name = document.createElement("td")
                        name.innerText = propData[i]["Name"] || ""
                        name.id = `row${i}name`
                        row.append(name)

                        let created_at = document.createElement("td")
                        let date = new Date(propData[i]["Created_at"]["seconds"] * 1000)
                        created_at.innerText = FormatDate(date)
                        created_at.id = `row${i}created_at`
                        row.append(created_at)

                        let updated_at = document.createElement("td")
                        date = new Date(propData[i]["Updated_at"]["seconds"] * 1000)
                        updated_at.innerText = FormatDate(date)
                        updated_at.id = `row${i}updated_at`
                        row.append(updated_at)

                        let warehouse = document.createElement("td")
                        warehouse.innerText = warehousesMap[propData[i]["Warehouse"]]
                        warehouse.id = `row${i}warehouse`
                        row.append(warehouse)

                        properties.append(row)
                        row.onmousedown = StartSelection;
                        row.onmouseup = EndSelections;
                    }
                })
        })
    return
    })
    .then(() => {
        BindOrderButtons()
})

document.querySelector("#doAction").addEventListener("click", DoAction)
document.querySelector("#groups").addEventListener("input", GetGroups)