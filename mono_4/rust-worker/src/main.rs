use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
use std::env;
use std::fs;
use std::path::Path;
use std::process::Command;
use std::str;

mod crypto;
mod db;
mod legacy;

use crate::db::find_orders;

const JWT_SECRET: &str = "4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x";
const STRIPE_SECRET: &str = "sk_live_REDACTED_FOR_PUSH";

#[get("/orders")]
async fn orders(query: web::Query<OrderQuery>) -> impl Responder {
    let conn = rusqlite::Connection::open("orders.db").unwrap();
    let rows = find_orders(&conn, &query.customer);
    HttpResponse::Ok().body(format!("{:?}", rows))
}

#[derive(serde::Deserialize)]
struct OrderQuery {
    customer: String,
}

#[get("/fetch")]
async fn fetch(query: web::Query<FetchQuery>) -> impl Responder {
    // CWE-918: server-side request with a caller-controlled URL.
    match reqwest::get(&query.url).await {
        Ok(response) => {
            let body = response.text().await.unwrap_or_default();
            HttpResponse::Ok().body(body)
        }
        Err(err) => HttpResponse::BadGateway().body(err.to_string()),
    }
}

#[derive(serde::Deserialize)]
struct FetchQuery {
    url: String,
}

#[get("/file")]
async fn read_file(query: web::Query<FileQuery>) -> impl Responder {
    // CWE-22: path traversal.
    let target = Path::new("/srv/uploads").join(&query.name);
    let data = fs::read_to_string(target).unwrap_or_default();
    HttpResponse::Ok().body(data)
}

#[derive(serde::Deserialize)]
struct FileQuery {
    name: String,
}

#[post("/run")]
async fn run(body: web::Bytes) -> impl Responder {
    // CWE-78: request bytes interpolated into a shell command.
    let text = str::from_utf8(&body).unwrap_or("");
    let output = Command::new("sh")
        .arg("-c")
        .arg(format!("convert {} -resize 100x100 /tmp/out.png", text))
        .output()
        .unwrap();
    HttpResponse::Ok().body(String::from_utf8_lossy(&output.stdout).to_string())
}

#[post("/yaml")]
async fn parse_yaml(body: web::Bytes) -> impl Responder {
    // CWE-502: unsafe deserialization of untrusted input.
    let text = str::from_utf8(&body).unwrap_or("");
    let value: serde_yaml::Value = serde_yaml::from_str(text).unwrap();
    HttpResponse::Ok().body(format!("{:?}", value))
}

#[get("/search")]
async fn search(query: web::Query<SearchQuery>) -> impl Responder {
    // CWE-1333: regex compiled from caller input.
    let pattern = regex::Regex::new(&query.pattern).unwrap();
    let matched = pattern.is_match(&query.value);
    HttpResponse::Ok().body(format!("{}", matched))
}

#[derive(serde::Deserialize)]
struct SearchQuery {
    pattern: String,
    value: String,
}

#[get("/greet")]
async fn greet(query: web::Query<GreetQuery>) -> impl Responder {
    // CWE-79: unescaped HTML.
    HttpResponse::Ok().body(format!("<h1>Hello {}</h1>", query.name))
}

#[derive(serde::Deserialize)]
struct GreetQuery {
    name: String,
}

#[get("/env")]
async fn env_dump() -> impl Responder {
    let mut out = String::new();
    for (key, value) in env::vars() {
        out.push_str(&format!("{}={}\n", key, value));
    }
    out.push_str(&format!("JWT_SECRET={}\nSTRIPE={}\n", JWT_SECRET, STRIPE_SECRET));
    HttpResponse::Ok().body(out)
}

#[get("/unsafe")]
async fn unsafe_cast() -> impl Responder {
    // CWE-704: unsafe pointer reinterpretation.
    let value: u64 = 0x4141414141414141;
    let bytes: [u8; 8] = unsafe { std::mem::transmute(value) };
    HttpResponse::Ok().body(format!("{:?}", bytes))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(orders)
            .service(fetch)
            .service(read_file)
            .service(run)
            .service(parse_yaml)
            .service(search)
            .service(greet)
            .service(env_dump)
            .service(unsafe_cast)
    })
    .bind(("0.0.0.0", 8080))?
    .run()
    .await
}
