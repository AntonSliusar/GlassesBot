package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func mainMenuKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Нове замовлення", "new_order"),
			tgbotapi.NewInlineKeyboardButtonData("В роботі", "active_orders"),
		),
	)
}

func framesKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Обідок", "frame_1"),
			tgbotapi.NewInlineKeyboardButtonData("Напівобідок", "frame_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Втулки", "frame_3"),
			tgbotapi.NewInlineKeyboardButtonData("Гвинти", "frame_4"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Металева жилка", "frame_5"),
		),
	)
}

func lensesKeyboard() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Пластик", "lenses_1"),
			tgbotapi.NewInlineKeyboardButtonData("Мінерал", "lenses_2"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Полікарбонат", "lenses_3"),
		),
	)
}

func orderActionKeyboard(orderID int, status string) tgbotapi.InlineKeyboardMarkup {
	var actionButton tgbotapi.InlineKeyboardButton
	switch status {
	case "paused":
		actionButton = tgbotapi.NewInlineKeyboardButtonData(
			"Продовжити",
			fmt.Sprintf("resume_%d", orderID),
		)
	case "inWork":
		actionButton = tgbotapi.NewInlineKeyboardButtonData(
			"Пауза",
			fmt.Sprintf("pause_%d", orderID),
		)
	}

	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			actionButton,
			tgbotapi.NewInlineKeyboardButtonData(
				"Завершити", fmt.Sprintf("finish_%d", orderID),
			),
		),
	)
}