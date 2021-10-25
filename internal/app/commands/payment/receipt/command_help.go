package receipt

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RCommander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help__payment__receipt - help\n"+
			"/list__payment__receipt - list products\n"+
			"/get__payment__receipt - get a receipt\n"+
			"/delete__payment__receipt - delete an existing receipt\n"+
			"/new__payment__receipt - create a new receipt\n"+
			"/edit__payment__receipt - edit a receipt",
	)

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.Help: error sending reply message to chat - %v", err)
		c.DisplayError(inputMessage, defaultErr)
		return
	}
}
