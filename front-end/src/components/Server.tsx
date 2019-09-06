import React from 'react'

interface Props {
    serverCount: number;
}

const Server = (props: Props) => {

    const getItem = (i: number) => (
        <div key={i} className="p-2 card mt-5 mr-2" style={{ width: "18rem" }}>
            <div className="card-body">
                <h5 className="card-title">Server1</h5>
                <p className="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                <a href="#" className="btn btn-primary">Go somewhere</a>
            </div>
        </div>
    );
    const servers = () => {
        const serverItems = [];
        for (let j = 0; j < props.serverCount; j++) {
            serverItems.push(
                getItem(j)
            );
        }

        return serverItems;
    }

    return <div className="d-flex justify-content-center">
        {servers().map(item => item)}
    </div>;
}

export default Server;