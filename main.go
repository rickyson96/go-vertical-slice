package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/rickyson96/go-vertical-slice/internal/server"
	"github.com/rickyson96/go-vertical-slice/internal/server/config"
	"github.com/rickyson96/go-vertical-slice/internal/server/middlewares"
	"github.com/spf13/viper"
)

var buildTime string = "now"

func init() {
	config.Setup()
	viper.Set("buildTime", buildTime)
}

//go:generate sqlc generate

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	conn, err := server.Conn(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		cancel()
		return
	}

	routes := middlewares.CORSHandler(server.Routes(conn))
	server := server.InitServer(ctx, routes)

	// Accepts graceful shutdowns when quitting via SIGINT (Ctrl + C)
	// SIGKILL, SIGQUIT or SIGTERM will not be caught and will forcefully shuts the application down.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Blocks until we receive graceful shutdown signal
	<-c

	server.Shutdown(ctx)

	cancel()
	<-ctx.Done()
}
