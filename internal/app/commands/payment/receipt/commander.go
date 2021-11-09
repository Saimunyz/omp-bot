package receipt

import (
	"fmt"
	"log"
	"strings"

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
		log.Printf("RCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

type CustomErr uint8

const (
	help   = "help"
	list   = "list"
	get    = "get"
	delete = "delete"
	edit   = "edit"
	new    = "new"

	defaultErr CustomErr = iota
	jsonNewErr
	jsonEditErr
	emptySliceErr
	wrongIndex
)

func (c *RCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case help:
		c.Help(msg)
	case list:
		c.List(msg)
	case get:
		c.Get(msg)
	case delete:
		c.Delete(msg)
	case edit:
		c.Edit(msg)
	case new:
		c.New(msg)
	default:
		c.Default(msg)
	}
}

func (c *RCommander) DisplayError(inputMsg *tgbotapi.Message, typeErr CustomErr) {

	indexes := c.receiptService.AvailIndex()
	msgErr := strings.Builder{}

	switch typeErr {
	case defaultErr:
		msgErr.WriteString("Error: somethig bad happend, Sorry!\n")
		msgErr.WriteString("Try again or come back later")
	case jsonNewErr:
		msgErr.WriteString("You have to write an index not from")
		msgErr.WriteString(fmt.Sprintf(" %v\n", indexes))
		msgErr.WriteString("and data in json format:\n")
		msgErr.WriteString("{\"ID\": <new id>,\n")
		msgErr.WriteString("\"Descr\": \"<some description text>\"\n")
		msgErr.WriteString("\"Goods\": {\"<tool>\": <price>}}\n")
		msgErr.WriteString("ALL ARGUMENTS ARE NOT REQUIRED AND")
		msgErr.WriteString(" ID IS SET ON MAX IF IT IS NOT SET")
	case jsonEditErr:
		msgErr.WriteString("You have to write an index from")
		msgErr.WriteString(fmt.Sprintf(" %v\n", indexes))
		msgErr.WriteString("and data in json format:\n")
		msgErr.WriteString("{\"ID\": <new id>,\n")
		msgErr.WriteString("\"Descr\": \"<some description text>\"\n")
		msgErr.WriteString("\"Goods\": {\"<tool>\": <price>}}\n")
		msgErr.WriteString("ALL ARGUMENTS ARE NOT REQUIRED")
	case emptySliceErr:
		msgErr.WriteString("There is no any existed receipt yet")
	case wrongIndex:
		msgErr.WriteString("You have to write an index from ")
		msgErr.WriteString(fmt.Sprintf("%v", indexes))
	}

	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		msgErr.String(),
	)
	_, _ = c.bot.Send(msg)
}
