var timeout = new Date()
var needSearch = false
var symbolNumber = 0

async function PassiveObserver() {
    let input = document.querySelector("#employee")
    let datalist = document.querySelector("#staff")
    while (true) {
        let delta = new Date() - timeout

        if (delta >= 200 && needSearch && symbolNumber >= 3) {
            needSearch = false
            POST(requestURL + "/private/staff/GetStaff", {"Search": input.value})
                .then(function (response) {
                    let ul = document.createElement("ul")
                    let staff = response["Staff"]
                    datalist.innerHTML = ""
                    datalist.append(ul)
                    for (let employee of staff) {
                        let li = document.createElement("li")
                        li.innerText = employee["Name"]
                        li.title = `${employee["Name"]} - ${employee["JobTitle"]} из ${employee["Department"]}`
                        li.addEventListener("mousedown", FixEmployee)
                        li.employee = employee
                        ul.append(li)
                    }
                    ShowDatalist()
                })
        } else if (symbolNumber < 3) {
            datalist.innerHTML = ""
            HideDatalist()
        }
        await new Promise(r => setTimeout(r, 200))
    }
}

function ChangeValue() {
    needSearch = true
    timeout = new Date()
    symbolNumber = this.value.length
}

function FixEmployee() {
    let input = document.querySelector("#employee")
    input.employee = this.employee
    input.value = input.employee["Name"]
    input.title = `${input.employee["Name"]} - ${input.employee["Job"]} из ${input.employee["Department"]}`
    ShowEmployeeInfo()
    ShowTable()
}

function ClearEmployee() {
    let input = document.querySelector("#employee")
    input.employee = null
    input.value = ""
    input.title = ""
    document.querySelector(".datalist").innerHTML = ""
}

function ShowDatalist() {
    let input = document.querySelector("#employee")
    if (!input.classList.contains("fixed")) {
        let datalist = document.querySelector("#staff")
        datalist.style.visibility = "visible"
        datalist.style.top = `${input.getBoundingClientRect().top + window.scrollY + 40}px`
        datalist.style.left = `${input.getBoundingClientRect().left + window.scrollX}px`


        setTimeout(()=>{
            datalist.style.height = "auto"
        }, 20)
    }

}

function HideDatalist() {
    let datalist = document.querySelector("#staff")
    datalist.style.height = ""

    setTimeout(()=>{
        datalist.style.visibility = "hidden"
    }, 100)
}

document.querySelector("#employee").addEventListener("keyup", ChangeValue)
document.querySelector("#employee").addEventListener("focusin", ShowDatalist)
document.querySelector("#employee").addEventListener("focusout", HideDatalist)
PassiveObserver()