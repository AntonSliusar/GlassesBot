package telegram

import (
	"fmt"
	"glassesbot/internal/domain"
	"glassesbot/internal/service"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleStart (message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "👓")
	msg.ReplyMarkup = mainMenuKeyboard()
	b.bot.Send(msg)
}

func (b *Bot) handleNewOrder(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID

	orderID := b.orderManager.CreateOrder()

	b.SetOrderState(chatID, &service.OrderState{
		OrderId: orderID,
		Stage: service.STAGE_AWAITING_FRAME,
	})

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"Оберіть тип оправи",
	)
	keyboard := framesKeyboard()
	msg.ReplyMarkup = &keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleActiveOrders(callback *tgbotapi.CallbackQuery) {
	orders := b.orderManager.GetAllOrders()
	if len(orders) == 0 {
		msg := tgbotapi.NewEditMessageText(
			callback.Message.Chat.ID,
			callback.Message.MessageID,
			"Активних замволнень немає",
		)
		b.bot.Send(msg)
		return
	}

	for id, order := range orders {
		text := fmt.Sprintf(
			"Тип оправи: %s\nТип лінз: %s\nСтатус: %s",
			order.Frame, order.Lenses, order.Status,
		)
		msg := tgbotapi.NewMessage(callback.Message.Chat.ID, text)
		fmt.Println(order.Status)
		fmt.Println(order.Frame)
		keyboard := orderActionKeyboard(int(id), order.Status)
		msg.ReplyMarkup = &keyboard
		b.bot.Send(msg)
	}
}

func (b *Bot) handleFrameSelection(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	orderState := b.getOrderState(chatID)
	if orderState == nil {
		orderId := b.orderManager.CreateOrder()
		b.SetOrderState(chatID, &service.OrderState{
			OrderId: orderId,
			Stage: service.STAGE_AWAITING_LENSES,
		})
	}
	b.orderManager.Orders[orderState.OrderId].Frame = domain.GetFrameByID(callback.Data)
	orderState.Stage = service.STAGE_AWAITING_LENSES

	b.SetOrderState(chatID, orderState)

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"Оберіть тип лінз",
	)
	keyboard := lensesKeyboard()
	msg.ReplyMarkup = &keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleLensesSelection(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	orderState := b.getOrderState(chatID)
	if orderState == nil {
		b.SendMessage(chatID, "Помилка: створіть нове замовлення")
	}

	b.orderManager.Orders[orderState.OrderId].Lenses = domain.GetLensesByID(callback.Data)
	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"Замовлення створено!",
	)
	keyboard := mainMenuKeyboard()
	msg.ReplyMarkup = &keyboard
	b.bot.Send(msg)
	b.ClearOrderState(chatID)
}

func (b *Bot) handlePauseAction(callback *tgbotapi.CallbackQuery) {
	orderID, err := strconv.ParseInt(strings.TrimPrefix(callback.Data, "pause_"), 10, 64)
	
	if err != nil {
		b.SendMessage(callback.Message.Chat.ID, "Невірний ID замовлення")
		return
	}
	b.orderManager.PauseOrder(orderID)

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"Замовлення поставлено на паузу",
	)
	keyboard := mainMenuKeyboard()
	msg.ReplyMarkup = &keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleResumeAction(callback *tgbotapi.CallbackQuery) {
	orderID, err := strconv.ParseInt(strings.TrimPrefix(callback.Data, "resume_"), 10, 64)
	if err != nil {
		b.SendMessage(callback.Message.Chat.ID, "Невірний ID замовлення")
		return
	}
	b.orderManager.ResumeOrder(orderID)

	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"Замовлення в роботі",
	)
	keyboard := mainMenuKeyboard()
	msg.ReplyMarkup = &keyboard
	b.bot.Send(msg)
}

func (b *Bot) handleFinishAction(callback *tgbotapi.CallbackQuery) {
	orderID, err := strconv.ParseInt(strings.TrimPrefix(callback.Data, "finish_"), 10, 64)
	if err != nil {
		b.SendMessage(callback.Message.Chat.ID, "Невірний ID замовлення")
		return
	}
fmt.Println(b.orderManager.Orders[orderID])

	b.orderManager.FinishOrder(orderID)
	msg := tgbotapi.NewEditMessageText(
		callback.Message.Chat.ID,
		callback.Message.MessageID,
		"Замовлення завершено",
	)
	keyboard := mainMenuKeyboard()
	msg.ReplyMarkup = &keyboard
	b.bot.Send(msg)
}

