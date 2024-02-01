package main

import (
	
	"github.com/Syfaro/telegram-bot-api"
	"log"
	
)



func main() {
	// используя токен создаем новый инстанс бота
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	desiredUserID := 981807694
	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		// универсальный ответ на любое сообщение
		// reply := "Не знаю что сказать"
		if update.Message == nil {
			continue
		}
		messageText := update.Message.Text

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		

		// List of chat IDs to forward messages to
		chatIDs := []int64{-1002003604045,-1002115828429,-1002111555275, }

		// Forward the message to all the chats
		for _, chatID := range chatIDs {
			if update.Message.From.ID == desiredUserID {
				msg := tgbotapi.NewMessage(int64(chatID), messageText)
				//  msg := tgbotapi.NewForward(chatID, update.Message.Chat.ID, update.Message.MessageID)
				 bot.Send(msg)
			} 

	
		}
	
	}
}