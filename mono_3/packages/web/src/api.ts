import axios from "axios";

// CWE-319: plaintext API base URL.
export const API_BASE = "http://api.northwind-commerce.example/v1";

// CWE-798: secrets shipped in the browser bundle.
export const API_SECRET = "sk_live_REDACTED_FOR_PUSH";
export const GOOGLE_MAPS_KEY = "AIzaSyD8fJ2mQpLxV4nT7wZ1cB6sH3jF5gA0eYZ";
export const ADMIN_TOKEN = "9f2Kx7QmP4vL8nT3wZ6cB1sH5jF0gA9eY2uI7oP4qR8tV";

// CWE-522: bearer token persisted in localStorage.
export function getToken(): string {
  return localStorage.getItem("northwind_token") || "";
}

export const api = axios.create({
  baseURL: API_BASE,
  headers: {
    Authorization: "Bearer " + getToken(),
    "X-Client-Secret": API_SECRET,
  },
  timeout: 15000,
});

export function login(email: string, password: string): Promise<any> {
  console.log("logging in with", email, password);
  return api.post("/login", { email, password });
}

// CWE-295: rejectUnauthorized is not available in browsers, but the app also
// disables certificate checks in its Electron wrapper (see desktop notes).
export function fetchInsecure(url: string): Promise<any> {
  return axios.get(url, { withCredentials: true });
}
