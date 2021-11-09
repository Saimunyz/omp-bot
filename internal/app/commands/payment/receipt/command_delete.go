package receipt

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *RCommander) Delete(inputMsg *tgbotapi.Message) {
	if c.receiptService.Len() == 0 {
		log.Printf("fail to delete from empty slice")
		c.DisplayError(inputMsg, emptySliceErr)
		return
	}

	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("wrong args: %v", args)
		c.DisplayError(inputMsg, wrongIndex)
		return
	}

	_, err = c.receiptService.Remove(uint64(idx))
	if err != nil {
		log.Printf("fail to delete receipt with idx %d: %v", idx, err)
		c.DisplayError(inputMsg, wrongIndex)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Successful deleted idx %d", idx),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.Delete: error sending reply message to chat - %v", err)
		c.DisplayError(inputMsg, defaultErr)
		return
	}
}
