const printPopup = document.querySelector("#PrintPopUp")
const background = document.querySelector("#BlackBackGround")
var currentPopup

/*Функционал появления и исчезания меню*/
function OpenPrintPopup() {
    printPopup.style.visibility = "visible"
    printPopup.style.top = "30%"
    printPopup.style.left = "30%"
    printPopup.style.height = "30%"
    printPopup.style.width = "40%"
    background.style.visibility = "visible"
    printPopup.querySelector("#PrintMessage").innerHTML = "Введите номер заявки"

    printPopup.querySelector("#NameI").value = selectedRow.querySelector(".Name").innerText

    let button = printPopup.querySelector(".inputButton")
    button.addEventListener("click", PrintAction)
    CloseContextMenu()

    setTimeout(()=>{
        window.removeEventListener("click", CheckMouseClick)
        window.addEventListener("mousedown", CheckMouseClickPopup)
        currentPopup = printPopup
    }, 200)
}

function ClosePrintPopup() {
    printPopup.style.top = "0"
    printPopup.style.left = "0"
    printPopup.style.height = "0"
    printPopup.style.width = "0"
    background.style.visibility = "hidden"
    let inputs = printPopup.querySelectorAll("#PopUpInput")
    for (let input of inputs) {
        input.value = ""
    }
    let button = printPopup.querySelector(".inputButton")
    button.removeEventListener("click", PrintAction)


    setTimeout(()=>{
        printPopup.style.visibility = "hidden"
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

function PrintAction() {
    let ticket = printPopup.querySelector("#Ticket").value
    let name = printPopup.querySelector("#NameI").value
    let recordId = selectedRow.recordId

    if (ticket === "") {
        printPopup.querySelector("#CreateMessage").innerHTML = "Введите номер заявки. Без заявки никак."
        return
    }

    let form = document.createElement("form")
    form.method = "POST"
    form.style.display = "none"
    form.action = "/view/print"
    form.name = "print"

    let tic = document.createElement("input")
    tic.name = "ticket"
    tic.value = ticket
    form.appendChild(tic)

    let nm = document.createElement("input")
    nm.name = "name"
    nm.value = name
    form.appendChild(nm)

    let rec = document.createElement("input")
    rec.name = "recordId"
    rec.value = `${recordId}`
    form.appendChild(rec)

    document.body.appendChild(form)
    form.submit()

    /*
    POST(requestURL + "/view/print", data)
        .then((data)=>{
            if (data["Ok"]) {
                window.location.reload()
            } else {
                printPopup.querySelector("#CreateMessage").innerHTML = `Ошибка: ${data["Message"]}`
            }
        })
    */
}

/*Бинды*/
function BindPopUps() {
    printPopup.ClosePopUp = ClosePrintPopup
}

BindPopUps()