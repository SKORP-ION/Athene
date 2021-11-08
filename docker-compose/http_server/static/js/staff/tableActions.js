var selectedRow

function CheckEvent(e) {
    selectedRow = this
    if (e.button === 0) {
        OpenProperty(this.propertyId)
    } else if (e.button === 1) {
        OpenPropertyInNewTab(this.propertyId)
    } else if (e.button === 2) {
        OpenContextMenu(e)
    }
}

function OpenProperty(propId) {
    window.location.href = requestURL + `/view/property?id=${propId}`
}

function OpenPropertyInNewTab(propId) {
    window.open(requestURL + `/view/property?id=${propId}`)
}