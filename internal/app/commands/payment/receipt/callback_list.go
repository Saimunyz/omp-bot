package receipt

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	CurrPage        uint64
	ReceiptsPerPage uint64
}

func (c *RCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("RCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	receipts, _ := c.receiptService.List(parsedData.CurrPage, parsedData.ReceiptsPerPage)
	var outputMsgText string

	for _, p := range receipts {
		outputMsgText += p.String()
		outputMsgText += fmt.Sprintf("\n%20s\n", "----------------------------")
	}

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		outputMsgText,
	)

	buttons := keybord(CallbackListData{
		CurrPage:        parsedData.CurrPage,
		ReceiptsPerPage: receiptsPerPage,
	})

	if len(buttons) != 0 {
		tmp := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				buttons...,
			),
		)
		msg.ReplyMarkup = &tmp
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("RCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
