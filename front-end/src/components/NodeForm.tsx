import React from 'react'

interface Props {
    upNodes: (nodeCount: number) => void;
}

const NodeForm = (props: Props) => {
    const [nodeCount, setNodeCount] = React.useState<string>("0");

    const handleInput = (e: any) => {
        setNodeCount(e.target.value);
    }

    const handleSubmit = (e: any) => {
        e.preventDefault();
        props.upNodes(parseInt(nodeCount));
    }

    return (
        <div>
            <form className="mt-3" onSubmit={(e: any) => handleSubmit(e)}>
                <div className="form-group">
                    <label htmlFor="tenant">Node Count</label>
                    <input type="tenant" className="form-control" id="tenant" placeholder="Enter Server Count" onChange={handleInput} />
                </div>
                <button type="submit" className="btn btn-primary">Up</button>
            </form>
        </div>
    )
}

export default NodeForm;