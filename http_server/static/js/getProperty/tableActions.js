var selectedId = []
var startedRow
var selectedRows = []


function StartSelection(click) {
    //Средняя кнопка мыши
    if (click.button === 1) {
        return
    }

    //Если выделение началось правой кнопкой мыши
    if (click.button === 2 && this.classList.contains("selected")) {
        OpenContextMenu(click)
        return
    } else if (click.button === 2 && !this.classList.contains("selected")) {
        selectedId = []
        selectedRows = []
        ResetSelections()
        selectedId.push(this.propertyId)
        selectedRows.push(this)
        this.classList.add("selected")
        OpenContextMenu(click)
        return
    }
    //Событие для нажатой Ctrl
    if (click.ctrlKey && !this.classList.contains("selected")) {
        //Если элемент ещё не выделен
        this.classList.add("selected")
        if (!selectedId.includes(this.propertyId)) {
            selectedId.push(this.propertyId)
            selectedRows.push(this)
        }
        return
    } else if (click.ctrlKey && this.classList.contains("selected")) {
        //Если элемент выделен
        this.classList.remove("selected")

        for (i = 0; i < selectedId.length; i++) {
            let elem = selectedId[i]
            if (elem === this.propertyId) {
                selectedId.splice(i, 1)
                break
            }
        }
        for (i = 0; i < selectedRows.length; i++) {
            let row = selectedRows[i]
            if (row.isEqualNode(this)) {
                selectedRows.splice(i, 1)
                break
            }
        }
        return
    }

    selectedId = []
    selectedRows = []

    let rows = document.querySelectorAll(".PropertyTable > tr")
    for (i = 0; i < rows.length; i++) {
        let row = rows[i]
        row.addEventListener("mouseenter", SelectRow)
        row.addEventListener("mouseout", UnselectRow)
        row.classList.remove("selected")
    }
    this.classList.add("selected")
    startedRow = this
}

function SelectRow() {
    let startIndex = startedRow.rowIndex

    if (startIndex === this.rowIndex) {
        return
    }

    let row = this

    for (i = this.rowIndex; row.rowIndex > startedRow.rowIndex; i--) {
        row = document.querySelectorAll(".PropertyTable > tr")[i-1]
        row.classList.add("selected")
    }
    for (i = this.rowIndex; row.rowIndex < startedRow.rowIndex; i++) {
        row = document.querySelectorAll(".PropertyTable > tr")[i-1]
        row.classList.add("selected")
    }

}

function UnselectRow() {
    let startIndex = startedRow.rowIndex

    if (startIndex === this.rowIndex) {
        return
    }

    let row = this

    for (i = this.rowIndex; row.rowIndex < startedRow.rowIndex; i++) {
        row = document.querySelectorAll(".PropertyTable > tr")[i-1]
        row.classList.remove("selected")
    }
    for (i = this.rowIndex; row.rowIndex > startedRow.rowIndex; i--) {
        row = document.querySelectorAll(".PropertyTable > tr")[i-1]
        row.classList.remove("selected")
    }

}

function EndSelections(click) {
    if (click.button === 1) {
        OpenPropertyInNewTab(this)
        return
    } else if (click.button === 2) {
        OpenContextMenu(click)
        return
    }

    let rows = document.querySelectorAll(".PropertyTable > tr")
    for (i = 0; i < rows.length; i++) {
        let row = rows[i]
        row.removeEventListener("mouseenter", SelectRow)
        row.removeEventListener("mouseout", UnselectRow)
    }

    if (this === startedRow) {
        OpenProperty(this)
        return
    }

    rows = document.querySelectorAll("tr.selected")

    for (let row of rows) {
        if (!selectedId.includes(row.propertyId)) {
            selectedId.push(row.propertyId)
            selectedRows.push(row)
        }
    }

    console.log(selectedId)
}

function ResetSelections() {
    selectedId = []
    selectedRows = []
    let rows = document.querySelectorAll(".PropertyTable > tr")
    for (i = 0; i < rows.length; i++) {
        let row = rows[i]
        row.classList.remove("selected")
    }
}

function CheckMouseClick(click) {
    const target = click.target

    if (!target.closest(".PropertyTable") && (!target.closest("#ctxMenu"))) {
        ResetSelections()
        CloseContextMenu()
    } else if (!target.closest("#ctxMenu")) {
        click.preventDefault()
        CloseContextMenu()
    }
}

function OpenProperty(sender) {
    console.log("Open property")
    let id = sender.propertyId
    window.location.href = requestURL + `/view/property?id=${id}`
}

function OpenPropertyInNewTab(sender) {
    console.log("Open property in new tab")
    let id = sender.propertyId
    window.open(requestURL + `/view/property?id=${id}`)
}

window.addEventListener("click", CheckMouseClick)
let table = document.querySelector("#PropertyTable")
table.onmousedown = function (e) {if (e.button === 1) return false;}
table.oncontextmenu = function (e) {e.preventDefault()}