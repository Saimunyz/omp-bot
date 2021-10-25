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

	if c.receiptService.Len() == 0 {
		log.Printf("fail to edit empty slice")
		c.DisplayError(inputMsg, emptySliceErr)
		return
	}

	args := strings.SplitN(inputMsg.CommandArguments(), " ", 2)

	idx, err := strconv.Atoi(args[0])
	if err != nil || len(args) < 2 {
		log.Printf("wrong args: %v", args)
		c.DisplayError(inputMsg, jsonEditErr)
		return
	}

	receipt, _ := c.receiptService.Describe(uint64(idx))
	//{"ID":1,"Descr":"First purchase","Goods":{"1-tool":200,"2-tool":150}}
	err = json.Unmarshal([]byte(args[1]), receipt)
	if err != nil {
		log.Printf("RCommander.Edit: %v", err)
		c.DisplayError(inputMsg, jsonEditErr)
		return
	}

	err = c.receiptService.Update(uint64(idx), *receipt)
	if err != nil {
		log.Printf("Cannot update receipt: %v", args)
		c.DisplayError(inputMsg, jsonEditErr)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		fmt.Sprintf("Successful updeted idx %d", idx),
	)
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.Edit: error sending reply message to chat - %v", err)
		c.DisplayError(inputMsg, defaultErr)
		return
	}
}
