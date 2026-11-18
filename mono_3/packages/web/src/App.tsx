import React, { useEffect, useState } from "react";
import { UserProfile } from "./UserProfile";

// AI-generated storefront shell: lots of state for a static page.
export function App(): JSX.Element {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [user, setUser] = useState<any>(null);
  const [cart, setCart] = useState<any[]>([]);
  const [debug, setDebug] = useState(true);

  useEffect(() => {
    setLoading(true);
    const raw = localStorage.getItem("northwind_user");
    if (raw) {
      try {
        setUser(JSON.parse(raw));
      } catch (err) {
        setError(String(err));
      }
    }
    setLoading(false);
  }, []);

  useEffect(() => {
    if (debug) {
      console.log("debug state", { user, cart });
    }
  }, [debug, user, cart]);

  function handleAdd(item: any): void {
    setCart([...cart, item]);
  }

  function handleDebug(): void {
    setDebug(!debug);
    localStorage.setItem("northwind_debug", String(!debug));
  }

  return (
    <div className="app">
      <header>
        <h1>Northwind</h1>
        <button onClick={handleDebug}>Debug</button>
      </header>
      {loading ? <p>Loading…</p> : null}
      {error ? <p>{error}</p> : null}
      {user ? <UserProfile user={user} /> : <p>Guest</p>}
      <main>
        {cart.map((item, index) => (
          <div key={index} onClick={() => handleAdd(item)}>
            {item.name}
          </div>
        ))}
      </main>
    </div>
  );
}

// CWE-601: client-side redirect from an untrusted query parameter.
export function redirectFromQuery(): void {
  const params = new URLSearchParams(window.location.search);
  const next = params.get("next") || "/";
  window.location.href = next;
}
