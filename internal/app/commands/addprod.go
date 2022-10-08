package commands

import (
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-demo/internal/services/product"
)

type Sku struct {
	Title       string
	Description string
	Price       float64
}

func (c *Commander) Addprod(inputMessage *tgbotapi.Message) {

	args := inputMessage.CommandArguments() // args is after command like /start 111_222_333
	newbie := strings.Split(args, "_")      // conv into int. letters will be 0 (zeroes)
	if len(newbie) != 3 {
		log.Println("wrong argument: ", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong argument.\nExpecting /addprod <product Name_Description_Price>\nFor example:\n/addprod Huggy Waggy_toy_100"))
		return
	}
	price, err := strconv.ParseFloat(newbie[2], 64)
	if err != nil {
		log.Println("Wrong Price format. Expecting numbers like 123.45")
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong Price format. Expecting numbers like 123.45"))
		return
	}
	newsku := product.Sku{
		Title:       newbie[0],
		Description: newbie[1],
		Price:       price,
	}
	idn := c.productService.Addprod(&newsku)

	err = c.productService.RewriteStorage()
	if err != nil {
		log.Printf("fail to rewrite data to storage", err)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "fail to rewrite data to storage"))
		return
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, newsku.Title+" added with ID = "+strconv.Itoa(idn))
	c.bot.Send(msg)
}
func init() { //register command in list
	registeredCommands["addprod"] = (*Commander).Addprod
}
