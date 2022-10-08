package commands

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-demo/internal/services/product"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

type CommandData struct {
	Task     string
	Parametr int
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

// main hendler
func (c *Commander) HandleUpdate(update tgbotapi.Update) { //hendler
	defer func() { //if panic
		if panicValue := recover(); panicValue != nil { // short format for "if"
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()

	// if button (collback) then operate "get_1"
	if update.CallbackQuery != nil {
		//other way is: args := strings.Split(update.CallbackQuery.Data, "_") //pars in "Data" of button
		parsedData := CommandData{}
		json.Unmarshal([]byte(update.CallbackQuery.Data), &parsedData)
		switch parsedData.Task {
		case "Pagenum":
			c.Pagenum(update.CallbackQuery.Message, parsedData.Parametr)
			log.Printf("Next page")
		default:
			log.Printf("wrong Task in Button")
		}
		return
	}

	if update.Message == nil { // If we got a non-message
		return
	}

	command, ok := registeredCommands[update.Message.Command()] // if command
	if ok {
		command(c, update.Message) // start command from list
	} else {
		c.Default(update.Message)
	}

}
