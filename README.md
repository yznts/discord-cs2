
<p align="center">
    <img width="200" style="border-radius:50%" src=".github/resources/logo.jpg" />
</p>

<h1 align="center">CS2 Manager Bot</h1>

<p align="center">
	Manage your Counter-Strike 2 game server with Discord
</p>

```bash
cs2bot \
	-token      <DISCORD_TOKEN> \
	-rcon-addr  <RCON_ADDR> \
	-rcon-pass  <RCON_PASS>
```

A great way to bring some fun to your Discord server.
Pairs nice with [this](https://github.com/joedwards32/CS2) CS2 server docker image.

Provides raw RCON interface, as well as some useful, user-friendly shortcuts.

## Installation & Configuration

There are multiple ways to configure and run the bot. Just pick options that suit you best.

You can either download the latest binary release from the [releases](https://github.com/yuriizinets/discord-cs2/releases) page, or use the docker image.

To configure the bot, you can provide command line arguments or environment variables.

| Argument     | Environment Variable | Description |
|--------------|----------------------|-------------|
| `-token`     | `TOKEN`              | Required. Discord bot token |
| `-rcon-addr` | `RCON_ADDR`          | Required. CS2 server RCON address |
| `-rcon-pass` | `RCON_PASS`          | Required. CS2 server RCON password |
| `-serv-addr` | `SERV_ADDR`          | Optional. CS2 server address. Used for description. |
| `-serv-pass` | `SERV_PASS`          | Optional. CS2 server password. Used for description. |

Example using binary:

```bash
cs2bot \
	-token      <DISCORD_TOKEN> \
	-rcon-addr  <RCON_ADDR> \
	-rcon-pass  <RCON_PASS> \
	-serv-addr  <SERV_ADDR> \
	-serv-pass  <SERV_PASS>
```

Example using docker:

```bash

docker run -d  \
	-e TOKEN=<TOKEN> \
	-e RCON_ADDR=<RCON_ADDR> \
	-e RCON_PASS=<RCON_PASS> \
	-e SERV_ADDR=<SERV_ADDR> \
	-e SERV_PASS=<SERV_PASS> \
    --name cs2bot \
    ghcr.io/yuriizinets/discord-cs2:latest
```

Example using docker-compose:

```yaml
version: '3'

services:
    cs2bot:
        image: ghcr.io/yuriizinets/discord-cs2:latest
		container_name: cs2bot
		restart: unless-stopped
		environment:
			- TOKEN=<TOKEN>
			- RCON_ADDR=<RCON_ADDR>
			- RCON_PASS=<RCON_PASS>
			- SERV_ADDR=<SERV_ADDR>
			- SERV_PASS=<SERV_PASS>
```

## Usage

| Command     | Arguments       | Description |
|-------------|-----------------|-------------|
| `/about`    | -               | Show bot description and server info |
| `/rcon`     | `<command:str>` | Send raw RCON command |
| `/mode`     | `<mode:str>`    | Change mode (choices provided) |
| `/map`      | `<map:str>`     | Change map (choices provided) |
| `/restart`  | -               | Restart match |
| `/warm`     | -               | Warmup skip (means "I'm already warm") |
| `/pause`    | -               | Pause match |
| `/unpause`  | -               | Unpause match |

Please note that after changing mode you have to set map again.
