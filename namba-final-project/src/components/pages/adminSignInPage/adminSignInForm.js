import React, {useState} from "react";
import {APIBase} from "../../../config";
import {useAppContext} from "../../../lib/contextLib";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "./adminSignInForm.css"
import httpClient from "../../../httpClient"
import { useNavigate } from "react-router-dom";


export const AdminSignInForm = () => {
    const { userHasAuthenticated } = useAppContext();
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();

    function validateForm() {
    return username.length > 0 && password.length > 0;
  }

    function handleSubmit(event) {
        event.preventDefault();
        //alert("submitted")

                if (username === "gitit" && password === "1234"){
                    let path = "/main";
                    navigate(path);
                    userHasAuthenticated(true);
                }else{
                    alert("User and/or password is not correct");
                }


        /*httpClient.get(APIBase + "/cookie")
            .then(res => {
                console.log(res);
            });
        httpClient.get(APIBase + "/cookie")
            .then(res => {
                console.log(res);
            });*/
        /*fetch(APIBase + "/auth_admin/" + username + "/" + password, {method: 'POST', mode: "cors"})
            .then(res => res.json())
            .then(data => {
                //console.log(data);
                if (data.result === true){
                    //console.log("auth is true");
                    userHasAuthenticated(true);
                    fetch(APIBase + "/cookie", {method: 'GET', mode: "cors"})
                        .then(res => res.json())
                        .then(data => {
                            alert("status????")
                        });
                    fetch(APIBase + "/cookie", {method: 'GET', mode: "cors"})
                        .then(res => res.json())
                        .then(data => {
                            alert("status????")
                        });
                }
                else{
                    alert("User and/or password is not correct");
                }
            });*/
    }

    return(
        <div className = "SignInForm">
            <Form onSubmit={handleSubmit}>
                <Form.Group size="lg" controlId="username">
                <Form.Label>Name</Form.Label>
                    <Form.Control type="username"  placeholder="Enter username" value={username}
                           onChange={(e) => setUsername(e.target.value)}/>
                    </Form.Group>
                <Form.Group size="lg" controlId="password">
                <Form.Label>Password:</Form.Label>
                    <Form.Control type="password"  placeholder="Enter password" value={password}
                           onChange={(e) => setPassword(e.target.value)}/>

                     </Form.Group>
                <Button className="custom-btn" type="submit" block size="lg" disabled={!validateForm()} > SignIn</Button>
            </Form>
        </div>
    );
}


/*export class AdminSignInForm extends React.Component{
    constructor(props) {
    super(props);
    this.username = {username: ''};
    this.password = {password: ''};

    this.handleChangeUser = this.handleChangeUser.bind(this);
    this.handleChangePassword = this.handleChangePassword.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleChangeUser(event) {
        this.setState({username: event.target.value});
  }
  handleChangePassword(event) {
        this.setState({password: event.target.value});
  }

  handleSubmit(event) {
        event.preventDefault();
        *//*fetch(APIBase + "/add_admin/" + this.state.username + "/" + this.state.password, {method: 'POST', mode: "cors"})
            .then(res => res.json())
            .then(data => {
                console.log(data);
            });*/
        /*fetch(APIBase + "/auth_admin/" + this.state.username + "/" + this.state.password, {method: 'POST', mode: "cors"})
            .then(res => res.json())
            .then(data => {
                console.log(data);
                if (data.result === true){
                    console.log("auth is true");
                }
            });
  }

  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <label>
          Name:
            <input name="username" type="username" className="form-control" placeholder="Enter username"
                   onChange={this.handleChangeUser}/>
        </label>
          <br></br>
          <label>
          Password:
            <input name="password" type="password" className="form-control" placeholder="Enter password"
                   onChange={this.handleChangePassword}/>
        </label>
          <br></br>
        <input type="submit" value="Submit" />
      </form>
    );
  }
}*/