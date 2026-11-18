import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// Intentionally permissive dev/preview server.
export default defineConfig({
  plugins: [react()],
  server: {
    host: "0.0.0.0",
    cors: true,
    strictPort: false,
  },
  preview: {
    host: "0.0.0.0",
    cors: true,
  },
});
