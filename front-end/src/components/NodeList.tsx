import React from 'react'
import NodeCard from './NodeCard';


interface Props {
    nodes: Node[];
}

export interface Node {
    ProcessID: number;
    Port: number;
    database: any;
}

const NodeList = (props: Props) => {

    return <div className="d-flex justify-content-center">
        {
            props.nodes && props.nodes.map((node: Node, index: number) => {
                return <NodeCard key={index} node={node} />
            })
        }

    </div>;
}

export default NodeList;