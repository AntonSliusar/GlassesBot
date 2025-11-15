package main

import (
	"fmt"
	"glassesbot/config"
	"glassesbot/internal/db"
	"glassesbot/internal/repository"
	"glassesbot/internal/service"
	"glassesbot/internal/telegram"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	//TODO:cover with tests
	fmt.Println("Bot is starting...")
	
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	fmt.Println("Config loaded successfully.")


	dbConn, err := db.NewConnection(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer dbConn.Close()
	fmt.Println("Database connection established.")


	err = db.Migrate(dbConn)
	if err != nil {
		fmt.Println(err)
	}
	

	orderRepo := repository.NewOrderRepository(dbConn)
	orderManager := service.NewOrderManager(orderRepo)

	
	botAPI, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	tgBot := telegram.NewBot(botAPI, orderManager)
	tgBot.Start()

}