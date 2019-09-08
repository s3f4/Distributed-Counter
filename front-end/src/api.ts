export const insertItem = (item: any) => {
    return fetch(`http://127.0.0.1:3001/items`, {
        method: "POST",
        headers: {
            Accept: 'application/json',
        },
        body: JSON.stringify(item)
    })
        .then(response => {
            return response.json();
        })
        .catch(err => {
            console.log(err);
        });
}

export const getNodeDB = (port: number) => {
    return fetch(`http://127.0.0.1:${port}/database`, {
        method: "GET",
    })
        .then(response => {
            return response.json();
        })
        .catch(err => {
            return {
                error: err
            };
        });
}

export const getNodes = () => {
    return fetch(`http://127.0.0.1:3001/nodes`, {
        method: "GET",
    })
        .then(response => {
            return response.json();
        })
        .catch(err => {
            return {
                error: err
            };
        });
}

export const upNodes = (nodeCount: number) => {
    return fetch(`http://127.0.0.1:3001/upNodes/${nodeCount}`, {
        method: "GET",
    })
        .then(response => {
            return response.json();
        })
        .catch(err => {
            return {
                error: err
            };
        });
}

export const getCount = () => {
    return fetch(`http://127.0.0.1:3001/count`, {
        method: "GET"
    })
        .then(response => {
            return response.json();
        })
        .catch(error => console.log(error));
}