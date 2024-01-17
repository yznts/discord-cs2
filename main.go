package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gorcon/rcon"
	errx "go.kyoto.codes/zen/v3/errorsx"
)

// Global configuration, vars, etc
var (
	// TOKEN is a bot access token
	TOKEN string
	// SERV_ADDR is a server address (used for information)
	SERV_ADDR string
	// SERV_PASS is a server password (used for information)
	SERV_PASS string
	// RCON_ADDR is a rcon connection address
	RCON_ADDR string
	// RCON_PASS is a rcon connection password
	RCON_PASS string
	// RCON_DIAL is a rcon dialer
	RCON_DIAL *rcon.Conn

	// DISCORD is a discord session instance
	DISCORD *discordgo.Session
	// COMMANDS are bot commands we need to register
	COMMANDS = []Command{
		AboutCommand,
		RconCommand,
		MapCommand,
		RestartCommand,
		WarmCommand,
		PauseCommand,
		UnpauseCommand,
		PracticeCommand,
		// For Tedo
		PivoCommand,
	}
)

func server(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Call corresponding command handler
	for _, command := range COMMANDS {
		if command.Command.Name == i.ApplicationCommandData().Name {
			command.Handler(s, i)
		}
	}
}

// Initialize bot application
func init() {
	// Parse command line arguments
	flag.StringVar(&TOKEN, "token", "", "Bot Token")
	flag.StringVar(&SERV_ADDR, "serv-addr", "", "Server Address")
	flag.StringVar(&SERV_PASS, "serv-pass", "", "Server Password")
	flag.StringVar(&RCON_ADDR, "rcon-addr", "", "RCON Address")
	flag.StringVar(&RCON_PASS, "rcon-pass", "", "RCON Password")
	flag.Parse()
	// Load env
	if TOKEN == "" {
		TOKEN = os.Getenv("TOKEN")
	}
	if SERV_ADDR == "" {
		SERV_ADDR = os.Getenv("SERV_ADDR")
	}
	if SERV_PASS == "" {
		SERV_PASS = os.Getenv("SERV_PASS")
	}
	if RCON_ADDR == "" {
		RCON_ADDR = os.Getenv("RCON_ADDR")
	}
	if RCON_PASS == "" {
		RCON_PASS = os.Getenv("RCON_PASS")
	}
	// Validate arguments
	if TOKEN == "" {
		panic("token is required")
	}
	if RCON_ADDR == "" {
		panic("rcon-addr is required")
	}
	if RCON_PASS == "" {
		panic("rcon-pass is required")
	}
	// Initialize rcon
	RCON_DIAL = errx.Must(rcon.Dial(RCON_ADDR, RCON_PASS))
	// Initialize discord
	DISCORD = errx.Must(discordgo.New(fmt.Sprintf("Bot %s", TOKEN)))
	// React only to message events
	DISCORD.Identify.Intents = discordgo.IntentsGuildMessages
	// Add handler
	DISCORD.AddHandler(server)
}

func main() {
	// Open a websocket connection to Discord and begin listening
	err := DISCORD.Open()
	if err != nil {
		panic(fmt.Errorf("can't open ws connection: %e", err))
	}
	// Remove existing commands
	for _, cmd := range errx.Must(DISCORD.ApplicationCommands(DISCORD.State.User.ID, "")) {
		DISCORD.ApplicationCommandDelete(DISCORD.State.User.ID, "", cmd.ID)
	}
	// Register commands
	for _, cmd := range COMMANDS {
		log.Println("Registering command", cmd.Command.Name)
		_, err := DISCORD.ApplicationCommandCreate(DISCORD.State.User.ID, "", cmd.Command)
		if err != nil {
			panic(err)
		}
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running, press CTRL-C to exit")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanup
	DISCORD.Close()
	RCON_DIAL.Close()
}
