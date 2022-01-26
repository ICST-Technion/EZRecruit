import React, {useState} from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import {APIBase} from "../../../config";
import httpClient from "../../../httpClient";

const AddAdminPage = () => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    function validateForm(){
        return username.length > 0 && password.length > 0;
    }

    function handleSubmit(event) {
        event.preventDefault();
        httpClient.post(APIBase + "/add_admin/" + username + "/" + password)
            .then(res => {
                if (res["data"].result === true){
                    //console.log("auth is true");
                    alert("Admin added!!")
                }
                else{
                    alert("Can't add admin\n" + res["data"].result);
                }
            }).catch(e => {
            console.log(e);
            alert("An error occurred");
        });

        /*fetch(APIBase + "/add_admin/" + username + "/" + password, {method: 'POST', mode: "cors"})
            .then(res => res.json())
            .then(data => {
                //console.log(data);
                if (data.result === true){
                    //console.log("auth is true");
                    alert("Admin added!!")
                }
                else{
                    alert("Can't add admin\n" + data.result);
                }
            })
            .catch(e => {
                console.log(e);
                alert("An error occurred");
            });*/

    }

return(
    <header className="App-header">
    <div className="AddAdminForm">
        <h1> Add a new admin</h1>
        <Form onSubmit={handleSubmit}>
            <Form.Group size="lg" controlId="username">
                <Form.Label>Name</Form.Label>
                <Form.Control type="username" placeholder="Enter username" value={username}
                              onChange={(e) => setUsername(e.target.value)}/>
            </Form.Group>
            <Form.Group size="lg" controlId="password">
                <Form.Label>Password:</Form.Label>
                <Form.Control type="password" placeholder="Enter password" value={password}
                              onChange={(e) => setPassword(e.target.value)}/>

            </Form.Group>
            <div className={"longShot"}>
                <Button className="custom-btn" type="submit" block size="lg" disabled={!validateForm()}> Add
                    Admin</Button>
            </div>
        </Form>
    </div>
        </header>
    );


}

export default AddAdminPage;