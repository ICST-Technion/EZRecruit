import React, {useEffect} from 'react';
import Button from "react-bootstrap/Button";
import { useNavigate } from "react-router-dom";


export const AdminMainPage = () => {

    const navigate = useNavigate();
    const pathname = window.location.pathname;

    useEffect(() => {
        if (pathname === "/"){
            navigate("/main");
        }
    },[] )

    const routeChangeToPollCreation = () =>{
    let path = `/createPoll`;
    navigate(path);
  }

  const routeChangeToAddAdmin = () =>{
    let path = `/addAdmin`;
    navigate(path);
  }

  const routeChangeToPollResults = () =>{
    let path = `/pollsResults`;
    navigate(path);
  }

    const routeChangeToUsersPage = () =>{
        let path = `/usersPage`;
        navigate(path);
    }

    return(
        <header className="App-header">
            <div className={"adminMainPage"}>
                <h1>Admin MAIN page</h1>
                <Button className="custom-btn" onClick={routeChangeToPollCreation} > Create a new job </Button>
                <br></br>
                <Button className="custom-btn" onClick={routeChangeToPollResults} > See open jobs </Button>
                <br></br>
                <Button className="custom-btn" onClick={routeChangeToUsersPage} > See applicants </Button>
            </div>
        </header>
    )
}


