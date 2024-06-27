use log::info;
use poise::serenity_prelude as serenity;
use serenity::all::GatewayIntents;
use std::env;

use crate::bot::commands::general::help;
use crate::core::Data;

mod commands;

pub async fn start() -> Result<(), String> {
    let guild_id_str = env::var("DISCORD_GUILD_ID").expect("Error get guild id");
    let guild_id_u64 = guild_id_str.parse::<u64>().expect("Error parsing guild id");
    let guild_id = serenity::GuildId::from(guild_id_u64);

    let framework = poise::Framework::builder()
        .options(poise::FrameworkOptions {
            commands: vec![help()],
            prefix_options: poise::PrefixFrameworkOptions {
                prefix: Some("!".into()),
                ..Default::default()
            },
            ..Default::default()
        })
        .setup(move |ctx, _ready, framework| {
            Box::pin(async move {
                poise::builtins::register_in_guild(ctx, &framework.options().commands, guild_id).await?;
                Ok(Data {})
            })
        })
        .build();

    let token = env::var("DISCORD_TOKEN").expect("Expected token in env");

    let intents = GatewayIntents::GUILD_MESSAGES
        | GatewayIntents::DIRECT_MESSAGES
        | GatewayIntents::MESSAGE_CONTENT;

    let mut client = serenity::ClientBuilder::new(&token, intents)
        .framework(framework)
        .await
        .expect("Err while creating client");

    info!("starting bot");

    if let Err(why) = client.start().await {
        return Err(format!("Client error {why:?}"));
    }

    Ok(())
}
