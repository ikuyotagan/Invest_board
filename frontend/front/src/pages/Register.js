import React from "react";
import { useState } from "react";
import { Redirect } from "react-router-dom";
import { Form, Button } from "react-bootstrap";

const Register = (props) => {
  const [email, setEmail] = useState();
  const [password, setPassword] = useState();
  const [redirect, setRedirect] = useState(false);
  const [error, setError] = useState("");

  const submit = async (e) => {
    e.preventDefault();

    const response = await fetch(props.api + "/users", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        email,
        password,
      }),
    });

    const data = await response.json();

    if (response.ok) {
      setRedirect(true);
    } else {
      setError(data.error);
    }
  };

  if (redirect) {
    return <Redirect to="/login" />;
  }

  let errorResponse = (
    <h1
      style={{ color: "red", fontSize: "15px" }}
      className="h3 mb-3 fw-normal"
    >
      {error}
    </h1>
  );

  return (
    <form onSubmit={submit}>
      <h1 className="h3 mb-3 fw-normal">Please Register</h1>
      {errorResponse}
      <Form.Group className="mb-3">
        <Form.Control
          type="email"
          className="form-control"
          placeholder="name@example.com"
          required
          onChange={(e) => setEmail(e.target.value)}
        />
        <Form.Control
          type="password"
          className="form-control"
          placeholder="Password"
          required
          onChange={(e) => setPassword(e.target.value)}
        />
      </Form.Group>
      <Button className="w-100 btn btn-lg btn-primary" type="submit">
        Register
      </Button>
    </form>
  );
};

export default Register;
