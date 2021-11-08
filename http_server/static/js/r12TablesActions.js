var selectedNewProperties = []
var selectedOldProperties = []
var ScrollOffset = 0
var ScrollTarget = null
var PreviousData = {}

/*------Выделения------*/
function StartSelection(click) {
    //Средняя кнопка мыши
    if (click.button === 1) {
        return
    }

    //Если выделение началось правой кнопкой мыши
    if (click.button === 2 && this.classList.contains("selected")) {
        EndSelections(click)
        return
    }

    //Если элемент помечен завершенным или ошибкой
    if (this.classList.contains("completed") || this.classList.contains("error")) {
        return
    }

    //Событие для нажатой Ctrl
    if (click.ctrlKey && !this.classList.contains("selected")) {
        //Если элемент уже выделен
        this.classList.add("selected")
        if (this.id.includes("New")) {
            let property = GetPropertyData(this.id)
            if (!selectedNewProperties.join().includes(property.join())) {
                selectedNewProperties.push(property)
            }
        } else if (this.id.includes("Old")) {
            selectedOldProperties.push([
                document.querySelector(`#${this.id} > .Inventory`).innerText,
                document.querySelector(`#${this.id} > .Serial`).innerText,
                document.querySelector(`#${this.id} > .Name`).innerText
            ])
        }
        return
    } else if (click.ctrlKey && this.classList.contains("selected")) {
        //Если элемент еще не выделен
        this.classList.remove("selected")
        let inventory = this.querySelector(".Inventory").innerText

        if (this.id.includes("New")) {
            for (i = 0; i < selectedNewProperties.length; i++) {
                let elem = selectedNewProperties[i]
                if (elem[0] === inventory) {
                    selectedNewProperties.splice(i, 1)
                    break
                }
            }
        } else if (this.id.includes("Old")) {
            for (i = 0; i < selectedOldProperties.length; i++) {
                let elem = selectedOldProperties[i]
                if (elem[0] === inventory) {
                    selectedOldProperties.splice(i, 1)
                    break
                }
            }
        }
        return
    }

    selectedNewProperties = []
    selectedOldProperties = []

    function ResetRows(selectors) {
        let rows = document.querySelectorAll(selectors)
        for (i = 0; i < rows.length; i++) {
            let row = rows[i]
            row.addEventListener("mouseenter", SelectRow)
            row.classList.remove("selected")
        }
    }

    ResetRows(".ContentTable > tbody > tr")
    if (this.id.includes("New")) {

        this.classList.add("selected")
        let property = GetPropertyData(this.id)
        if (!selectedNewProperties.join().includes(property.join())) {
            selectedNewProperties.push(property)
        }
        ScrollTarget = document.querySelector("#NewTableBody")
        ScrollTarget.addEventListener("mousemove", SetScrollOffset)
        ScrollOffset = 0
    } else if (this.id.includes("Old")) {
        this.classList.add("selected")
        let property = GetPropertyData(this.id)
        if (!selectedOldProperties.join().includes(property.join())){
            selectedOldProperties.push(property)
        }
        ScrollTarget = document.querySelector("#OldTableBody")
        ScrollTarget.addEventListener("mousemove", SetScrollOffset)
        ScrollOffset = 0
    }
    if (click.button === 2) {
        EndSelections(click)
    }


}

function SelectRow() {

    if (this.classList.contains("completed") || this.classList.contains("error")) {
        return
    }


    this.classList.add("selected")
    if (this.id.includes("New")) {
        let property = GetPropertyData(this.id)
        if (!selectedNewProperties.join().includes(property.join())) {
            selectedNewProperties.push(property)
        }
    } else if (this.id.includes("Old")) {
        let property = GetPropertyData(this.id)
        if (!selectedOldProperties.join().includes(property.join())){
            selectedOldProperties.push(property)
        }
    }
}

function EndSelections(click) {
    if (click.button === 2) {
        OpenContextMenu(click)
    }

    this.removeEventListener("mousemove", Scroll)

    let rows = document.querySelectorAll(".ContentTable > tbody > tr")
    for (i = 0; i < rows.length; i++) {
        let row = rows[i]
        row.removeEventListener("mouseenter", SelectRow)
    }

    console.log(selectedNewProperties || selectedOldProperties)

    let tables = document.querySelectorAll(".TableBody")

    for (let table of tables) {
        table.removeEventListener("mousemove", SetScrollOffset)
    }
    ScrollTarget = null
}

function ResetSelections() {
    selectedNewProperties = []
    selectedOldProperties = []
    let rows = document.querySelectorAll("table > tbody > tr")
    for (i = 0; i < rows.length; i++) {
        let row = rows[i]
        row.classList.remove("selected")
    }
}

function GetPropertyData(rowId) {
    return [
        document.querySelector(`#${rowId} > .Inventory`).innerText,
        document.querySelector(`#${rowId} > .Serial`).innerText,
        document.querySelector(`#${rowId} > .Name`).innerText
    ]
}

/*------Скроллинг------*/
function Scroll() {
    if (ScrollTarget !== null && ScrollOffset !== 0) {
        ScrollTarget.scrollTop += ScrollOffset
    }
}

function SetScrollOffset(click) {
    let height = this.clientHeight
    let bottomPoint = height + this.offsetTop

    if (click.pageY > bottomPoint - 100) {
         offset = 200 / (bottomPoint - click.pageY )
         ScrollOffset = Math.min(offset, 10)
    } else if (click.pageY < this.offsetTop + 100) {
        offset = 200 / (this.offsetTop - click.pageY)
        ScrollOffset = Math.max(offset, -10)
    } else {
        ScrollOffset = 0
    }

}

/*------Функции действий------*/
function SendProperty(note="") {
    let data = {
        "note": note,
        "properties": []
    }
    for (let prop of selectedNewProperties) {
        data["properties"].push({
            "inventory": parseInt(prop[0]),
            "serial": prop[1],
            "name": prop[2],
        })
    }
    PreviousData = data
    AddPropertyRequest(requestURL + "/private/history/AddProperty", data)
        .then((response) => {
            if(response === "Success") {
                for (let line of selectedNewProperties) {
                    let row = document.querySelector(`#row${line[0]}New`)
                    row.classList.add("completed")
                    row.classList.remove("selected")
                }
            } else {
                for (let line of selectedNewProperties) {
                    let row = document.querySelector(`#row${line[0]}New`)
                    row.classList.add("error")
                    row.classList.remove("selected")
                }
            }
            selectedNewProperties = []
        })
}

function SendToArchive(note = "") {

}

function Skip() {
    //Удаление выделенных элементов
    for (i = 0; i < selectedNewProperties.length; i++) {
        let elem = selectedNewProperties[i]

        let row = document.querySelector(`#row${elem[0]}New`)
        row.remove()
    }

    for (i = 0; i < selectedOldProperties.length; i++) {
        let elem = selectedOldProperties[i]

        let row = document.querySelector(`#row${elem[0]}Old`)
        row.remove()
    }
    selectedNewProperties = []
    selectedOldProperties = []
}

function OpenPopUpWindow(click) {
    let popup = document.querySelector("#PopUpWindow")
    if (this.id === "saveWithNote") {
        popup.querySelector("#PopUpButton").addEventListener("click", ()=>{
            let message = popup.querySelector("#PopUpInput").value
            SendProperty(message)
            ClosePopUpWindow()
        })
    } else if (this.id === "sendToArchive") {
        popup.querySelector("#PopUpButton").addEventListener("click", ()=>{
            let message = popup.querySelector("#PopUpInput").value
            SendToArchive(message)
            ClosePopUpWindow()
        })
    }
    popup.style.visibility = "visible"
    popup.style.top = "30%"
    popup.style.left = "30%"
    popup.style.height = "30%"
    popup.style.width = "40%"
    let background = document.querySelector("#BlackBackGround")
    background.style.visibility = "visible"

    setTimeout(()=>{
        window.removeEventListener("click", CheckMouseClick)
        window.addEventListener("click", CheckMouseClickPopup)
    }, 200)
}

function ClosePopUpWindow() {
    let popup = document.querySelector("#PopUpWindow")
    popup.style.top = "0"
    popup.style.left = "0"
    popup.style.height = "0"
    popup.style.width = "0"
    let background = document.querySelector("#BlackBackGround")
    background.style.visibility = "hidden"
    let input = popup.querySelector("#PopUpInput")
    input.value = ""

    setTimeout(()=>{
        popup.style.visibility = "hidden"
        window.removeEventListener("click", CheckMouseClickPopup)
        window.addEventListener("click", CheckMouseClick)
    }, 200)
}

function SelectAll() {
    let ctx = document.querySelector("#ctxMenu")
    for (let row of document.querySelectorAll(`${ctx.currentTableSelector}  tr`)) {
        let property = GetPropertyData(row.id)

        if (ctx.currentTableSelector === ".AddedTable" &&
            !row.classList.contains("selected") &&
            !row.classList.contains("completed") &&
            !row.classList.contains("error") &&
            !selectedNewProperties.join().includes(property.join())){
            selectedNewProperties.push(property)
            setTimeout(()=>{row.classList.add("selected")}, 0)
        } else if (ctx.currentTableSelector === ".DeleteTable" &&
            !row.classList.contains("selected") &&
            !row.classList.contains("completed") &&
            !row.classList.contains("error") &&
            !selectedOldProperties.join().includes(property.join())){
            selectedOldProperties.push(property)
            setTimeout(()=>{row.classList.add("selected")}, 0)
        }
    }
}

/*------Контекстное меню------*/
function OpenContextMenu(click) {
    let ctx = document.querySelector("#ctxMenu")
    ctx.oncontextmenu = function (e) {e.preventDefault()}
    ctx.style.visibility = "visible"
    ctx.style.left = (click.pageX) + "px"
    ctx.style.top = (click.pageY) + "px"
    let menus = ctx.querySelectorAll(".menu")

    for (let menu of menus) {
        menu.style.height = "30px";
        menu.style.width = "250px";
    }

    newProperty = document.querySelector(".NewProperty")
    oldProperty = document.querySelector(".OldProperty")
    if (click.path.includes(newProperty)) {
        let archive = document.querySelector("#ctxMenu > #sendToArchive")
        //archive.style.visibility = "hidden";
        archive.style.display = "none";
        let saveWithNote = document.querySelector("#ctxMenu > #saveWithNote")
        saveWithNote.style.display = "block";
        let saveWithoutNote = document.querySelector("#ctxMenu > #saveWithoutNote")
        saveWithoutNote.style.display = "block";
        ctx.currentTableSelector = ".AddedTable"
    } else if (click.path.includes(oldProperty)) {
        let archive = document.querySelector("#ctxMenu > #sendToArchive")
        archive.style.display = "block";
        let saveWithNote = document.querySelector("#ctxMenu > #saveWithNote")
        saveWithNote.style.display = "none";
        let saveWithoutNote = document.querySelector("#ctxMenu > #saveWithoutNote")
        saveWithoutNote.style.display = "none";
        ctx.currentTableSelector = ".DeleteTable"
    }
}

function CloseContextMenu(click) {
    let ctx = document.querySelector("#ctxMenu")
    let menus = ctx.querySelectorAll(".menu")

    for (let menu of menus) {
        menu.style.height = "0";
        menu.style.width = "0";
        menu.senderId = ""
    }
    setTimeout(()=>{ctx.style.visibility = "hidden"}, 100)

}


/*------Остальное------*/
function CheckMouseClick(click) {
    const target = click.target
    CloseContextMenu(click)

    /*
    if (!target.closest("#PopUpWindow") &&
        document.querySelector("#PopUpWindow").style.visibility === "visible") {
        ClosePopUpWindow()
        return
    }
    */

    if (!target.closest(".TableBody") && !target.closest("#ctxMenu")) {
        ResetSelections()
    }
}

function CheckMouseClickPopup(click) {
    const target = click.target

    if (!target.closest("#PopUpWindow")) {
        ClosePopUpWindow()
    }
}

window.addEventListener("click", CheckMouseClick)
let tables = document.querySelectorAll(".TableContainer")
for (i = 0; i < tables.length; i++) {
    let table = tables[i]
    table.onmousedown = function (e) {if (e.button === 1) return false;}
    table.oncontextmenu = function (e) {e.preventDefault()}
}

let rows = document.querySelectorAll(".ContentTable > tbody > tr")
for (i = 0; i < rows.length; i++) {
    let row = rows[i]
    row.onmousedown = StartSelection
    row.onmouseup = EndSelections
}

document.querySelector("#selectAll").addEventListener("click", SelectAll)
document.querySelector("#saveWithNote").addEventListener("click", OpenPopUpWindow)
document.querySelector("#sendToArchive").addEventListener("click", OpenPopUpWindow)
document.querySelector("#saveWithoutNote").addEventListener("click", ()=>{SendProperty()})
document.querySelector("#skip").addEventListener("click", Skip)


setInterval(Scroll, 10)