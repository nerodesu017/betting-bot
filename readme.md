# Gambling Discord Bot Readme

# [WARNING! I DO NOT CONDONE GAMBLING! BE GAMBLE AWARE!](https://www.begambleaware.org/)

## How to run this bot

1. Create a discord bot ([https://discord.com/developers/applications](https://discord.com/developers/applications))
2. Enable SERVER MEMBERS INTENT and MESSAGE CONTENT INTENT ([https://discord.com/developers/applications/<<APPLICATION_ID>>/bot](https://discord.com/developers/applications/<<APPLICATION_ID>>/bot)) 
3. Add the bot to your server (`https://discord.com/api/oauth2/authorize?client_id=<<BOT_ID>>&permissions=0&scope=bot%20applications.commands`)
4. Rename `.env_copy` to `.env` and fill in the environment variables
5. On the first run, run the app with the `-addCommands` flag (and on any subsequent run that you wish to add/edit existing slash commands)
6. On every other run, run it without the flag as there is no need

(Oh, and have a PostgreSQL database set up)

## Add commands

1. To add commands, head over to the `/src/discord`, which we’ll also denote as DISCORD_DIR
2. Find where `DISCORD_DIR/commands` your command should be
3. Create your own command file (for an example you can head over to `DISCORD_DIR/commands/games/coins/coin_flip.go`)
4. Go to `DISCORD_DIR/commands.go`, import the specific package of your new function, and at the end of the `init` function add your new command as other are added
5. Go to `DISCORD_DIR/commands/help.go` and add your new command to the list

## What else needs to be done?

1. The code needs to be cleaned
2. Better error checking for SQL queries
3. Let the owner of the bot (or any other person you’d like) seed the RNG
4. Edit the interaction of the user with slash commands
5. Add new games
6. Create leaderboards (i.e. biggest wins, highest balances)

## Why?

Earlier this year I got my hands on a pretty awesome book by Mark Bollman, “Basic Gambling Mathematics: The Numbers Behind The Neon” that sparked an interest in how luck-based games work, having had experiences with CS:GO gambling websites in the past and online casinos (which I do NOT encourage at all, since the house always has an advantage). Leading to it, I told myself that I should create a discord bot for these luck games and see where it gets me

## Is the RNG good enough?

Yes, and no. For what it’s worth, for a demo no-real-money for-fun discord gambling bot, it is. If you wish to use it in actual production for a real casino, I’d recommend not to.

It is random enough for the chances to be fair.

### How the RNG was tested

The RNG (basically, math.rand) was tested on running 10,000 simulations, each with 1,000,000 random values for a coin flip. I have chosen this as the coin flip is a valid binomial experiment, thus making use of its distribution and having its distribution and standard deviation be easily calculated. The values haven’t been far from the [68-95-99.7 rule](https://en.wikipedia.org/wiki/68%E2%80%9395%E2%80%9399.7_rule) (results were: 67.32-95.49-99.8)