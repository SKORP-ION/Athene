async function AddPropertyRequest(url, postData) {
    const response = await fetch(url, {
        method: 'POST',
        headers: {
            "Content-type": "application/json"
        },
        body: JSON.stringify(postData)
    })
    return await response.json();
}