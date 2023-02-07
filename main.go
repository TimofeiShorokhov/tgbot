package main

import (
	"log"
	"tgbot/configs"
	"tgbot/dao"
	"tgbot/services"
)
const API = "5892442599:AAF-RhAz3-wmE8DVUkkVvOXb1xowS2WoZGU"
var (
	// глобальная переменная в которой храним токен
	telegramBotToken string
)

/*
var clients map[int]*data.Client

func init() {
	clients = make(map[int]*data.Client)
}

 */
/*
func main() {
	var (
		bot *tgbotapi.BotAPI
		err error
		updChannel tgbotapi.UpdatesChannel
		updConfig tgbotapi.UpdateConfig
		update tgbotapi.Update
		user tgbotapi.User
	)

	cfg, err := configs.ParseConfig()
	if err != nil {
		log.Fatalf("error in config parsing: %s", err)
	}

	db, err := dao.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s",err)
	}

	repository := dao.NewRepository(db)


	bot, err = tgbotapi.NewBotAPI(API)
	if err != nil {
		log.Fatalf("bot init error: %s\n", err)
	}

	user, err = bot.GetMe()
	if err != nil {
		log.Fatalf("bot getme error: %s\n", err)
	}

	fmt.Printf("Authorization succeed, bot is: %s\n", user.FirstName)

	updConfig.Timeout = 60
	updConfig.Limit = 1
	updConfig.Offset = 0

	updChannel, err = bot.GetUpdatesChan(updConfig)
	if err != nil {
		log.Fatalf("bot getUpdatesChan error: %s", err)
	}

	for {
		update = <- updChannel

		if update.Message != nil {
			if update.Message.IsCommand() {
				comText := update.Message.Command()
				if comText == "start" {
					msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID,
						"Привет! Я бот для записи на реснички у Ксении! Все необходимое найдешь в меню :)")
					msgConfig.ReplyMarkup = mainMenu
					bot.Send(msgConfig)
					fmt.Printf("Chat id: %v", update.Message.Chat.ID)
				} else if comText == "menu" {
					msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Главное меню")
					msgConfig.ReplyMarkup = mainMenu
					bot.Send(msgConfig)
				}
			} else {
				if update.Message.Text == mainMenu.Keyboard[0][2].Text {
					msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID,
						"Классика: " + data.Classic +
						"\n2D: " + data.D2 +
						"\n3D: " + data.D3 +
						"\n4D: " + data.D4)
					bot.Send(msgConfig)
				} else if update.Message.Text == mainMenu.Keyboard[0][1].Text{
					clients[update.Message.From.ID] = new(data.Client)
					clients[update.Message.From.ID].State = data.StateName
					fmt.Printf("message: %s\n", update.Message.Text)
					msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID,"Введите Ваше имя:")
					msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
					bot.Send(msgConfig)
				} else {
					cl, ok := clients[update.Message.From.ID]
					if ok {
						if cl.State == data.StateName {
							cl.ID = update.Message.From.ID
							cl.Name = update.Message.Text
							msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите Ваш телефон:")
							bot.Send(msgConfig)
							cl.State = 1
						} else if cl.State == data.StateNumber {
							cl.Number = update.Message.Text
							msgConfig := tgbotapi.NewMessage(update.Message.Chat.ID, "Данные записаны, я Вам перезвоню для потверждения записи :)\n" +
								"Ваши данные: Ваше имя - " + cl.Name + ", Ваш номер - " + cl.Number)
							msgConfig.ReplyMarkup = mainMenu
							//msg := tgbotapi.NewMessage(616237237, "К тебе записался клиент. Данные клиента: Имя - " + cl.Name + ", телефон - " + cl.Number)
							//bot.Send(msg)
							bot.Send(msgConfig)
							delete(clients, update.Message.From.ID)
							repository.SaveData(cl)
							err = dao.Log(cl)
							if err != nil {
								log.Fatalf("error on saving user: %s",err)
							}
						}
						fmt.Printf("state: %+v\n", cl)
					}
				}
			}
		} else {
			fmt.Printf("not message: %+v\n", update)
		}
	}
	bot.StopReceivingUpdates()
}

 */

func main() {
	cfg, err := configs.ParseConfig()
	if err != nil {
		log.Fatalf("error in config parsing: %s", err)
	}

	db, err := dao.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("failed to initialize db: %s",err)
	}

	repository := dao.NewRepository(db)

	service := services.NewService(repository)

	service.TgBotInit(API)
}
