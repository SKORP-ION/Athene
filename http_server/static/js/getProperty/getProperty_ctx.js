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
}

function CloseContextMenu() {
    let ctx = document.querySelector("#ctxMenu")
    let menus = ctx.querySelectorAll(".menu")

    for (let menu of menus) {
        menu.style.height = "0";
        menu.style.width = "0";
        menu.senderId = ""
    }
    setTimeout(()=>{ctx.style.visibility = "hidden"}, 100)

}

function OpenCopySubmenu() {
    let subCtx = document.querySelector("#ctxMenu .subCtx")

    subCtx.style.visibility = "visible"
    document.addEventListener("mousemove", CheckMousePosition)
}

function CloseCopySubmenu() {
    let subCtx = document.querySelector("#ctxMenu .subCtx")

    subCtx.style.visibility = "hidden"
}

function CheckMousePosition(event) {
    let target = event.target

    if (!target.closest("#ctxMenu .subCtx") && !target.closest("#copySelect") ) {
        event.preventDefault()
        CloseCopySubmenu()
        document.removeEventListener("mousemove", CheckMousePosition)
    }

}

function BindFunctionsForMenus() {
    document.querySelector("#openProperty").addEventListener("click", Open)
    document.querySelector("#openInNewTab").addEventListener("click", OpenInNewTab)
    document.querySelector("#selectAll").addEventListener("click", SelectAll)
    document.querySelector("#copySelect").addEventListener("mouseenter", OpenCopySubmenu)
    document.querySelector("#copySelected").addEventListener("click", CopyAll)
    document.querySelector("#copyInventory").addEventListener("click", CopyInventories)
    document.querySelector("#copySerial").addEventListener("click", CopySerials)
    document.querySelector("#copyName").addEventListener("click", CopyNames)
    document.querySelector("#action").addEventListener("click", DoAction)
}

/*Функции для пунктов контекстного меню*/

function CopyAll() {
    let message = ""

    for (let row of selectedRows) {
        let action = row.querySelector("img").title || ""
        let inventory = row.querySelector(`#${row.id}inventory`).innerText || ""
        let serial = row.querySelector(`#${row.id}serial`).innerText || ""
        let name = row.querySelector(`#${row.id}name`).innerText || ""
        let created_at = row.querySelector(`#${row.id}created_at`).innerText || ""
        let updated_at = row.querySelector(`#${row.id}updated_at`).innerText || ""
        let warehouse = row.querySelector(`#${row.id}warehouse`).innerText || ""
        message += `${action}\t${inventory}\t${serial}\t${name}\t${created_at}\t${updated_at}\t${warehouse}\n`
    }
    try{
        navigator.clipboard.writeText(message)
    } catch {
        alert("Копирование не будет доступно, пока сайт не заработает по https")
    }
    CloseCopySubmenu()
    CloseContextMenu()
}

function CopyInventories() {
    let message = ""

    for (let row of selectedRows) {
        let inventory = row.querySelector(`#${row.id}inventory`).innerText || ""
        message += `${inventory}\n`
    }
    try{
        navigator.clipboard.writeText(message)
    } catch {
        alert("Копирование не будет доступно, пока сайт не заработает по https")
    }
    CloseCopySubmenu()
    CloseContextMenu()
}

function CopySerials() {
    let message = ""

    for (let row of selectedRows) {
        let serial = row.querySelector(`#${row.id}serial`).innerText || ""
        message += `${serial}\n`
    }
    try{
        navigator.clipboard.writeText(message)
    } catch {
        alert("Копирование не будет доступно, пока сайт не заработает по https")
    }
    CloseCopySubmenu()
    CloseContextMenu()
}

function CopyNames() {
    let message = ""

    for (let row of selectedRows) {
        let name = row.querySelector(`#${row.id}name`).innerText || ""
        message += `${name}\n`
    }
    try{
        navigator.clipboard.writeText(message)
    } catch {
        alert("Копирование не будет доступно, пока сайт не заработает по https")
    }
    CloseCopySubmenu()
    CloseContextMenu()
}

function SelectAll(click) {
    let rows = document.querySelectorAll(".PropertyTable > tr")
    selectedRows = []
    selectedId = []

    for (let row of rows) {
        row.classList.add("selected")
        selectedRows.push(row)
        selectedId.push(row.propertyId)
    }
    CloseContextMenu(click)
}

function Open() {
    console.log("Open property")
    let id = selectedId[0]
    window.location.href = requestURL + `/view/property?id=${id}`
    CloseContextMenu()
}

function OpenInNewTab() {
    for (let id of selectedId) {
        window.open(requestURL + `/view/property?id=${id}`)
    }
    CloseContextMenu()
}

BindFunctionsForMenus()