import React from 'react'

interface Props {
    upServers: (serverCount: number) => void;
}

const ServerForm = (props: Props) => {
    const [serverCount, setServerCount] = React.useState<string>("0");

    const handleInput = (e: any) => {
        setServerCount(e.target.value);
    }

    const handleSubmit = (e: any) => {
        e.preventDefault();
        props.upServers(parseInt(serverCount));
    }

    return (
        <div>
            <form className="mt-3" onSubmit={(e: any) => handleSubmit(e)}>
                <div className="form-group">
                    <label htmlFor="tenant">Server Count</label>
                    <input type="tenant" className="form-control" id="tenant" placeholder="Enter Server Count" onChange={handleInput} />
                </div>
                <button type="submit" className="btn btn-primary">Up</button>
            </form>
        </div>
    )
}

export default ServerForm;