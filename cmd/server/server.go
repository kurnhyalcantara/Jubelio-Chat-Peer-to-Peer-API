package server

import (
	"fmt"
	"log"

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

	app := fiber.New()

	serverAddr := fmt.Sprintf("%s:%d", appConfig.HOST, appConfig.PORT)
	if errInitServer := app.Listen(serverAddr); errInitServer != nil {
		log.Fatalf("Oops.. init server error: %v", errInitServer)
	}
}