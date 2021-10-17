package receipt

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PaymentReceiptCommander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}

	receipt, err := c.receiptService.Get(idx)
	if err != nil {
		log.Printf("fail to get receipt with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		receipt.String(),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PaymentReceiptCommander.Get: error sending reply message to chat - %v", err)
	}
}
