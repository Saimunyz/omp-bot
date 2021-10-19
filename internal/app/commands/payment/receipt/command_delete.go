package receipt

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RCommander) Delete(inputMsg *tgbotapi.Message) {
	if c.receiptService.Len() != 0 {

		args := inputMsg.CommandArguments()

		idx, err := strconv.Atoi(args)
		if err != nil {
			log.Println("wrong args", args)

			indexes := c.receiptService.AvailIndex()
			msg := tgbotapi.NewMessage(
				inputMsg.Chat.ID,
				fmt.Sprintf("You have to write an index from %v", indexes),
			)

			_, _ = c.bot.Send(msg)
			return
		}

		res, err := c.receiptService.Remove(uint64(idx))

		if err != nil {
			log.Printf("fail to delete receipt with idx %d: %v", idx, err)

			indexes := c.receiptService.AvailIndex()
			msg := tgbotapi.NewMessage(
				inputMsg.Chat.ID,
				fmt.Sprintf("You have to write an index from %v", indexes),
			)
			_, _ = c.bot.Send(msg)

			return
		}

		if res {
			msg := tgbotapi.NewMessage(
				inputMsg.Chat.ID,
				fmt.Sprintf("Successeful deleted idx %d", idx),
			)
			_, _ = c.bot.Send(msg)
		}

	} else {
		log.Printf("fail to delete from empty slice")

		msg := tgbotapi.NewMessage(
			inputMsg.Chat.ID,
			"there is no existed receipt",
		)
		_, _ = c.bot.Send(msg)
	}
}
