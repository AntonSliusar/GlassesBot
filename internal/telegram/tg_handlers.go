package telegram

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) {
	switch message.Command() {
	case "start":
		b.handleStart(message)
	default:
		b.bot.Send(tgbotapi.NewMessage(message.Chat.ID, "Невідома команда"))  // 
	}
}

func (b *Bot) handleCallback(callback *tgbotapi.CallbackQuery) {
	switch callback.Data {
	case "new_order":
		b.handleNewOrder(callback)
	case "active_orders" :
		b.handleActiveOrders(callback)
	default: // TODO
		if strings.HasPrefix(callback.Data, "frame_") {
			b.handleFrameSelection(callback)
		}
		if strings.HasPrefix(callback.Data, "lenses_") {
			b.handleLensesSelection(callback)
		}
		if strings.HasPrefix(callback.Data, "pause_") {
			b.handlePauseAction(callback)
		}
		if strings.HasPrefix(callback.Data, "resume_") {
			b.handleResumeAction(callback)
		}
		if strings.HasPrefix(callback.Data, "finish_") {
			b.handleFinishAction(callback)
		}
	}

}