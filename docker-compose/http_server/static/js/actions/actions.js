const tablebody = document.querySelector("#PropertyTable > tbody")
const states = document.querySelector("#states")
var statesMap = {}
var GeneralInventory = false
var GeneralName = false
var Action = 1

var Checked = false
var CheckTime = new Date()

/*Действия с таблицей*/
function ChangeNameEditable() {
    let status = this.checked
    let NameText = document.querySelector("#InputName")
    GeneralName = status
    if (status === true) {
        NameText.style.visibility = "visible"
        NameText.style.height = "30px"
        NameText.style.width = "300px"
        NameText.addEventListener("input", ChangeNameInTable)
        for (let row of tablebody.children) {
            let name = row.querySelector(".Name")
            if (!name.classList.contains("fixed")) {
                name.classList.remove("Editable")
                name.contentEditable = false
                name.prevText = name.innerHTML
                name.innerHTML = NameText.value
                CheckAction(name)
            }
        }
    } else {
        NameText.style.height = "0"
        NameText.style.width = "0"
        setTimeout(()=>{
            NameText.style.visibility = "hidden"
        }, 100)
        NameText.removeEventListener("input", ChangeNameInTable)
        for (let row of tablebody.children) {
            let name = row.querySelector(".Name")
            if (!name.classList.contains("fixed")) {
                name.classList.add("Editable")
                name.contentEditable = true
                name.innerHTML = name.prevText || ""
                CheckAction(name)
            }
        }
    }
}

function ChangeNameInTable() {
    let NameText = document.querySelector("#InputName")

    for (let row of tablebody.children) {
        let name = row.querySelector(".Name")
        if (!name.classList.contains("fixed")) {
            name.innerHTML = NameText.value
        }
    }
}

function ChangeInventoryEditable() {
    let status = this.checked
    let InventoryText = document.querySelector("#InputInventory")
    GeneralInventory = status
    if (status === true) {
        InventoryText.style.visibility = "visible"
        InventoryText.style.height = "30px"
        InventoryText.style.width = "300px"
        InventoryText.addEventListener("input", ChangeInventoryInTable)
        for (let row of tablebody.children) {
            let inventory = row.querySelector(".Inventory")
            if (!inventory.classList.contains("fixed")) {
                inventory.classList.remove("Editable")
                inventory.contentEditable = false
                inventory.prevText = inventory.innerHTML
                inventory.innerHTML = InventoryText.value
            }
        }
    } else {
        InventoryText.style.height = "0"
        InventoryText.style.width = "0"
        setTimeout(()=>{
            InventoryText.style.visibility = "hidden"
        }, 100)
        InventoryText.removeEventListener("input", ChangeInventoryInTable)
        for (let row of tablebody.children) {
            let inventory = row.querySelector(".Inventory")
            if (!inventory.classList.contains("fixed")) {
                inventory.classList.add("Editable")
                inventory.contentEditable = true
                inventory.innerHTML = inventory.prevText || ""
            }
        }
    }
}

function ChangeInventoryInTable() {
    let InventoryText = document.querySelector("#InputInventory")

    for (let row of tablebody.children) {
        let inventory = row.querySelector(".Inventory")
        if (!inventory.classList.contains("fixed")) {
            inventory.innerHTML = InventoryText.value
        }
    }
}

function CreateRow() {
    let tr = document.createElement("tr")

    let remove = document.createElement("td")
    remove.classList.add("Remove")
    tr.append(remove)

    let serial = document.createElement("td")
    serial.classList.add("Serial")
    serial.classList.add("Editable")
    serial.contentEditable = true
    serial.addEventListener("keydown", CheckInput)
    serial.addEventListener("paste", PasteFromClipboard)
    tr.append(serial)

    let inventory = document.createElement("td")
    inventory.classList.add("Inventory")
    if (!GeneralInventory) {
        inventory.contentEditable = true
        inventory.classList.add("Editable")
    } else {
        inventory.innerHTML = document.querySelector("#InputInventory").value
    }
    inventory.addEventListener("keydown", CheckInput)
    tr.append(inventory)

    let name = document.createElement("td")
    name.classList.add("Name")
    if (!GeneralName) {
        name.contentEditable = true
        name.classList.add("Editable")
    }  else {
        name.innerHTML = document.querySelector("#InputName").value
    }
    name.addEventListener("keydown", CheckInput)
    tr.append(name)

    let status = document.createElement("td")
    status.classList.add("Status")
    tr.append(status)

    tablebody.append(tr)
}

function RemoveRow() {
    let row = this.parentNode.parentNode
    let elements = tablebody.querySelectorAll("tr")

    if (elements.length <= 1) {
        return
    }
    row.remove()
}

function PasteFromClipboard(event) {
    event.preventDefault()
    let data = event.clipboardData.getData("text").split("\n")
    for (let row of data) {
        if (row === "") {
            continue
        }
        if (row.split("\t").length === 1) {
            let td = tablebody.querySelector("tr:last-child .Serial")
            td.innerText = row
            AddRemove(td)
            CreateRow()
            CheckAction(td)
        } else if (row.split("\t").length === 2) {
            let data = row.split("\t")
            let td = tablebody.querySelector("tr:last-child .Serial")
            td.innerText = data[0]
            let inventory = tablebody.querySelector("tr:last-child .Inventory")
            inventory.innerText = data[1]
            AddRemove(td)
            CreateRow()
            CheckAction(td)
        } else if (row.split("\t").length === 3) {
            let data = row.split("\t")
            let td = tablebody.querySelector("tr:last-child .Serial")
            td.innerText = data[0]
            let inventory = tablebody.querySelector("tr:last-child .Inventory")
            inventory.innerText = data[1]
            let name = tablebody.querySelector("tr:last-child .Name")
            name.innerText = data[2].slice(0, -1)
            AddRemove(td)
            CheckAction(td)
            CreateRow()
        }
    }
    event.stopPropagation()
}

/*Формы ввода*/
function ShowSelect(selector) {
    let slc = document.querySelector(selector)
    slc.style.display = "block"

    setTimeout(()=>{
        slc.style.height = "34px"
        slc.style.width = "400px"
        slc.style.visibility = "visible"
    }, 20)
}

function HideSelect(selector) {
    let slc = document.querySelector(selector)
    slc.style.height = ""
    slc.style.width = ""

    setTimeout(()=>{
        slc.style.display = "none"
        slc.style.visibility = ""
    }, 100)
}

/*Проверки*/
function CheckInput(event) {
    let elements = tablebody.querySelectorAll("tr")

    if (this.classList.contains("Serial")) {
        if (CheckCyrillic(this)){
            if (IsLastRow(this)) {
                AddRemove(this)
                CreateRow()
            }
            ChangeFocus(this)
            return
        } else if (CheckSpecSymbols(this)) {
            if (IsLastRow(this)) {
                AddRemove(this)
                CreateRow()
            }
            ChangeFocus(this)
            return
        }
    }

    if (!Checked) {
        StartCheckTimer(this)
    } else {
        CheckTime = new Date()
    }

    if (event.keyCode === 13) {
        event.preventDefault()
        if (IsLastRow(this)) {
            AddRemove(this)
            CreateRow()
        }
        ChangeFocus(this)
        return
    }

}

async function StartCheckTimer(node) {
    Checked = true
    CheckTime = new Date()

    while (true) {
        let delta = new Date() - CheckTime
        if (delta >= 200) {
            CheckAction(node)
            Checked = false
            break
        }

        await new Promise(r => setTimeout(r, 200))
    }
}

function CheckAction(node) {
    switch (Action) {
        case 3:
        case 4: HideSelect("#groups"); HideSelect("#warehouses"); ShowSelect("#employee");break;
        case 10: HideSelect("#employee"); HideSelect("#groups"); ShowSelect("#warehouses");break;
        case 15:
        case 16: HideSelect("#employee"); HideSelect("#warehouses"); ShowSelect("#groups");break;
        default: HideSelect("#employee"); HideSelect("#groups");HideSelect("#warehouses");break;
    }

    switch (Action) {
        case 1: CheckIsCreated(node);break;
        case 3: CheckIsReadyToGive(node);break;
        case 4: CheckIsWithEmployee(node);break;
        case 5:
        case 7:
        case 11: CheckIsInStock(node);break;
        case 6: CheckIsOnWorkspace(node);break;
        case 8: CheckIsNeedsRepair(node);break;
        case 9: CheckIsUnderRepair(node);break;
        case 10: CheckIsOnWarehouse(node);break;
        case 12: CheckIsNotCreated(node);UnFixInventories();break;
        case 13: CheckIsNotCreated(node);UnFixNames();break;
        case 14: CheckIsInArchive(node);break;
        case 15: CheckIsInGroup(node);break;
        case 16: CheckIsNotInGroup(node);break;
    }

    if (node.classList.contains("serial")) {
        ChangeFocus(node)
    }
}

function ChangeFocus(node) {
    if (node.classList.contains("Serial")) {
        if (!node.parentNode.querySelector(".Inventory").classList.contains("fixed") && !GeneralInventory) {
            node.parentNode.querySelector(".Inventory").focus()
        } else if (
            (node.parentNode.querySelector(".Inventory").classList.contains("fixed") || GeneralInventory)
            &&
            (!node.parentNode.querySelector(".Name").classList.contains("fixed")) && !GeneralName) {
            node.parentNode.querySelector(".Name").focus()
        } else {
            node.parentNode.parentNode.querySelector("tr:last-child .Serial").focus()
        }
    } else if (node.classList.contains("Inventory")) {
        if (!node.parentNode.querySelector(".Name").classList.contains("fixed") && !GeneralName) {
            node.parentNode.querySelector(".Name").focus()
        } else if (node.parentNode.querySelector(".Name").classList.contains("fixed") || GeneralName) {
            node.parentNode.parentNode.querySelector("tr:last-child .Serial").focus()
        }
    } else if (node.classList.contains("Name")) {
        node.parentNode.parentNode.querySelector("tr:last-child .Serial").focus()
    }
}

function ChangeAction() {
    Action = parseInt(this.value)

    let elements = tablebody.querySelectorAll("tr")

    FixExistsRows()

    for (let row of elements) {
        CheckAction(row.querySelector(".Serial"))
    }
}

function ChangeSelect() {
    let elements = tablebody.querySelectorAll("tr")

    FixExistsRows()

    for (let row of elements) {
        CheckAction(row.querySelector(".Serial"))
    }
}

function CreateAndAction() {
    let rows = tablebody.querySelectorAll("tr")
    let note = document.querySelector("#Description").value

    let toCreate = []
    let createDone = true
    let toAction = []

    async function action() {
        while (true) {
            if (createDone) {
                if (toAction.length >= 1 && ![1, 3, 4, 10, 15, 16].includes(Action)) {
                    ActionRequest(requestURL + GetUrl(), {
                        "note": note,
                        "properties": toAction
                    }).then((data)=>{
                        if (data["Ok"]) {
                            window.location.href = "/"
                            return
                        } else {
                            alert(`Error: ${data['Message']}`)
                            return
                        }
                    })
                    return
                } else if ([15, 16].includes(Action)) {
                    let group = document.querySelector("#groups").value
                    let ids = []
                    for (let action of toAction) {
                        ids.push(action["id"])
                    }
                    ActionRequest(requestURL + GetUrl(), {
                        "Note": note,
                        "Ids": ids,
                        "GroupId": parseInt(group),
                    }).then((data)=>{
                        if (data["Ok"]) {
                            window.location.href = "/"
                            return
                        } else {
                            alert(`Error: ${data['Message']}`)
                            return
                        }
                    })
                    return
                } else if (Action === 10) {
                    let warehouse = document.querySelector("#warehouses").value
                    let ids = []
                    for (let action of toAction) {
                        ids.push(action["id"])
                    }
                    ActionRequest(requestURL + GetUrl(), {
                        "note": note,
                        "Ids": ids,
                        "WarehouseId": parseInt(warehouse),
                    }).then((data)=>{
                        if (data["Ok"]) {
                            window.location.href = "/"
                            return
                        } else {
                            alert(`Error: ${data['Message']}`)
                            return
                        }
                    })
                    return
                } else if (Action === 3){
                    let employee = document.querySelector("#employee").employee
                    employee["Name"] = employee["Name"].replace("*", "")
                    if (!("Id" in employee)) {
                        POST(requestURL + "/private/staff/CreateEmployee", employee)
                            .then((data)=> {
                                if (data["Ok"]) {
                                    employee["Id"] = data["Employee"]["Id"]
                                } else {
                                    employee["Id"] = 0
                                }
                            })
                    }
                    //Ожидание получения id при создании
                    (async function Next() {
                        while (true) {
                            if ("Id" in employee) {
                                break
                            }
                            await new Promise(r => setTimeout(r, 20))
                        }
                        if (employee["Id"] === 0) {
                            alert("error")
                            return
                        }
                        let req = {
                            "EmployeeId": employee["Id"],
                            "Ids": [],
                        }
                        for (let elem of toAction) {
                            req.Ids.push(elem["id"])
                        }
                        POST(requestURL + GetUrl(), req)
                            .then((data)=>{
                                if (data["Ok"]) {
                                    window.location.href = "/"
                                    return
                                } else {
                                    alert(`Error: ${data['Message']}`)
                                    return
                                }
                            })
                    })().then(null)
                    return

                } else if (Action === 4) {
                    let employee = document.querySelector("#employee").employee

                    let req = {
                        "EmployeeId": employee["Id"],
                        "Ids": [],
                    }
                    for (let elem of toAction) {
                        req.Ids.push(elem["id"])
                    }
                    POST(requestURL + GetUrl(), req)
                        .then((data)=>{
                            if (data["Ok"]) {
                                window.location.href = "/"
                            } else {
                                alert(`Error: ${data['Message']}`)
                            }
                        })
                    return
                } else {
                    window.location.href = "/"
                }
            }
            await new Promise(r => setTimeout(r, 20))
        }
    }

    for (let row of Array.from(rows).slice(0, -1)) {
        if (row.toCreate) {
            let serial = row.querySelector(".Serial").innerHTML
            let inventory = row.querySelector(".Inventory").innerHTML || ""
            let name = row.querySelector(".Name").innerHTML || ""
            let elem = {
                "serial": serial,
                "inventory": inventory,
                "name": name
            }
            if (!toCreate.includes(elem)) {
                toCreate.push(elem)
            }
        }
        if (row.toAction && !row.toCreate) {
            let id = parseInt(row.id)
            let serial = row.querySelector(".Serial").innerHTML
            let inventory = row.querySelector(".Inventory").innerHTML
            let name = row.querySelector(".Name").innerHTML
            let elem = {
                "id": id,
                "serial": serial,
                "inventory": inventory,
                "name": name
            }
            if (!toAction.includes(elem)) {
                toAction.push(elem)
            }
        }
    }

    if (toCreate.length >= 1) {
        createDone = false
        let CreateNote = "Автоматически созданная карточка."

        if (Action === 1) {
            CreateNote = note
        }

        ActionRequest(requestURL + "/private/history/CreateCard", {
            "note": CreateNote,
            "properties": toCreate
        }).then((data)=> {
                if (data["Ok"]) {
                    for (let prop of data["Properties"]) {
                        toAction.push({
                            "id": prop["Id"],
                            "serial": prop["Serial"],
                            "inventory": prop["Inventory"],
                            "name": prop["Name"]
                        })
                    }
                } else {
                    alert(`Error: ${data['Message']}`)
                    createDone = true
                    return
                }
                createDone = true
            })
        action()
        return
    }

    if (toAction.length >= 1) {
        action()
    } else if (toCreate.length === 0 && toAction.length === 0) {
        alert("Нечего выполнять!")
    }
}

function GetPropertiesFromCookies() {
    const cookie = document.cookie

    if (!cookie) {
        return
    }
    const cookieValue = cookie
        .split('; ')
        .find(row => row.startsWith('properties='))
        .split('=')[1]

    let data = JSON.parse(cookieValue)
    console.log(data)
    return data
}

function WriteProperties() {
    let data = GetPropertiesFromCookies()

    if (data) {
        for (let prop of data["props"]) {
            let tr = document.createElement("tr")

            let remove = document.createElement("td")
            remove.classList.add("Remove")
            tr.append(remove)

            let serial = document.createElement("td")
            serial.classList.add("Serial", "fixed")
            serial.addEventListener("keydown", CheckInput)
            serial.innerHTML = prop["serial"]
            tr.append(serial)

            let inventory = document.createElement("td")
            inventory.classList.add("Inventory", "fixed")
            inventory.innerHTML = prop["inventory"]
            inventory.addEventListener("keydown", CheckInput)
            tr.append(inventory)

            let name = document.createElement("td")
            name.classList.add("Name", "fixed")
            name.innerHTML = prop["name"]
            name.addEventListener("keydown", CheckInput)
            tr.append(name)

            let status = document.createElement("td")
            status.classList.add("Status")
            tr.append(status)
            AddRemove(status)
            CheckAction(status)

            tablebody.append(tr)
        }
    }
}

function WriteBackUrl() {
    document.querySelector("#goBack").href = document.referrer
}

function FixExistsRows() {
    for (let row of tablebody.querySelectorAll("tr")) {
        if (row.id) {
            let inventory = row.querySelector(".Inventory")
            inventory.classList.add("fixed")
            inventory.classList.remove("Editable")
            inventory.contentEditable = false
            let name = row.querySelector(".Name")
            name.classList.add("fixed")
            name.classList.remove("Editable")
            name.contentEditable = false
        }
    }
}

function UnFixNames() {
    for (let row of tablebody.querySelectorAll("tr")) {
        let status = document.querySelector("#CheckName").checked
        if (row.id && !status) {
            let name = row.querySelector(".Name")
            name.classList.remove("fixed")
            name.contentEditable = true
            name.classList.add("Editable")
        }
    }
}

function UnFixInventories() {
    for (let row of tablebody.querySelectorAll("tr")) {
        let status = document.querySelector("#CheckInventory").checked
        if (row.id && !status) {
            let inventory = row.querySelector(".Inventory")
            inventory.classList.remove("fixed")
            inventory.classList.add("Editable")
            inventory.contentEditable = true
        }
    }
}

document.querySelector("#CheckName").addEventListener("click", ChangeNameEditable)
document.querySelector("#CheckInventory").addEventListener("click", ChangeInventoryEditable)
document.querySelector("#actionButton1").addEventListener("mouseup", CreateAndAction)
document.querySelector("#actionButton2").addEventListener("mouseup", CreateAndAction)
document.addEventListener("paste", (e)=>{
    e.preventDefault()
    let text = e.clipboardData.getData("text/plain")
    document.execCommand("insertHTML", false, text)
})
states.addEventListener("change", ChangeAction)
document.querySelector("#warehouses").addEventListener("change", ChangeSelect)
document.querySelector("#groups").addEventListener("change", ChangeSelect)


getStates(requestURL + "/private/property/getActions")
    .then((data)=>{
        //Запись в глобальную мапу
        for (i = 0; i < data.length; i++) {
            statesMap[data[i]["Id"]] = data[i]["Action"]
        }

        states.innerHTML = ""
        for (let i = 0; i < data.length; i++) {
            /*
            if ([3, 4].includes(data[i]["Id"])) {
                //Не добавляются действия, связанные с пользователем
                continue
            }
             */
            let opt = document.createElement("option")
            opt.value = data[i]["Id"]
            opt.innerText = data[i]["Action"]
            states.append(opt)
        }
    })

GET(requestURL + "/private/groups/GetGroups")
    .then((data)=>{
        let slc = document.querySelector("#groups")
        slc.innerHTML = ""
        for (let grp of data) {
            let option = document.createElement("option")
            option.value = grp["Id"]
            option.innerText = grp["Name"]
            slc.append(option)
        }
    })

GET(requestURL + "/private/property/getWarehouses")
    .then((data)=>{
        let slc = document.querySelector("#warehouses")
        slc.innerHTML = ""
        for (let elem of data) {
            let option = document.createElement("option")
            option.value = elem["Id"]
            option.innerText = elem["Name"]
            slc.append(option)
        }
    })

WriteProperties()
WriteBackUrl()
CreateRow()