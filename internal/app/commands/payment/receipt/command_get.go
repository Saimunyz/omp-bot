package receipt

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)

		indexes := c.receiptService.AvailIndex()
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("You have to write an index from %v", indexes),
		)

		_, _ = c.bot.Send(msg)
		return
	}

	receipt, err := c.receiptService.Describe(uint64(idx))
	if err != nil {
		log.Printf("fail to get receipt with idx %d: %v", idx, err)

		indexes := c.receiptService.AvailIndex()
		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			fmt.Sprintf("You have to write an index from %v", indexes),
		)
		_, _ = c.bot.Send(msg)

		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		receipt.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.Get: error sending reply message to chat - %v", err)
	}
}
