import React from 'react'
import { getNodeDB } from '../api';

interface Item {
    TenantID: string;
    ItemID: string;
}

const colors = [
    '#e6194b', '#3cb44b', '#ffe119', '#4363d8', '#f58231', '#911eb4', '#46f0f0', '#f032e6', '#bcf60c', '#fabebe',
];

const NodeCard = (props: any) => {
    const { node } = props;
    const [error, setError] = React.useState<any>(false);
    const [db, setDB] = React.useState<any>([])

    React.useEffect(() => {
        getNodeDB(node.Port).then(data => {
            if (data.error) {
                setError(data.error.message);
            } else {
                setError(false);
                setDB(data);
            }
        });
    }, [node])


    return (
        <div className="p-2 card mt-5 mr-2" style={{ width: "18rem" }}>
            <div className="card-body">
                <h5 className="card-title">pID: {node.ProcessID} - Port: {node.Port}</h5>
                <div className="card-text">
                    {error}
                    {
                        db && db.items && db.items.map((item: Item, index: number) => {
                            return <div key={index} style={{ backgroundColor: colors[parseInt(item.TenantID) - 1], borderBottom: "1px solid black", padding: "2px 5px 2px 10px" }}>{item.TenantID} - {item.ItemID}</div>
                        })
                    }
                </div>
            </div>
        </div>
    );
}

export default NodeCard;