const CyrillicRegexp = /[А-ЯЁ]+|[а-яё]+/
const SpecSymbolRexexp = /[~!@#$&*=:\/,;?+]/
const HtmlRegexp = /<[A-Za-z]+>/

function CheckIsCreated(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            let name = node.parentNode.querySelector(".Name")
            let property = data["Property"]
            if (ok) {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/stop.png"
                img.title = "Уже существует карточка с таким серийным номером"
                a.href = `/view/property?id=${property["Id"]}`
                node.parentNode.toCreate = false
                node.parentNode.toAction = false
                a.append(img)
                status.append(a)
            } else {
                if (name.innerHTML === ""){
                    let a = document.createElement("a")
                    a.href = "#"
                    let img = document.createElement("img")
                    img.src = "/static/img/almost_ok.png"
                    img.title = "Проверка пройдена. Укажите название"
                    node.parentNode.toCreate = false
                    node.parentNode.toAction = false
                    a.append(img)
                    status.append(a)
                } else {
                    let a = document.createElement("a")
                    a.href = "#"
                    let img = document.createElement("img")
                    img.src = "/static/img/ok.png"
                    img.title = "Готово к добавлению"
                    node.parentNode.toCreate = true
                    node.parentNode.toAction = false
                    a.append(img)
                    status.append(a)
                }
            }
            CheckDuplicate(node)
        })
        .catch((error)=>{
            Error(status, error)
        })
}

function CheckIsNotCreated(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            let name = node.parentNode.querySelector(".Name")
            let property = data["Property"]
            if (!ok) {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/stop.png"
                img.title = "Карточка с таким серийным номером отсутствует"
                a.href = `/view/property?id=${property["Id"]}`
                node.parentNode.toCreate = false
                node.parentNode.toAction = false
                a.append(img)
                status.append(a)
            } else {
                if (name.innerHTML === ""){
                    let a = document.createElement("a")
                    a.href = "#"
                    let img = document.createElement("img")
                    img.src = "/static/img/almost_ok.png"
                    img.title = "Проверка пройдена. Укажите название"
                    node.parentNode.toCreate = false
                    node.parentNode.toAction = false
                    a.append(img)
                    status.append(a)
                } else {
                    let a = document.createElement("a")
                    a.href = "#"
                    let img = document.createElement("img")
                    img.src = "/static/img/ok.png"
                    img.title = "Готово к изменению"
                    node.parentNode.toCreate = false
                    node.parentNode.toAction = true
                    a.append(img)
                    status.append(a)
                }
            }
            CheckDuplicate(node)
        })
        .catch((error)=>{
            Error(status, error)
        })
}

function CheckIsOnWorkspace(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                Is(requestURL + `/private/history/isOnWorkspace?serial=${serial}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok !== true) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/warning.png"
                            img.title = "Нельзя снять с рабочего места то, что там не находится."
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию"
                            node.parentNode.toAction = true
                            a.append(img)
                            status.append(a)

                        }
                    })
            }
            CheckDuplicate(node)
        })
}

function CheckIsInStock(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                Is(requestURL + `/private/history/isInStock?serial=${serial}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (!ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/stop.png"
                            img.title = "Устройство не на складе. Подробнее по ссылке."
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию"
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = true
                            a.append(img)
                            status.append(a)

                        }
                    })
            }
            CheckDuplicate(node)
        })
}

function CheckIsNeedsRepair(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                Is(requestURL + `/private/history/isNeedsRepair?serial=${serial}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok !== true) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/warning.png"
                            img.title = "Не нуждается в ремонте в данный момент."
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = true
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию"
                            a.append(img)
                            status.append(a)

                        }
                    })
            }
            CheckDuplicate(node)
        })
}

function CheckIsUnderRepair(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                //IsUnderRepair
                Is(requestURL + `/private/history/isUnderRepair?serial=${serial}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok !== true) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/warning.png"
                            img.title = "Не нуждается в ремонте в данный момент."
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию"
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = true
                            a.append(img)
                            status.append(a)

                        }
                    })
            }
            CheckDuplicate(node)
        })
}

function CheckIsInArchive(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                //IsUnderRepair
                Is(requestURL + `/private/history/isInArchive?serial=${serial}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (!ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/stop.png"
                            img.title = "Не в архиве."
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию"
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = true
                            a.append(img)
                            status.append(a)

                        }
                    })
            }
            CheckDuplicate(node)
        })
}

function CheckIsInGroup(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/stop.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе."
                node.parentNode.toCreate = false
                node.parentNode.toAction = false
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                let group = document.querySelector("#groups").value
                //IsUnderRepair
                Is(requestURL + `/private/groups/IsInGroup?id=${property["Id"]}&group=${group}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/stop.png"
                            img.title = "Уже состоит в этой группе."
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию"
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = true
                            a.append(img)
                            status.append(a)
                        }
                    })
            }
        })
}

function CheckIsNotInGroup(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/stop.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе."
                node.parentNode.toCreate = false
                node.parentNode.toAction = false
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                tr.id = property["Id"]
                let group = document.querySelector("#groups").value
                //IsUnderRepair
                Is(requestURL + `/private/groups/IsInGroup?id=${property["Id"]}&group=${group}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (!ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/stop.png"
                            img.title = "Не состоит в этой группе."
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к удалению"
                            node.parentNode.toCreate = false
                            node.parentNode.toAction = true
                            a.append(img)
                            status.append(a)
                        }
                    })
            }
        })
}

function CheckIsOnWarehouse(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                let id = property["Id"]
                tr.id = id
                let warehouse = document.querySelector("#warehouses").value

                Is(requestURL + `/private/property/isOnWarehouse?id=${id}&warehouse=${warehouse}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/stop.png"
                            img.title = "Уже на этом складе."
                            node.parentNode.toAction = false
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            if (data["Message"] === "Record not found"){
                                let a = document.createElement("a")
                                a.href = `/view/property?id=${property["Id"]}`
                                let img = document.createElement("img")
                                img.src = "/static/img/ok.png"
                                img.title = "Готово к действию"
                                node.parentNode.toAction = true
                                a.append(img)
                                status.append(a)
                            } else if (data["Message"] === "Not in stock") {
                                let a = document.createElement("a")
                                a.href = `/view/property?id=${property["Id"]}`
                                let img = document.createElement("img")
                                img.src = "/static/img/stop.png"
                                img.title = "В данный момент используется. Сначала нужно забрать на склад."
                                node.parentNode.toAction = false
                                a.append(img)
                                status.append(a)
                            }
                        }
                    })
            }
        })
}

function CheckIsReadyToGive(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                let id = property["Id"]
                tr.id = id
                let employee = document.querySelector("#employee").employee

                if (!("Id" in employee)) {
                    let a = document.createElement("a")
                    a.href = `/view/property?id=${property["Id"]}`
                    let img = document.createElement("img")
                    img.src = "/static/img/warning.png"
                    img.title = "Пользователь еще не существует в базе. Будет создана новая запись."
                    node.parentNode.toAction = true
                    a.append(img)
                    status.append(a)
                    return
                }

                Is(requestURL + `/private/history/isInStock?serial=${serial}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию."
                            node.parentNode.toAction = true
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            let a = document.createElement("a")
                            a.href = `/view/property?id=${property["Id"]}`
                            let img = document.createElement("img")
                            img.src = "/static/img/stop.png"
                            img.title = "Не на складе. Подробнее по ссылке."
                            node.parentNode.toAction = false
                            a.append(img)
                            status.append(a)
                        }
                    })
            }
        })
}

function CheckIsWithEmployee(node) {
    let status = node.parentNode.querySelector(".Status")
    status.innerHTML = ""
    let serial = node.parentNode.querySelector(".Serial").innerHTML

    if (["", " "].includes(serial)) {
        let a = document.createElement("a")
        a.href = "#"
        let img = document.createElement("img")
        img.src = "/static/img/error.png"
        img.title = "Укажите нормальный серийный номер!"
        node.parentNode.toCreate = false
        node.parentNode.toAction = false
        a.append(img)
        status.append(a)
        return
    }

    Is(requestURL + `/private/history/isCreated?serial=${serial}`)
        .then((data)=>{
            let ok = data["Ok"]
            if (!ok) {
                let a = document.createElement("a")
                let img = document.createElement("img")
                img.src = "/static/img/warning.png"
                img.title = "Имущество с таким серийным номером отсутствует в базе. Будет создана новая запись."
                node.parentNode.toCreate = true
                node.parentNode.toAction = true
                a.href = `#`
                a.append(img)
                status.append(a)
            } else {
                FixRow(node, data)
                if (IsLastRow(node)) {
                    CreateRow()
                }
                let property = data["Property"]
                let tr = node.parentNode
                let id = property["Id"]
                tr.id = id
                let employee = document.querySelector("#employee").employee

                if (!("Id" in employee)) {
                    let a = document.createElement("a")
                    a.href = `/view/property?id=${property["Id"]}`
                    let img = document.createElement("img")
                    img.src = "/static/img/stop.png"
                    img.title = "Пользователь еще не существует в базе."
                    node.parentNode.toAction = false
                    a.append(img)
                    status.append(a)
                    return
                }

                Is(requestURL + `/private/staff/isWithEmployee?id=${id}&employee=${employee["Id"]}`)
                    .then((data)=>{
                        let ok = data["Ok"]
                        if (ok) {
                            let a = document.createElement("a")
                            let img = document.createElement("img")
                            img.src = "/static/img/ok.png"
                            img.title = "Готово к действию."
                            node.parentNode.toAction = true
                            a.href = `/view/property?id=${property["Id"]}`
                            a.append(img)
                            status.append(a)
                        } else {
                            if (data["Message"] === "Record not found"){
                                let a = document.createElement("a")
                                a.href = `/view/property?id=${property["Id"]}`
                                let img = document.createElement("img")
                                img.src = "/static/img/stop.png"
                                img.title = "Не выдано."
                                node.parentNode.toAction = false
                                a.append(img)
                                status.append(a)
                            } else if (data["Message"].includes("Last owner")) {
                                let a = document.createElement("a")
                                a.href = `/view/property?id=${property["Id"]}`
                                let img = document.createElement("img")
                                img.src = "/static/img/stop.png"
                                img.title = "Выдано другому пользователю."
                                node.parentNode.toAction = false
                                a.append(img)
                                status.append(a)
                            }
                        }
                    })
            }
        })
}

function CheckDuplicate(node) {
    let rows = tablebody.querySelectorAll("tr > .Serial")
    let number = 0

    for (let row of rows) {
        if (node.innerHTML === row.innerHTML) {
            number += 1
        }
    }

    if (number > 1) {
        setTimeout(()=>{
            Error(node.parentNode.querySelector(".Status"), "Duplicate", "Дубликат")
        }, 500)

    }
}

function CheckCyrillic(node) {
    let serial = node.innerText
    let matches = serial.match(CyrillicRegexp)

    if (matches) {
        Error(node.parentNode.querySelector(".Status"), "match Cyrillic",
            "В поле ввода обнаружена кириллица")
        return true
    }
    return false
}

function CheckSpecSymbols(node) {
    let serial = node.innerText
    let matches = serial.match(SpecSymbolRexexp)

    if (matches) {
        Error(node.parentNode.querySelector(".Status"), "match spec symbols",
            "Зарпещено использовать следующие спецсимволы для серийного номера: [~!@#$&*=:/,;?+]")
        return true
    }
    return false
}

function FixRow(node, data) {
    let property = data["Property"]
    let serial = node.parentNode.querySelector(".Serial")
    let inventory = node.parentNode.querySelector(".Inventory")
    let name = node.parentNode.querySelector(".Name")
    node.parentNode.id = property["Id"]

    serial.classList.add("fixed")
    serial.classList.remove("Editable")
    serial.contentEditable = false

    inventory.classList.add("fixed")
    inventory.innerHTML = property["Inventory"] || ""
    inventory.classList.remove("Editable")
    inventory.contentEditable = false

    name.classList.add("fixed")
    name.classList.remove("Editable")
    name.innerHTML = property["Name"] || ""
    inventory.contentEditable = false
    AddRemove(node)
}

function AddRemove(node) {
    let remove = node.parentNode.querySelector(".Remove")
    remove.innerHTML = ""
    let img = document.createElement("img")
    img.src = "/static/img/remove.png"
    img.classList.add("statusIcon")
    img.addEventListener("click", RemoveRow)
    remove.append(img)
}

function Error(status, error, message = "Error") {
    status.parentNode.toCreate = false
    status.parentNode.toAction = false
    console.log(error)
    let a = document.createElement("a")
    a.href = "#"
    let img = document.createElement("img")
    img.src = "/static/img/error.png"
    img.title = message
    a.append(img)
    status.innerHTML = ""
    status.append(a)
}

function IsLastRow(node) {
    let elements = tablebody.querySelectorAll("tr")

    if (elements[elements.length - 1].isEqualNode(node.parentNode)) {
        return true
    } else {
        return false
    }
}
