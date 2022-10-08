package commands

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kormiltsev/tbot-demo/internal/services/product"
)

func (c *Commander) Editprod(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	newbie := strings.Split(args, "_")
	if len(newbie) != 4 {
		log.Println("wrong argument: ", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong argument.\nExpecting /editprod <id to change>_<New product Name>_<New Description>_<New Price>\nFor example:\n/editprod 2_Huggy Waggy_toy_100"))
		return
	}

	idx, err := strconv.Atoi(newbie[0])
	if err != nil {
		log.Println("wrong argument number: ", args)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong argument number.\nExpecting 2_Huggy Waggy_toy_100"))
		return
	}

	productold, ok := c.productService.Get(idx)
	if ok == false {
		log.Printf("fail to get product with id = %v", newbie[0])
		return
	}
	status := fmt.Sprintf("id: %v\nOld product:\nTitle: %s\nDescription: %s\nPrice: %f\n\n", idx, productold.Title, productold.Description, productold.Price)

	price, err := strconv.ParseFloat(newbie[3], 64)
	if err != nil {
		log.Println("Wrong Price format. Expecting numbers like 123.45")
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "Wrong Price format. Expecting numbers like 123.45"))
		return
	}
	// newsku := product
	// newsku.Title = newbie[1]
	// newsku.Description = newbie[2]
	// newsku.Price = price
	newsku := product.Sku{
		Title:       newbie[1],
		Description: newbie[2],
		Price:       price,
	}
	if c.productService.Editprod(&newsku, idx) {
		status += fmt.Sprintf("New product:\nTitle: %s\nDescription: %s\nPrice: %f\n", newsku.Title, newsku.Description, newsku.Price)

	} else {
		status += fmt.Sprintf("is not updated")
	}

	err = c.productService.RewriteStorage()
	if err != nil {
		log.Printf("fail to rewrite data to storage", err)
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "fail to rewrite data to storage"))
		return
	}
	//msg := tgbotapi.NewMessage(inputMessage.Chat.ID, fmt.Sprintf("successfully parsed argument : %v", arg))
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, status)
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}

}
func init() {
	registeredCommands["editprod"] = (*Commander).Editprod
}
