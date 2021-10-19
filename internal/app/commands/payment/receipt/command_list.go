package receipt

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/model/payment"
)

const receiptsPerPage, currPage = uint64(3), uint64(1)

func keybord(data CallbackListData) []tgbotapi.InlineKeyboardButton {
	keyboard := []tgbotapi.InlineKeyboardButton{}
	var nextPageIndex, prevPageIndex uint64

	receiptsLength := uint64(len(payment.AllEntities))

	if (data.CurrPage - 1) > 0 {

		prevPageIndex = data.CurrPage - 1

		serializedData, _ := json.Marshal(CallbackListData{
			CurrPage:        prevPageIndex,
			ReceiptsPerPage: receiptsPerPage,
		})

		callbackPath := path.CallbackPath{
			Domain:       "payment",
			Subdomain:    "receipt",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData("Prev page", callbackPath.String()))
	}

	if (data.CurrPage+1)*data.ReceiptsPerPage-data.ReceiptsPerPage < receiptsLength {
		nextPageIndex = data.CurrPage + 1

		serializedData, _ := json.Marshal(CallbackListData{
			CurrPage:        nextPageIndex,
			ReceiptsPerPage: receiptsPerPage,
		})

		callbackPath := path.CallbackPath{
			Domain:       "payment",
			Subdomain:    "receipt",
			CallbackName: "list",
			CallbackData: string(serializedData),
		}

		keyboard = append(keyboard, tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()))
	}

	return keyboard
}

func (c *RCommander) List(inputMessage *tgbotapi.Message) {
	receipts, err := c.receiptService.List(currPage, receiptsPerPage)
	if err != nil {
		log.Println("Slice is empty")

		msg := tgbotapi.NewMessage(
			inputMessage.Chat.ID,
			"There is no any receipt yet",
		)

		_, _ = c.bot.Send(msg)
		return
	}
	var outputMsgText string

	for _, p := range receipts {
		outputMsgText += p.String()
		outputMsgText += fmt.Sprintf("\n%20s\n", "----------------------------")
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgText)

	buttons := keybord(CallbackListData{
		CurrPage:        currPage,
		ReceiptsPerPage: receiptsPerPage,
	})

	if len(buttons) != 0 {
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(
				buttons...,
			),
		)
	}

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("RCommander.List: error sending reply message to chat - %v", err)
	}
}
