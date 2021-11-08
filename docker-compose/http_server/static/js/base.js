const requestURL = "https://" + document.domain + ":5000"
//TODO:Вставить нужный порт
const domain = document.domain

const archiveIcon = document.createElement("img")
archiveIcon.src = "/static/img/archive.png"
archiveIcon.classList.add("statusIcon")
archiveIcon.title = "В архиве"

const stockIcon = document.createElement("img")
stockIcon.src = "/static/img/storage.png"
stockIcon.classList.add("statusIcon")
stockIcon.title = "На складе"

const workIcon = document.createElement("img")
workIcon.src = "/static/img/computer.png"
workIcon.classList.add("statusIcon")
workIcon.title = "В работе"

async function POST(url, postData) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            "Content-type": "application/json"
        },
        body: JSON.stringify(postData)
    })
    if (response.status !== 200) {
        alert(response.statusText)
    }
    return await response.json();
}

async function GET(url) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        },
    })
    if (response.status !== 200) {
        alert(response.statusText)
    }
    return await response.json();
}

function FormatDate(date) {
    function pad(s) {return (s < 10) ? "0" + s: s;}
    return [pad(date.getDate()), pad(date.getMonth() + 1), date.getFullYear()]
        .join(".") + " " + [pad(date.getHours()), pad(date.getMinutes())].join(":")
}

function WriteBackUrl() {
    document.querySelector("#goBack").href = document.referrer
}