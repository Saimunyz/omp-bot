package receipt

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("wrong args: %v", args)
		c.DisplayError(inputMessage, wrongIndex)
		return
	}

	receipt, err := c.receiptService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get receipt with idx %d: %v", idx, err)
		c.DisplayError(inputMessage, wrongIndex)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		receipt.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.Get: error sending reply message to chat - %v", err)
		c.DisplayError(inputMessage, defaultErr)
		return
	}
}
