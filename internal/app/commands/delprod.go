package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Delprod(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	idx, err := strconv.Atoi(args) // args to int. (letter converts to "0")
	if err != nil {
		log.Println("wrong argument number: ", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong argument number.\nExpecting: /delprod <product ID>\n/delprod 2\nWARNING! Deleting will shift product ID. To keep it use /editprod"))
		return
	}

	product, ok := c.productService.Get(idx)
	if !ok {
		log.Printf("fail to get product with idx = %v", idx)
		return
	}

	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfully parsed argument : %v", arg))
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("id: %v\nTitle: %s\nDescription: %s\nPrice: %f\n", idx, product.Title, product.Description, product.Price))
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}

	if c.productService.Delprod(idx) {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("DELETED\nIDs other product was shifted"))
		if _, err := c.bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
	err = c.productService.RewriteStorage()
	if err != nil {
		log.Printf("fail to rewrite data to storage", err)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "fail to rewrite data to storage"))
		return
	}
}
func init() {
	registeredCommands["delprod"] = (*Commander).Delprod
}
