import React from "react";
import "../notFoundPage/NotFoundPage.css";
import { Link } from 'react-router-dom';

export default function UnauthorizedPage() {
  return (
      <header className="App-header">
          <div className="NotFound text-center">
              <h3>401 - unauthorized</h3>
              <Link className="HomeLink" to="/">Back to login page</Link>
          </div>
      </header>
  );
}