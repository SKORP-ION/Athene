
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

function CheckMouseClick(click) {
    const target = click.target

    if (!target.closest(".GroupsTable tr") && (!target.closest("#ctxMenu"))) {
        CloseContextMenu()
    } else if (!target.closest("#ctxMenu")) {
        click.preventDefault()
        CloseContextMenu()
    }
}

function BindFunctionsForMenus() {
    document.querySelector("#openProperty")
        .addEventListener("click", ()=>{OpenProperty(selectedRow.propertyId)})
    document.querySelector("#openPropertyInNewTab")
        .addEventListener("click", ()=>{OpenPropertyInNewTab(selectedRow.propertyId)})
    document.querySelector("#action").addEventListener("click", DoAction)
    document.querySelector("#printAction").addEventListener("click", OpenPrintPopup)
}

/*Функции для пунктов контекстного меню*/

function DoAction() {
    let prop = {}

    prop["id"] = selectedRow.propertyId
    prop["serial"] = selectedRow.querySelector(".Serial").innerText
    prop["inventory"] = selectedRow.querySelector(".Inventory").innerText
    prop["name"] = selectedRow.querySelector(".Name").innerText

    document.cookie = `properties=${JSON.stringify({"props": [prop]})}; path=/view/action; domain=${domain};`
    window.location.href = "/view/action"
}

window.addEventListener("click", CheckMouseClick)
document.querySelector("#GroupsTable").onmousedown = function (e) {if (e.button === 1) return false;}
document.querySelector("#GroupsTable").oncontextmenu = function (e) {e.preventDefault()}


BindFunctionsForMenus()