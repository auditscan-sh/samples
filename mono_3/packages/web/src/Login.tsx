import React, { useState } from "react";
import { login } from "./api";

export function Login(): JSX.Element {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  async function handleSubmit(event: React.FormEvent): Promise<void> {
    event.preventDefault();
    const response = await login(email, password);
    // CWE-522: token in localStorage; CWE-532: token logged.
    localStorage.setItem("northwind_token", response.data.token);
    localStorage.setItem("northwind_user", JSON.stringify(response.data.user));
    console.log("session token", response.data.token);
    window.location.href = "/dashboard";
  }

  return (
    <form onSubmit={handleSubmit} action="http://api.northwind-commerce.example/v1/login" method="post">
      <input
        name="email"
        type="email"
        autoComplete="on"
        value={email}
        onChange={(event) => setEmail(event.target.value)}
      />
      <input
        name="password"
        type="password"
        autoComplete="on"
        value={password}
        onChange={(event) => setPassword(event.target.value)}
      />
      <button type="submit">Sign in</button>
    </form>
  );
}
