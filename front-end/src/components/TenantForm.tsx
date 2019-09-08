import React from 'react'
import { insertItem } from '../api';

const TenantForm = (props: any) => {
    const initalState = {
        TenantID: "",
        ItemID: ""
    }

    const [error, setError] = React.useState<any>(undefined);
    const [values, setValues] = React.useState<any>(initalState);

    const handleChange = (name: string) => (event: any) => {
        const value = event.target.value;
        setValues({ ...values, [name]: value });
    }

    const send = (e: any) => {
        e.preventDefault();
        insertItem(values).then(data => {
            if (data.error) {
                setError(data.error)
            } else {

            }
            console.log(data);
        });
        props.reRender();
    }

    return (
        <form onSubmit={send} className="mt-3">
            <div className="form-group">
                <label htmlFor="tenant">Tenant</label>
                <input onChange={handleChange("TenantID")} type="text" name="TenantID" className="form-control" id="tenant" aria-describedby="emailHelp" placeholder="Enter tenant" />
            </div>
            <div className="form-group">
                <label htmlFor="item">Item</label>
                <input onChange={handleChange("ItemID")} type="text" name="ItemID" className="form-control" id="item" placeholder="Item Id" />
            </div>
            <button type="submit" className="btn btn-primary">Send</button>
        </form>
    )
}

export default TenantForm;