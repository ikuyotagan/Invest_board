import React from "react";
import { useState } from "react";
import { Redirect } from "react-router-dom";

const Register = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);
  const [error, setError] = useState(" ");


  const submit = async (e) => {
    e.preventDefault();

    const response = await fetch("http://localhost:8080/users", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        email,
        password,
      }),
    });

    const data = await response.json();

    if (response.ok) {
      setRedirect(true);
    } else{
      setError(data.error)
    }
  };

  if (redirect) {
    return <Redirect to="/login" />;
  }

  let errorResponse = ""

  if (error !== ""){
    errorResponse = (<h1 style={{color: "red", fontSize: "15px",}} className="h3 mb-3 fw-normal">{error}</h1>)
  }

  return (
    <form onSubmit={submit}>
      <h1 className="h3 mb-3 fw-normal">Please register</h1>
      {errorResponse}
      <input
        type="email"
        className="form-control"
        placeholder="name@example.com"
        required
        onChange={(e) => setEmail(e.target.value)}
      />
      <input
        type="password"
        className="form-control"
        placeholder="Password"
        required
        onChange={(e) => setPassword(e.target.value)}
      />
      <button className="w-100 btn btn-lg btn-primary" type="submit">
        Submit
      </button>
    </form>
  );
};

export default Register;
