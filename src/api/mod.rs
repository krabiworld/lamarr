use axum::routing::get;
use axum::{Json, Router};
use std::env;
use log::info;
use serde_json::{json, Value};

pub async fn start() -> std::io::Result<()> {
    let app = Router::new().route("/", get(root));

    let addr = env::var("API_ADDR").expect("Expected addr in env");

    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();

    info!("starting api");

    axum::serve(listener, app).await
}

async fn root() -> Json<Value> {
    Json(json!({"data": "world"}))
}
