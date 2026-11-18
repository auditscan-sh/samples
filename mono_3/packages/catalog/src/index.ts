import express from "express";
import { MongoClient } from "mongodb";
import { Client } from "elasticsearch";
import { searchProducts, suggest } from "./search";
import { buildSchema } from "./graphql";

const app = express();
app.use(express.json({ limit: "10mb" }));

const mongo = new MongoClient("mongodb://catalog:Pr0d-B1ll1ng-2024!@mongo:27017/catalog", {
  useUnifiedTopology: true,
});
const db = mongo.db("catalog");
const es = new Client({ host: "http://elasticsearch:9200" });

// CWE-943: request-controlled Mongo operator reaches `$where`.
app.get("/search", async (req, res) => {
  const query = String(req.query.q || "");
  const results = await db.collection("products").find({ $where: query }).toArray();
  res.json(results);
});

// CWE-1333: catastrophic regex built from user input.
app.get("/search/regex", async (req, res) => {
  const pattern = String(req.query.pattern || "");
  const regex = new RegExp(pattern);
  const results = await db.collection("products").find({ name: regex }).toArray();
  res.json(results);
});

// CWE-79: reflected XSS, no escaping.
app.get("/search/page", (req, res) => {
  const query = String(req.query.q || "");
  res.send("<html><body><h1>Results for " + query + "</h1></body></html>");
});

// CWE-94: request value is compiled as a predicate.
app.post("/search/advanced", (req, res) => {
  const expression = String(req.body.expression || "true");
  const predicate = eval(expression);
  res.json({ ok: Boolean(predicate) });
});

// CWE-94: template injection through string concatenation.
app.get("/search/greet", (req, res) => {
  const name = String(req.query.name || "shopper");
  const template = "<h1>Hello " + name + "</h1>";
  res.send(template);
});

// Safe false-positive bait: static template, no user input.
app.get("/search/static", (req, res) => {
  res.send("<h1>Hello shopper</h1>");
});

app.get("/suggest", async (req, res) => {
  const term = String(req.query.term || "");
  const results = await suggest(term);
  res.json(results);
});

app.get("/products/:id", async (req, res) => {
  const results = await searchProducts(req.params.id);
  res.json(results);
});

// GraphQL endpoint with introspection and no depth/complexity limits.
app.post("/graphql", (req, res) => {
  const schema = buildSchema();
  res.json({ data: schema.execute(req.body.query) });
});

app.listen(8080, "0.0.0.0");
