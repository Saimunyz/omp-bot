package receipt

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/payment"
)

func (c *RCommander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	receipt := payment.Receipt{}
	//{"ID":1,"Descr":"First purchase","Goods":{"1-tool":200,"2-tool":150}}
	err := json.Unmarshal([]byte(args), &receipt)
	if err != nil {
		log.Printf("RCommander.New: %v", err)
		c.DisplayError(inputMsg, jsonNewErr)
		return
	}

	indexes := c.receiptService.AvailIndex()
	if receipt.ID == 0 {
		idxLen := len(indexes)
		if idxLen != 0 {
			receipt.ID = indexes[idxLen-1] + 1
		} else {
			receipt.ID = 0
		}
	}

	idx, err := c.receiptService.Create(receipt)
	if err != nil {
		log.Printf("Cannot create new receipt: %v", receipt)
		c.DisplayError(inputMsg, jsonNewErr)
		return
	}

	newReceipt, _ := c.receiptService.Describe(idx)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Successful createt new receipt\n%s",
			newReceipt.String()),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.New: error sending reply message to chat - %v", err)
		c.DisplayError(inputMsg, defaultErr)
		return
	}
}
