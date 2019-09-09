import React from 'react'
import { shutdownNode } from '../api';

interface Props {
    upNodes: (nodeCount: number) => void;
    reRender: any;
}

const NodeForm = (props: Props) => {
    const [error, setError] = React.useState<string>("");
    const [nodeCountOrProcessID, setNodeCountOrProcessID] = React.useState<string>("0");
    const [operation, setOperation] = React.useState<string>("0");

    const handleInput = (e: any) => {
        setNodeCountOrProcessID(e.target.value);
    }

    const shutdown = () => {
        setError("");
        shutdownNode(parseInt(nodeCountOrProcessID)).then(data => {
            if (data.error) {
                setError(data.error.message)
            } else {
                props.reRender();
            }
        })
    }

    const alert = () => {
        if (error) {
            return (
                <div className="mt-3 alert alert-danger" role="alert">
                    {error}
                </div>
            );
        }
    }

    const handleSelect = (e: any) => {
        setNodeCountOrProcessID("");
        setOperation(e.target.value);
    }

    const showInput = () => {
        switch (operation) {
            case "1":
                return (
                    <React.Fragment>
                        <label htmlFor="tenant">Node count (Only even numbers)</label>
                        <input type="text" className="form-control" id="tenant" placeholder="Node count" onChange={handleInput} />
                    </React.Fragment>
                );
            case "2": return (
                <React.Fragment>
                    <label htmlFor="tenant">ProcessID</label>
                    <input value={nodeCountOrProcessID} type="text" className="form-control" id="tenant" placeholder="ProcessID" onChange={handleInput} />
                </React.Fragment>
            );
            case "3": return;
        }
    }

    const showButton = () => {
        switch (operation) {
            case "1":
                return (
                    <React.Fragment>
                        <button onClick={(e: any) => {
                            e.preventDefault();
                            setError("");
                            if (parseInt(nodeCountOrProcessID) > 10) {
                                setError("Node count can not be greather than 10");
                            } else {
                                props.upNodes(parseInt(nodeCountOrProcessID) % 2 ? (parseInt(nodeCountOrProcessID) + 1) : parseInt(nodeCountOrProcessID));
                            }

                        }} type="submit" className="btn btn-primary">Up</button>
                    </React.Fragment>
                );
            case "2": return (
                <React.Fragment>
                    <button onClick={(e: any) => {
                        e.preventDefault();
                        shutdown()
                    }} type="submit" className=" btn btn-danger">Shutdown</button>
                </React.Fragment>
            );
            case "3": return (
                <button onClick={(e: any) => {
                    e.preventDefault();
                }} type="submit" className="btn btn-success">Add and replicate</button>
            );
        }
    }

    return (
        <div>
            {alert()}
            <form className="mt-3">
                <div className="form-group">
                    <label htmlFor="operation">Operation</label>
                    <select onChange={handleSelect} className="form-control mb-4" id="operation">
                        <option>Select an operation</option>
                        <option value="1">Up nodes(only at the begining, this operation shutdowns other nodes)</option>
                        <option value="2">Shutdown a node by processId </option>
                        <option value="3">Add a new node and replicate - TODO</option>
                    </select>

                    {showInput()}
                </div>

                {showButton()}
            </form>
        </div >
    )
}

export default NodeForm;