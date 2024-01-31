package main

import (
	
	"github.com/Syfaro/telegram-bot-api"
	"log"
	
)

var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

// func init() {
// 	// принимаем на входе флаг -telegrambottoken
// 	flag.StringVar(&telegramBotToken, "telegrambottoken", "", "Telegram Bot Token")
// 	flag.Parse()

// 	// без него не запускаемся
// 	if telegramBotToken == "" {
// 		log.Print("-telegrambottoken is required")
// 		os.Exit(1)
// 	}
// }

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

		// chats, err := bot.GetChatAdministrators(tgbotapi.ChatConfig{ChatID: update.Message.Chat.ID})
        // if err != nil {
        //     log.Println(err)
        //     continue
        // }
		// for _, chat := range chats {
        //     chatID := chat.User.ID
        //     msg := tgbotapi.NewMessage(int64(chatID), messageText)
        //     _, err := bot.Send(msg)
        //     if err != nil {
        //         log.Println(err)
        //         continue
        //     }
		// }

		// List of chat IDs to forward messages to
		chatIDs := []int64{-1002112299520, -1001993949188,-4195035189 }

		// Forward the message to all the chats
		for _, chatID := range chatIDs {
			msg := tgbotapi.NewMessage(int64(chatID), messageText)
		//  msg := tgbotapi.NewForward(chatID, update.Message.Chat.ID, update.Message.MessageID)
		 bot.Send(msg)
		}
		// for update := range updates {
		// 	chatID := update.Message.Chat.ID
		// 	chatTitle := update.Message.Chat.Title
		// 	log.Printf("Chat ID: %d, Chat Title: %s", chatID, chatTitle)
		// 	continue
		// }

		// свитч на обработку комманд
		// комманда - сообщение, начинающееся с "/"
		// switch update.Message.Command() {
		// case "start":
		// 	reply = "Привет. Я телеграм-бот"
		// case "hello":
		// 	reply = "world"
		// }

		// // создаем ответное сообщение
		// msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// // отправляем
		// bot.Send(msg)
	}
}