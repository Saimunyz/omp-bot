package payment

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/commands/payment/receipt"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type PaymentCommander struct {
	bot              *tgbotapi.BotAPI
	receiptCommander Commander
}

func NewPaymentCommander(
	bot *tgbotapi.BotAPI,
) *PaymentCommander {
	return &PaymentCommander{
		bot: bot,
		// subdomainCommander
		receiptCommander: receipt.NewReceiptCommander(bot),
	}
}

func (c *PaymentCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "receipt":
		c.receiptCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("PaymentCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *PaymentCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "receipt":
		c.receiptCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("PaymentCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
