import React, { useState } from "react";
import { Form, Button } from "react-bootstrap";

const SetTinkoffKey = (props) => {
  const [tinkoffKey, setTinkoffKey] = useState("");
  const [isKey, setIsKey] = useState(props.tKey)

  const submit = async (e) => {
    e.preventDefault();

    const response = await fetch("/api/private/set_tinkoff", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({
        tinkoffapikey: tinkoffKey,
      }),
    });

    if (response.ok) {
      setIsKey(true);
    }
  };

  const change = async (e) => {
    e.preventDefault();

    setIsKey(false);
  }

  let content;

  if (isKey === true) {
    content = (
      <div>
        You already set Tinkoff API Key
        <Button onClick={change} className="w-100 btn btn-lg btn-primary">
          Change Tinkoff API Key
        </Button>
      </div>
    );
  } else {
    content = (
      <form onSubmit={submit}>
        <h1 className="h3 mb-3 fw-normal">Please enter Tinkoff Key</h1>
        <Form.Group className="mb-3" controlId="formBasicEmail">
          <Form.Control
            type="password"
            className="form-control"
            placeholder="Tinkoff API Key"
            required
            onChange={(e) => setTinkoffKey(e.target.value)}
          />
        </Form.Group>
        <Button className="w-100 btn btn-lg btn-primary" type="submit">
          Set Tinkoff API Key
        </Button>
      </form>
    );
  }

  return content;
};

export default SetTinkoffKey;
