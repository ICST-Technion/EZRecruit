import React from 'react';
import {AdminSignInForm} from "./adminSignInForm";
import logo from "./paamonim.png";

const AdminSignInPage = () => {
    return(
        <header className="App-header">
            <img src={logo} alt="Logo"/>
            <h1>Admin page</h1>
            <AdminSignInForm/>
        </header>
    )
}

export default AdminSignInPage;