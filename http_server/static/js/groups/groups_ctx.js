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

function BindFunctionsForMenus() {
    document.querySelector("#action").addEventListener("click", DoAction)
    document.querySelector("#delete").addEventListener("click", OpenDeleteGroup)
}

/*Функции для пунктов контекстного меню*/

function DoAction() {
    let name = selectedRow.querySelector(".Name").innerText
    POST(requestURL + "/private/property/getProperty", {"groups": [name]})
        .then((data)=>{
            let props = []
            for (let prop of data["Properties"]) {
                props.push({
                    "id": prop["Id"],
                    "serial": prop["Serial"] || "",
                    "inventory": prop["Inventory"] || "",
                    "name": prop["Name"] || ""
                })
            }
            document.cookie = `properties=${JSON.stringify({"props": props})}; path=/view/action; domain=${domain};`
            window.location.href = "/view/action"
        })
}

BindFunctionsForMenus()