package telegram

import (
	"glassesbot/internal/service"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	orderManager *service.OrderManager
	ordersStates map[int64]*service.OrderState
	statesMutex sync.Mutex
}

func NewBot(bot *tgbotapi.BotAPI, manager *service.OrderManager) *Bot {
	return &Bot{
		bot:          bot,
		orderManager: manager,
		ordersStates: make(map[int64]*service.OrderState),
	}
}

func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates := b.bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand(){
			b.handleCommand(update.Message)
		} else if update.CallbackQuery !=nil {
			b.handleCallback(update.CallbackQuery)
		}
	}
}

func (b * Bot) getOrderState(chatID int64) *service.OrderState {
	b.statesMutex.Lock()
	defer b.statesMutex.Unlock()
	return b.ordersStates[chatID]
}

func (b *Bot) SetOrderState(chatID int64, state *service.OrderState) {
	b.statesMutex.Lock()
	defer b.statesMutex.Unlock()
	b.ordersStates[chatID] = state
}

func (b *Bot) ClearOrderState(chatID int64) {
	b.statesMutex.Lock()
	defer b.statesMutex.Unlock()
	delete(b.ordersStates, chatID)
}

func (b *Bot) SendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	b.bot.Send(msg)
}