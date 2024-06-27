use poise::{serenity_prelude as serenity};
use crate::core::{Context, Error};

#[poise::command(slash_command, prefix_command)]
pub async fn help(ctx: Context<'_>, #[description = "Selected user"] user: Option<serenity::User>) -> Result<(), Error> {
    let u = user.as_ref().unwrap_or_else(|| ctx.author());
    ctx.say(format!("Help command {}", u.id)).await?;
    Ok(())
}
