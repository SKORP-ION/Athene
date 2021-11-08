const tbody = document.querySelector(".TableBody  tbody")
var id = 0
var statesMap = {}


function GetNameFromUrl() {
    let params = new URLSearchParams(window.location.search)

    if (params.get("name")) {
        return params.get("name")
    }
    return ""
}

function getSearchProperties() {
    var params = {
        "Inventory": "",
        "Serial": "",
        "Action": 0,
        "Warehouse": 0,
        "Offset": 0,
        "Limit": 10000,
        "Order": "",
        "Groups": [],
    }

    let urlParams = new URLSearchParams(window.location.search)

    if (urlParams.get("name") !== null) {
        params["Groups"] = urlParams.get("name").split(";")
    }


    return params
}

function OpenProperty() {
    console.log("Open property")
    let id = this.id
    window.location.href = requestURL + `/view/property?id=${id}`
}

GET(requestURL + `/private/groups/GetGroups?name=${GetNameFromUrl()}`)
    .then((data)=>{
        let grp = data[0]
        id = grp["Id"]
        document.querySelector("#Name").innerHTML = grp["Name"] || ""
        document.querySelector("#Description").innerHTML = grp["Description"] || ""
        document.querySelector("#Username").innerHTML = grp["WhoDisplayName"] || ""
        document.querySelector("#Created_at").innerHTML =
            FormatDate(new Date(grp["CreatedAt"]["seconds"] * 1000))
    })

GET(requestURL + "/private/property/getActions")
    .then((data)=>{
        for (let state of data) {
            statesMap[state["Id"]] = state["Name"]
        }

        POST(requestURL + "/private/property/getProperty", getSearchProperties())
            .then((data)=>{
                tbody.innerHTML = ""
                let props = data["Properties"]
                for (let prop of props) {
                    let tr = document.createElement("tr")
                    tr.id = prop["Id"]

                    let state = document.createElement("td")
                    let icon = document.createElement("img")
                    icon.src = `/static/img/states/${prop["State"]}.png`
                    icon.classList.add("statusIcon")
                    icon.title = statesMap[prop["State"]]
                    state.append(icon)
                    tr.append(state)

                    let serial = document.createElement("td")
                    serial.classList.add("Serial")
                    serial.innerText = prop["Serial"] || ""
                    tr.append(serial)

                    let inventory = document.createElement("td")
                    inventory.classList.add("Inventory")
                    inventory.innerText = prop["Inventory"] || ""
                    tr.append(inventory)

                    let name = document.createElement("td")
                    name.classList.add("Name")
                    name.innerText = prop["Name"] || ""
                    tr.append(name)

                    tr.addEventListener("click", OpenProperty)

                    tbody.append(tr)
                }
            })
    })

WriteBackUrl()