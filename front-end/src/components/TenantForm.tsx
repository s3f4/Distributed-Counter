import React from 'react'
import { insertItem } from '../api';

const TenantForm = (props: any) => {
    const [error, setError] = React.useState<any>("");
    const [TenantID, setTenantID] = React.useState<string>("1");

    const send = (e: any) => {
        e.preventDefault();
        insertItem({
            TenantID,
            ItemID: (new Date().getUTCMilliseconds().toString() + new Date().getTime().toString()).toString()
        }).then(data => {
            setError(data.error ? data.error.message : "")
        });
        props.reRender();
    }

    const range = (start: number, end: number): number[] => {
        if (start === end) return [start];
        return [start, ...range(start + 1, end)];
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

    return (
        <form onSubmit={send} className="mt-3">
            {alert()}
            <div className="form-group">
                <label htmlFor="tenant">Tenant</label>
                <select onChange={(e: any) => setTenantID(e.target.value)} name="TenantID" className="form-control" id="tenant">
                    {range(1, 10).map((id: number) => {
                        return <option value={id}>{id}</option>
                    })}
                </select>
            </div>
            <button type="submit" className="btn btn-primary">Send</button>
        </form>
    )
}

export default TenantForm;