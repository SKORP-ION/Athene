async function getStates(url) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        }
    })
    return await response.json();
}

async function Is(url) {
    const response = await fetch(url, {
        method: 'GET',
        headers: {
            "Content-type": "application/json"
        },
    })
    return await response.json();
}

async function ActionRequest(url, postData) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            "Content-type": "application/json"
        },
        body: JSON.stringify(postData)
    })
    return await response.json();
}

function GetUrl() {
    switch (Action) {
        case 1: return "/private/history/CreateCard";
        case 3: return "/private/staff/GiveToEmployee";
        case 4: return "/private/staff/TakeFromEmployee";
        case 5: return "/private/history/InstallOnWorkspace";
        case 6: return "/private/history/RemoveFromWorkspace";
        case 7: return "/private/history/NeedRepair";
        case 8: return "/private/history/SendToRepair";
        case 9: return "/private/history/ReceiveFromRepair";
        case 10: return "/private/property/sendToWarehouse";
        case 11: return "/private/history/Archive";
        case 12: return "/private/history/ChangeInventory";
        case 13: return "/private/history/ChangeName";
        case 14: return "/private/history/DeArchive";
        case 15: return "/private/groups/AddPropsToGroup";
        case 16: return "/private/groups/RemoveFromGroup";
    }
}