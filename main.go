package main

import (
	"context"
	"gitlab.com/systeric/internal/chat/backend/core/internal/server/redis"
	"log"
	"os"
	"os/signal"

	"github.com/spf13/viper"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
	"gitlab.com/systeric/internal/chat/backend/core/internal/server"
	"gitlab.com/systeric/internal/chat/backend/core/internal/server/config"
	"gitlab.com/systeric/internal/chat/backend/core/internal/server/middlewares"
)

var buildTime string = "now"

func init() {
	config.Setup()
	viper.Set("buildTime", buildTime)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	db, err := server.Conn(ctx)
	if err != nil {
		log.Fatal("Unable to connect to database: ", err)
		cancel()
		return
	}

	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatal("Unable to migrate database: ", err)
		cancel()
		return
	}

	redis.Setup()
	routes := middlewares.CORSHandler(server.Routes(db))
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
