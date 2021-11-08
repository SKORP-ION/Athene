const createPopup = document.querySelector("#CreatePopup")
const deletePopup = document.querySelector("#DeletePopup")
const createGroup = document.querySelector("#CreateGroup")
const background = document.querySelector("#BlackBackGround")
var currentPopup

/*Функционал появления и исчезания меню*/
function OpenCreateGroup() {
    createPopup.style.visibility = "visible"
    createPopup.style.top = "30%"
    createPopup.style.left = "30%"
    createPopup.style.height = "30%"
    createPopup.style.width = "40%"
    background.style.visibility = "visible"
    createPopup.querySelector("#CreateMessage").innerHTML = "Создание группы"
    let button = createPopup.querySelector(".inputButton")
    button.addEventListener("click", CreateGroup)

    setTimeout(()=>{
        window.removeEventListener("click", CheckMouseClick)
        window.addEventListener("mousedown", CheckMouseClickPopup)
        currentPopup = createPopup
    }, 200)
}

function CloseCreateGroup() {
    createPopup.style.top = "0"
    createPopup.style.left = "0"
    createPopup.style.height = "0"
    createPopup.style.width = "0"
    background.style.visibility = "hidden"
    let inputs = createPopup.querySelectorAll("#PopUpInput")
    for (let input of inputs) {
        input.value = ""
    }
    let button = createPopup.querySelector(".inputButton")
    button.removeEventListener("click", CreateGroup)


    setTimeout(()=>{
        createPopup.style.visibility = "hidden"
        window.removeEventListener("mousedown", CheckMouseClickPopup)
        window.addEventListener("click", CheckMouseClick)
        currentPopup = null
    }, 200)
}

function OpenDeleteGroup() {
    deletePopup.style.visibility = "visible"
    deletePopup.style.top = "30%"
    deletePopup.style.left = "30%"
    deletePopup.style.height = "30%"
    deletePopup.style.width = "40%"
    background.style.visibility = "visible"
    deletePopup.querySelector("#DeleteMessage").innerHTML = "Вы уверены? Чтобы подтвердить удаление, " +
        "введите точное название группы."
    let button = deletePopup.querySelector(".inputButton")
    button.addEventListener("click", DeleteGroup)
    CloseContextMenu()

    setTimeout(()=>{
        window.removeEventListener("click", CheckMouseClick)
        window.addEventListener("mousedown", CheckMouseClickPopup)
        currentPopup = deletePopup
    }, 200)
}

function CloseDeleteGroup() {
    deletePopup.style.top = "0"
    deletePopup.style.left = "0"
    deletePopup.style.height = "0"
    deletePopup.style.width = "0"
    background.style.visibility = "hidden"
    let inputs = deletePopup.querySelectorAll("#PopUpInput")
    for (let input of inputs) {
        input.value = ""
    }
    let button = deletePopup.querySelector(".inputButton")
    button.removeEventListener("click", DeleteGroup)

    setTimeout(()=>{
        deletePopup.style.visibility = "hidden"
        window.removeEventListener("mousedown", CheckMouseClickPopup)
        window.addEventListener("click", CheckMouseClick)
        currentPopup = null
    }, 200)
}

function CheckMouseClickPopup(click) {
    const target = click.target

    if (!target.closest(`#${currentPopup.id}`)) {
        currentPopup.ClosePopUp()
    }
}

/*Функции логики*/

function CreateGroup() {
    let name = createPopup.querySelector("#CreateName").value
    let description = createPopup.querySelector("#CreateDescription").value

    if (name === "") {
        createPopup.querySelector("#CreateMessage").innerHTML = "Введите корректное название"
        return
    }

    let data = {
        "Name": name,
        "Description": description,
    }
    POST(requestURL + "/private/groups/CreateGroup", data)
        .then((data)=>{
            if (data["Ok"]) {
                window.location.reload()
            } else {
                createPopup.querySelector("#CreateMessage").innerHTML = `Ошибка: ${data["Message"]}`
            }
        })
}

function DeleteGroup() {
    let name = deletePopup.querySelector("#DeleteName").value
    let id = selectedRow.id

    if (!CheckDelete(name)) {
        deletePopup.querySelector("#DeleteMessage").innerHTML = "Введите корректное название"
        return
    }

    let data = {
        "Id": parseInt(id)
    }
    POST(requestURL + "/private/groups/RemoveGroup", data)
        .then((data)=>{
            if (data["Ok"]) {
                window.location.reload()
            } else {
                deletePopup.querySelector("#DeleteMessage").innerHTML = `Ошибка: ${data["Message"]}`
            }
        })
}

function CheckDelete(text) {
    if (text === selectedRow.querySelector(".Name").innerText) {
        return true
    } else {
        return false
    }
}

/*Бинды*/
function BindPoUps() {
    createPopup.ClosePopUp = CloseCreateGroup
    deletePopup.ClosePopUp = CloseDeleteGroup
}

createGroup.addEventListener("click", OpenCreateGroup)
BindPoUps()