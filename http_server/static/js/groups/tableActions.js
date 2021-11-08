var selectedRow
function CheckMouseClick(click) {
    const target = click.target

    if (!target.closest(".GroupsTable") && (!target.closest("#ctxMenu"))) {
        CloseContextMenu()
    } else if (!target.closest("#ctxMenu")) {
        click.preventDefault()
        CloseContextMenu()
    }
}

function CheckRowActions(click) {
    selectedRow = this
    console.log(selectedRow.id)
    if (click.button === 0) {
        OpenGroup(this)
        return
    } else if (click.button === 1) {
        OpenGroupInNewTab(this)
        return
    } else if (click.button === 2) {
        OpenContextMenu(click)
        return
    }
}


window.addEventListener("click", CheckMouseClick)
let table = document.querySelector(".TableBody")
table.onmousedown = function (e) {if (e.button === 1) return false;}
table.oncontextmenu = function (e) {e.preventDefault()}