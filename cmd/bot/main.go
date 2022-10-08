package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NautiloosGo/telebot/internal/app/commands"
	"github.com/NautiloosGo/telebot/internal/services/product"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load() //upload from .env
	myToken := os.Getenv("TOKEN")
	if myToken == "" {
		log.Printf("input telegram bot unique token:") //if no .env then input via terminal
		fmt.Fscan(os.Stdin, &myToken)
		//scanner := bufio.NewScanner(os.Stdin)
		//myToken = os.Getenv(scanner.Text())
	} // also can start app via terminal: TOKEN="dfgdfg" go run main.go
	bot, err := tgbotapi.NewBotAPI(myToken)
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	//telegram delete incoming messages after 24 hours
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}
	updates := bot.GetUpdatesChan(u)

	product.GetJsonCatalog() //upload catalog from file

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		if update.CallbackQuery == nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		commander.HandleUpdate(update)
	}
}
