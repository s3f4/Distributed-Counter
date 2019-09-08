import React from 'react'
import { getNodeDB } from '../api';
import { Node } from "./NodeList";

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
                <p className="card-text">
                    {error}
                    {JSON.stringify(db)}
                </p>
            </div>
        </div>
    );
}

export default NodeCard;