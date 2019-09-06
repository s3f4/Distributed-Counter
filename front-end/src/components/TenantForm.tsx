import React from 'react'

const TenantForm = () => {
    return (
        <form className="mt-3">
            <div className="form-group">
                <label htmlFor="tenant">Tenant</label>
                <input type="email" className="form-control" id="tenant" aria-describedby="emailHelp" placeholder="Enter tenant" />
            </div>
            <div className="form-group">
                <label htmlFor="item">Item</label>
                <input type="item" className="form-control" id="item" placeholder="Item Id" />
            </div>
            <button type="submit" className="btn btn-primary">Send</button>
        </form>
    )
}

export default TenantForm;