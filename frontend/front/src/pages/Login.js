import React, { useState } from "react";
import { Redirect } from "react-router-dom";
import { Form, Button } from "react-bootstrap";

const Login = (props) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);
  const [error, setError] = useState(" ");

  const submit = async (e) => {
    e.preventDefault();

    const response = await fetch(props.api + "/sessions", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        email,
        password,
      }),
    });

    if (response.ok) {
      setRedirect(true);

      const user = await fetch(props.api + "/private/whoami", {
        credentials: "include",
      });

      if (user.ok) {
        const userContent = await user.json();
        props.setName(userContent.email);

        const isTinkoffKeyExist = await fetch(props.api + "/private/tinkoff/proverka", {
          credentials: "include",
        });

        if (isTinkoffKeyExist.ok) {
          props.setTKey(true);
        }
      }
    } else {
      const data = await response.json();
      setError(data.error);
    }
  };

  let errorResponse = "";

  if (error !== "") {
    errorResponse = (
      <h1
        style={{ color: "red", fontSize: "15px" }}
        className="h3 mb-3 fw-normal"
      >
        {error}
      </h1>
    );
  }

  if (redirect) {
    return <Redirect to="/" />;
  }

  return (
    <form onSubmit={submit}>
      <h1 className="h3 mb-3 fw-normal">Please Sign in</h1>
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
        Sign in
      </Button>
    </form>
  );
};

export default Login;
