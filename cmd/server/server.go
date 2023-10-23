package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/kurnhyalcantara/Jubelio-Chat-Peer-to-Peer-API/app/config"
	"github.com/kurnhyalcantara/Jubelio-Chat-Peer-to-Peer-API/app/database/postgresql"
)

func Serve() {
	appConfig, err := config.LoadAppConfig()
	if err != nil {
		log.Fatalf("error load app config: %v", err)
	}

	if errDB := postgresql.ConnectDB(); errDB != nil {
		log.Fatalf("error connect db: %v", errDB)
	}

	fmt.Println("Database Connected!")

	app := fiber.New()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	go func() {
		<-sigCh
		fmt.Printf("\nShutting down server...")
		_ = app.Shutdown()
	}()

	serverAddr := fmt.Sprintf("%s:%d", appConfig.HOST, appConfig.PORT)
	if errInitServer := app.Listen(serverAddr); errInitServer != nil {
		log.Fatalf("Oops.. init server error: %v", errInitServer)
	}
}