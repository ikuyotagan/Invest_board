import React, { useEffect, useState } from "react";
import "./App.css";
import Login from "./pages/Login";
import Nav from "./components/Nav";
import { BrowserRouter, Route } from "react-router-dom";
import Home from "./pages/Home";
import Register from "./pages/Register";
import Graph from "./pages/Graph";

function App() {
  const [name, setName] = useState("");

  useEffect(() => {
    (async () => {
      const response = await fetch("http://localhost:8080/private/whoami", {
        credentials: "include",
      });

      if (response.ok) {
        const content = await response.json();
        setName(content.email);
      }
    })();
  });

  return (
    <div className="App">
      <BrowserRouter>
        <Nav name={name} setName={setName} />
        <main className="form-signin">
          <Route path="/" exact component={() => <Home name={name} />} />
          <Route path="/login" component={() => <Login setName={setName} />} />
          <Route path="/register" component={Register} />
          <Route path="/graph" component={() => <Graph />} />
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
