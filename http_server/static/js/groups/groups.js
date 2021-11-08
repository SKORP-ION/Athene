const tbody = document.querySelector(".TableBody")

function FormatDate(date) {
    function pad(s) {return (s < 10) ? "0" + s: s;}
    return [pad(date.getDate()), pad(date.getMonth() + 1), date.getFullYear()]
        .join(".") + " " + [pad(date.getHours()), pad(date.getMinutes())].join(":")
}

function GetNameFromUrl() {
    let params = new URLSearchParams(window.location.search)

    if (params.get("name")) {
        return params.get("name")
    }
    return ""
}

function WriteFormFromUrl() {
    document.querySelector("#searchName").value = GetNameFromUrl()
}

function OpenGroup(node) {
    let name = node.querySelector(".Name").innerText
    window.location.href = `/view/group?name=${name}`
}

function OpenGroupInNewTab(node) {
    let name = node.querySelector(".Name").innerText
    window.open(requestURL + `/view/group?name=${name}`)
}



GET(requestURL + `/private/groups/GetGroups?name=${GetNameFromUrl()}`)
    .then((data)=>{
        for (let row of data) {
            let tr = document.createElement("tr")
            tr.id = row["Id"]

            let name = document.createElement("td")
            name.classList.add("Name")
            name.innerHTML = row["Name"] || ""
            tr.append(name)

            let desc = document.createElement("td")
            desc.classList.add("Description")
            desc.innerHTML = row["Description"] || ""
            tr.append(desc)

            let creat = document.createElement("td")
            creat.classList.add("CreatedAt")
            creat.innerHTML = FormatDate(new Date(row["CreatedAt"]["seconds"] * 1000))
            tr.append(creat)

            let user = document.createElement("td")
            user.classList.add("Username")
            user.innerHTML = row["WhoDisplayName"] || ""
            tr.append(user)
            tr.addEventListener("mousedown", CheckRowActions)

            tbody.append(tr)
        }
    })

WriteFormFromUrl()