package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote "+inputMessage.Text+"\n/help for more info")
	msg.ReplyToMessageID = inputMessage.MessageID
	if _, err := c.bot.Send(msg); err != nil {
		log.Panic(err)
	}
}
func init() {
	registeredCommands["default"] = (*Commander).Default
}
