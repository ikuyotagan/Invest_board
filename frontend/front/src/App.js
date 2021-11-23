import React, { useEffect, useState } from "react";
import "./App.css";
import Login from "./pages/Login";
import Navigation from "./components/Nav";
import { BrowserRouter, Route } from "react-router-dom";
import Home from "./pages/Home";
import Register from "./pages/Register";
import GraphInterface from "./pages/GraphInterface";
import SetTinkoffKey from "./pages/SetTinkoffKey";
import PersonalGraphInterface from "./pages/PersonalGraphInterface";

function App() {
  const [name, setName] = useState("");
  const [tKey, setTKey] = useState(false);

  useEffect(() => {
    (async () => {
      const user = await fetch("http://localhost:8080/private/whoami", {
        credentials: "include",
      });

      if (user.ok) {
        const userContent = await user.json();
        setName(userContent.email);
      }
    })();
  }, []);

  useEffect(() => {
    (async () => {
      const isTinkoffKeyExist = await fetch(
        "http://localhost:8080/private/tinkoff/proverka",
        {
          credentials: "include",
        }
      );

      if (isTinkoffKeyExist.ok) {
        setTKey(true);
      }
    })();
  }, [tKey]);

  return (
    <div className="App">
      <BrowserRouter>
        <Navigation
          name={name}
          setName={setName}
          tKey={tKey}
          setTKey={setTKey}
        />
        <main className="form-signin">
          <Route path="/" exact component={() => <Home name={name} />} />
          <Route
            path="/login"
            component={() => <Login setName={setName} setTKey={setTKey} />}
          />
          <Route path="/register" component={Register} />
          <Route path="/graph" component={() => <GraphInterface />} />
          <Route
            path="/set-key"
            component={() => <SetTinkoffKey tKey={tKey} />}
          />
          <Route
            path="/perosnal-graph"
            component={() => <PersonalGraphInterface />}
          />
        </main>
      </BrowserRouter>
    </div>
  );
}

export default App;
