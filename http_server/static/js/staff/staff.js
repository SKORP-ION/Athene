statesMap = {}


function ShowEmployeeInfo () {
    let employee = document.querySelector("#employee").employee
    let info = document.querySelector(".EmployeeInfo")
    info.employeeId = employee["Id"]
    let name = document.createElement("a")
    name.innerHTML = employee["Name"] || ""
    name.href = `https://w3c.portal.rt.ru/profiles/html/simpleSearch.do?searchBy=name&lang=ru&searchFor=`
    name.href += employee["Name"].replace(" ", "%20")
    info.querySelector("#Name").innerHTML = ""
    info.querySelector("#Name").append(name)
    info.querySelector("#Table").innerHTML = employee["Table"] || ""
    info.querySelector("#Job").innerHTML = employee["Job"] || ""
    info.querySelector("#Manager").innerHTML = employee["Manager"] || ""
    info.querySelector("#Department").innerHTML = employee["Department"] || ""
    info.querySelector("#Created_at").innerHTML = FormatDate(new Date(employee["CreatedAt"]["seconds"] * 1000))
    info.style.visibility = "visible"
}

function ShowTable () {
    let container = document.querySelector("#Properties")
    let tbody = container.querySelector("tbody")
    let employee = document.querySelector("#employee").employee
    id = parseInt(employee["Id"])

    GET(requestURL + `/private/staff/GetStaffProps?id=${id}`)
        .then((data)=>{
            let props = data["Props"]
            tbody.innerHTML = ""
            for (let prop of props) {
                let tr = document.createElement("tr")
                tr.propertyId = prop["Id"]
                tr.recordId = prop["RecordId"]
                tr.addEventListener("mouseup", CheckEvent)

                let state = document.createElement("td")
                state.classList.add("Action")
                let img = document.createElement("img")
                img.src = `/static/img/states/${prop["State"]}.png`
                img.classList.add("statusIcon")
                img.title = statesMap[prop["State"]]
                state.append(img)
                tr.append(state)

                let serial = document.createElement("td")
                serial.classList.add("Serial")
                serial.innerHTML = prop["Serial"] || ""
                tr.append(serial)

                let inventory = document.createElement("td")
                inventory.classList.add("Inventory")
                inventory.innerHTML = prop["Inventory"] || ""
                tr.append(inventory)

                let name = document.createElement("td")
                name.classList.add("Name")
                name.innerHTML = prop["Name"] || ""
                tr.append(name)

                let givenAt = document.createElement("td")
                givenAt.classList.add("GivenAt")
                givenAt.innerHTML = FormatDate(new Date(prop["GivenAt"]["seconds"] * 1000))
                tr.append(givenAt)

                tbody.append(tr)
            }
            container.style.visibility = "visible"
        })

}
GET(requestURL + "/private/property/getActions")
    .then((data)=>{
        for (i = 0; i < data.length; i++) {
            statesMap[data[i]["Id"]] = data[i]["Name"]
        }
    })