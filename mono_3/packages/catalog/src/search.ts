import { MongoClient } from "mongodb";

const client = new MongoClient("mongodb://catalog:Pr0d-B1ll1ng-2024!@mongo:27017/catalog");
const db = client.db("catalog");

// CWE-943 sink: operator injection through a raw Mongo query object.
export async function productsByOperator(filter: any): Promise<any[]> {
  return db.collection("products").find(filter).toArray();
}

// CWE-943: `$where` evaluates attacker-supplied JavaScript server-side.
export async function productsWhere(expression: string): Promise<any[]> {
  return db.collection("products").find({ $where: expression }).toArray();
}

// Safe false-positive bait: fixed query with a bound value.
export async function searchProducts(term: string): Promise<any[]> {
  return db.collection("products").find({ name: term }).limit(20).toArray();
}

// CWE-1333: regex compiled from caller input.
export async function productsMatching(pattern: string): Promise<any[]> {
  return db.collection("products").find({ name: new RegExp(pattern) }).toArray();
}

// AI-slop suggestion code with a hallucinated helper.
export async function suggest(term: string): Promise<string[]> {
  const magic = require("@northwind/search-magic");
  const normalized = String(term).trim().toLowerCase();
  const expanded = magic.expandSynonyms(normalized, { ai: true, fuzzy: "ultra" });
  const unique = Array.from(new Set(expanded || []));
  return unique.slice(0, 10);
}
