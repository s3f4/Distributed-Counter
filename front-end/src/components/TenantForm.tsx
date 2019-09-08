import React from 'react'
import { insertItem, getCount } from '../api';

const TenantForm = (props: any) => {
    const [error, setError] = React.useState<any>("");
    const [TenantID, setTenantID] = React.useState<string>("1");
    const [itemCount, setItemCount] = React.useState<number>(0);
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

    const count = (e: any) => {
        e.preventDefault();
        getCount(parseInt(TenantID)).then(data => {
            if (data.error) {
                setError(data.error.message)
            } else {
                setItemCount(data.count)
            }
        });
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
                        return <option key={id} value={id}>{id}</option>
                    })}
                </select>
            </div>
            <button type="submit" className="btn btn-primary">Add an item</button>
            <button type="submit" onClick={count} className="ml-2 btn btn-success">Count</button>
            <span className="ml-2 success">Result : </span> <span className="ml-2 badge badge-success">{itemCount}</span>
        </form>
    )
}

export default TenantForm;