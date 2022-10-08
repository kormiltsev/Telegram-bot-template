package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var pagin = 4

var pg int

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()
	if args == "" {
		pg = 1
	} else {

		pgn, err := strconv.Atoi(args)
		if err != nil {
			log.Println("wrong argument got by button. args = ", args)
			c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "wrong argument"))
			return
		}
		pg = pgn
	}

	outputMsgText := "List of products\n(Page #" + strconv.Itoa(pg) + ")\n"
	products := c.productService.List()
	if (pg-1)*pagin < 0 {
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "cant find next page"))
		return
	}
	i := (pg - 1) * pagin
	for i < pagin*pg && i < len(products) {
		outputMsgText += fmt.Sprintf("ID: %d, Title: %s\n", i, products[i].Title)
		i++
	}

	// _, p := range products {
	// 	for i <= pagin {
	// 		outputMsgText += p.Title
	// 		outputMsgText += "\n"
	// 		i++
	// 	}

	// }

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if pg*pagin < len(products) {
		//btnapp := "/list " + strconv.Itoa(pg+1)
		serializedData, _ := json.Marshal(CommandData{
			Task:     "Pagenum",
			Parametr: pg + 1,
		})

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup( // add button
			// tgbotapi.NewInlineKeyboardRow(
			// 	tgbotapi.NewInlineKeyboardButtonURL("yandex", "ya.ru"),
			// ),
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next Page", string(serializedData)),
			),
		)
	}

	//c.bot.Send(msg) correct if no need to check error
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func (c *Commander) Pagenum(inputMessage *tgbotapi.Message, pg int) {

	outputMsgText := "List of products\n(Page #" + strconv.Itoa(pg) + ")\n"
	products := c.productService.List()
	if (pg-1)*pagin < 0 {
		c.bot.Send(tgbotapi.NewMessage(inputMessage.Chat.ID, "cant find next page"))
		return
	}
	i := (pg - 1) * pagin
	for i < pagin*pg && i < len(products) {
		outputMsgText += fmt.Sprintf("ID: %d, Title: %s\n", i, products[i].Title)
		i++
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	if pg*pagin < len(products) {
		serializedData, _ := json.Marshal(CommandData{
			Task:     "Pagenum",
			Parametr: pg + 1,
		})

		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData("Next Page", string(serializedData)),
			),
		)
	}
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
