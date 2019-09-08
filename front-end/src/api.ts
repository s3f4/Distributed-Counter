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
            return {
                error: err
            };
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

export const shutdownNode = (pID: number) => {
    return fetch(`http://127.0.0.1:3001/shutdown/${pID}`, {
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

export const getCount = (tenantID: number) => {
    return fetch(`http://127.0.0.1:3001/items/${tenantID}/count`, {
        method: "GET"
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