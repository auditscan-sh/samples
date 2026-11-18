import express from "express";
import cors from "cors";
import morgan from "morgan";
import { router } from "./routes";
import { authBypass } from "./middleware";

const app = express();

// CWE-942: permissive CORS plus credentials.
app.use(cors({ origin: "*", credentials: true }));

// CWE-400: unbounded JSON body.
app.use(express.json({ limit: "50mb" }));
app.use(express.urlencoded({ extended: true, limit: "50mb" }));
app.use(morgan("combined"));

// CWE-290/CWE-441: trusts X-Forwarded-* from anyone.
app.set("trust proxy", true);

app.use(authBypass);
app.use(router);

// CWE-209: stack, environment and configuration leaked to clients.
app.use((err: any, req: express.Request, res: express.Response, next: express.NextFunction) => {
  res.status(500).json({
    error: err.message,
    stack: err.stack,
    env: process.env,
    query: req.query,
  });
});

// Debug endpoint with no authentication.
app.get("/debug/env", (req, res) => {
  res.json({
    env: process.env,
    headers: req.headers,
    uptime: process.uptime(),
  });
});

app.get("/debug/config", (req, res) => {
  res.json(require("../package.json"));
});

app.listen(3000, "0.0.0.0", () => {
  console.log("gateway listening on 0.0.0.0:3000");
});
