package receipt

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RCommander) Edit(inputMsg *tgbotapi.Message) {

	if c.receiptService.Len() != 0 {
		args := strings.SplitN(inputMsg.CommandArguments(), " ", 2)

		idx, err := strconv.Atoi(args[0])
		if err != nil || len(args) < 2 {
			log.Println("wrong args", args)

			indexes := c.receiptService.AvailIndex()
			msg := tgbotapi.NewMessage(
				inputMsg.Chat.ID,
				fmt.Sprintf("You have to write an index from %v\n"+
					" and data in json format:\n"+
					"{\"ID\": <new id>,\n"+
					"\"Descr\": \"<some description text>\"\n"+
					"\"Goods\": {\"<tool>\": <price>}}\n"+
					"ALL ARGUMENTS ARE NOT REQUIRED", indexes),
			)
			_, _ = c.bot.Send(msg)
			return
		}

		receipt, _ := c.receiptService.Describe(uint64(idx))
		//{"ID":1,"Descr":"First purchase","Goods":{"1-tool":200,"2-tool":150}}
		err = json.Unmarshal([]byte(args[1]), receipt)
		if err != nil {

			indexes := c.receiptService.AvailIndex()
			msg := tgbotapi.NewMessage(
				inputMsg.Chat.ID,
				fmt.Sprintf("You have to write an index from %v\n"+
					" and data in json format:\n"+
					"{\"ID\": <new id>,\n"+
					"\"Descr\": \"<some description text>\"\n"+
					"\"Goods\": {\"<tool>\": <price>}}\n"+
					"ALL ARGUMENTS ARE NOT REQUIRED", indexes),
			)
			_, _ = c.bot.Send(msg)
			return
		}

		err = c.receiptService.Update(uint64(idx), *receipt)
		if err != nil {
			log.Println("Cannot update receipt", args)

			indexes := c.receiptService.AvailIndex()
			msg := tgbotapi.NewMessage(
				inputMsg.Chat.ID,
				fmt.Sprintf("You have to write an index from %v\n"+
					" and data in json format:\n"+
					"{\"ID\": <new id>,\n"+
					"\"Descr\": \"<some description text>\"\n"+
					"\"Goods\": {\"<tool>\": <price>}}\n"+
					"ALL ARGUMENTS ARE NOT REQUIRED", indexes),
			)
			_, _ = c.bot.Send(msg)
			return
		}

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			fmt.Sprintf("Successful updeted idx %d", idx),
		)
		_, _ = c.bot.Send(msg)

	} else {
		log.Printf("fail to edit empty slice")

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			"there is no existed receipt",
		)
		_, _ = c.bot.Send(msg)
	}
}
