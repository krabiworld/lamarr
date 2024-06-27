mod api;
mod bot;
mod core;

use dotenv::dotenv;
use tokio::signal;
use log::info;

#[tokio::main]
async fn main() {
    dotenv().ok();

    env_logger::init();

    tokio::spawn(async {
        bot::start().await.unwrap();
    });

    tokio::spawn(async {
        api::start().await.unwrap();
    });

    signal::ctrl_c().await.unwrap();
    info!("gracefully shutdown")
}
