import React from "react";
import { Link } from "react-router-dom";
import { Navbar, Container, Nav } from "react-bootstrap";

const Navigation = (props) => {
  const logout = async () => {
    const response = await fetch(props.api + "/private/logout", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
    });

    if (response.ok) {
      props.setName("");
    }
  };

  let menu;

  if (props.name === "") {
    menu = (
      <Navbar bg="dark" variant="dark">
        <Container>
          <Nav>
            <Nav.Link href="/login">Login</Nav.Link>
            <Nav.Link href="/register">Register</Nav.Link>
          </Nav>
        </Container>
      </Navbar>
    );
  } else if (props.tKey === true) {
    menu = (
      <Navbar bg="dark" variant="dark">
        <Container>
          <Nav>
            <Nav.Link href="/login" onClick={logout}>
              Logout
            </Nav.Link>
            <Nav.Link href="/set-key">Set Tinkoff API Key</Nav.Link>
            <Nav.Link href="/perosnal-graph">Your Graphs</Nav.Link>
          </Nav>
        </Container>
      </Navbar>
    );
  } else {
    menu = (
      <Navbar bg="dark" variant="dark">
        <Container>
          <Nav>
            <Nav.Link href="/login" onClick={logout}>
              Logout
            </Nav.Link>
            <Nav.Link href="/set-key">Set Tinkoff API Key</Nav.Link>
            <Nav.Link href="/graph">Graphs</Nav.Link>
          </Nav>
        </Container>
      </Navbar>
    );
  }

  return (
    <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
      <div className="container-fluid">
        <Link to="/" className="navbar-brand">
          Home
        </Link>
        <div>{menu}</div>
      </div>
    </nav>
  );
};

export default Navigation;
