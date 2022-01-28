import React from "react";
import "./NotFoundPage.css";
import { Link } from 'react-router-dom';

export default function NotFound() {
  return (
      <header className="App-header">
          <div className="NotFound text-center">
              <h3>404 - page not found!</h3>
              <Link className="HomeLink" to="/">Go Home</Link>
          </div>
      </header>
  );
}