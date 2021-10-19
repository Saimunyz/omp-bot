package receipt

import (
	"log"

	"github.com/ozonmp/omp-bot/internal/app/path"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/service/payment/receipt"
)

type ReceiptCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type RCommander struct {
	bot            *tgbotapi.BotAPI
	receiptService *receipt.DummyReceiptService
}

func NewReceiptCommander(bot *tgbotapi.BotAPI) *RCommander {
	receiptService := receipt.NewDummyReceiptService()

	return &RCommander{
		bot:            bot,
		receiptService: receiptService,
	}
}

func (c *RCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("PaymentReceiptCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *RCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "list":
		c.List(msg)
	case "get":
		c.Get(msg)
	case "delete":
		c.Delete(msg)
	case "edit":
		c.Edit(msg)
	case "new":
		c.New(msg)
	default:
		c.Default(msg)
	}
}
