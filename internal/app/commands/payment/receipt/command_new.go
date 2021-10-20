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
		indexes := c.receiptService.AvailIndex()
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("You have to write an index not from %v\n"+
				" and data in json format:\n"+
				"{\"ID\": <new id>,\n"+
				"\"Descr\": \"<some description text>\"\n"+
				"\"Goods\": {\"<tool>\": <price>}}\n"+
				"ALL ARGUMENTS ARE NOT REQUIRED AND"+
				"ID SETS ON MAX IF IT UNSET", indexes),
		)
		_, _ = c.bot.Send(msg)
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
		log.Println("Cannot create new receipt", args)

		indexes := c.receiptService.AvailIndex()
		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("You have to write an index not from %v\n"+
				" and data in json format:\n"+
				"{\"ID\": <new id>,\n"+
				"\"Descr\": \"<some description text>}\"\n"+
				"\"Goods\": {\"<tool>\": <price>}}\n"+
				"ALL ARGUMENTS ARE NOT REQUIRED", indexes),
		)
		_, _ = c.bot.Send(msg)
		return
	}

	newReceipt, _ := c.receiptService.Describe(idx)

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		"Successful createt new receipt\n"+
			newReceipt.String(),
	)
	_, _ = c.bot.Send(msg)
}
