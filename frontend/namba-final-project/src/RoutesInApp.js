import React from "react";
import { Route, Routes } from "react-router-dom";
import AdminSignInPage from "./components/pages/adminSignInPage/adminSignInPage";
import {useAppContext} from "./lib/contextLib";
import {AdminMainPage} from "./components/pages/adminMainPage/adminMainPage";
import CreatePollPage from "./components/pages/createPollPage/createPollPage";
import NotFound from "./components/pages/notFoundPage/notFoundPage";
import AddAdminPage from "./components/pages/addAdminPage/addAdminPage";
import {PollsResultsPage} from "./components/pages/pollsResultsPage/pollsResultsPage";
import {UsersPage} from "./components/pages/usersPage/usersPage";
import UnauthorizedPage from "./components/pages/unauthorizedPage/unauthorizedPage";

export default function RoutesInApp() {
    const {isAuthenticated} = useAppContext();
    const {isAuthenticating} = useAppContext();
    const pathname = window.location.pathname;
  return (<>
          {isAuthenticating ? (<h1> Loading</h1>)
              : (
                  <Routes>
                      {isAuthenticated && (pathname === "/" || pathname === "" ) ? (<Route exact path="/" element={<AdminMainPage/>}/>)
                          : (<Route exact path="/" element={<AdminSignInPage/>}/>)}
                      {isAuthenticated ? (<Route exact path="/main" element={<AdminMainPage/>}/>)
                          : (<Route exact path="/main" element={<UnauthorizedPage/>}/>)}
                      {isAuthenticated ? (<Route exact path="/createPoll" element={<CreatePollPage/>}/>)
                          : (<Route exact path="/createPoll" element={<UnauthorizedPage/>}/>)}
                      {isAuthenticated ? (<Route exact path="/pollsResults" element={<PollsResultsPage/>}/>)
                          : (<Route exact path="/pollsResults" element={<UnauthorizedPage/>}/>)}
                      {isAuthenticated ? (<Route exact path="/usersPage" element={<UsersPage/>}/>)
                          : (<Route exact path="/usersPage" element={<UnauthorizedPage/>}/>)}
                      {isAuthenticated ? (<Route exact path="/addAdmin" element={<AddAdminPage/>}/>)
                          : (<Route exact path="/addAdmin" element={<UnauthorizedPage/>}/>)}
                      <Route path="*" element={<NotFound/>}/>
                  </Routes>)}
      </>
  );
}