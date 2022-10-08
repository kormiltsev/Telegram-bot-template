package commands

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Findprod(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	if args == "" {
		log.Println("no argument", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "no argument\nExpecting:\n/findprod <name>\nExample:\n/findprod apple"))
		return
	}

	products := c.productService.List()

	ans := fmt.Sprintf("Found on '%s':\n\n", args)
	for idx, sku := range products {
		if strings.Contains(sku.Title, args) {
			ans += fmt.Sprintf("\nid: %v, Title: %s", idx, sku.Title)
		}
	}
	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfully parsed argument : %v", arg))
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ans)
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
func init() {
	registeredCommands["findprod"] = (*Commander).Findprod
}
