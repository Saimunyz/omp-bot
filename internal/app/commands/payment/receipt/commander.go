package receipt

import (
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/payment/receipt"
)

type PaymentReceiptCommander struct {
	bot            *tgbotapi.BotAPI
	receiptService *receipt.Service
}

func NewPaymentReceiptCommander(bot *tgbotapi.BotAPI) *PaymentReceiptCommander {
	receiptService := receipt.NewService()

	return &PaymentReceiptCommander{
		bot:            bot,
		receiptService: receiptService,
	}
}

func (c *PaymentReceiptCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("PaymentReceiptCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *PaymentReceiptCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	default:
		c.Default(msg)
	}
}
