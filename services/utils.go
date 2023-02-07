package services

import (
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"tgbot/dao"
	"tgbot/data"
)

var mainMenu = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🏠 Главная"),
		tgbotapi.NewKeyboardButton("🗒 Запись"),
		tgbotapi.NewKeyboardButton("💵 Цены"),
		tgbotapi.NewKeyboardButton("Проверить запись"),
		tgbotapi.NewKeyboardButton("Редактировать запись"),
	),
)

type TgBot struct {
	err error
	bot *tgbotapi.BotAPI
	updChannel tgbotapi.UpdatesChannel
	updConfig tgbotapi.UpdateConfig
	update tgbotapi.Update
	user tgbotapi.User
}

type TgService struct {
	repo dao.Repository
}

func NewTgService(repo dao.Repository) *TgService {
	return &TgService{repo: repo}
}

func(s *TgService) TgBotInit(api string) {
	tgbot := TgBot{}
	err := tgbot.err

	var clients map[int]*data.Client

	clients = make(map[int]*data.Client)

	tgbot.bot, err = tgbotapi.NewBotAPI(api)
	if err != nil {
		log.Fatalf("bot init error: %s\n", err)
	}

	tgbot.user, err = tgbot.bot.GetMe()
	if err != nil {
		log.Fatalf("bot getme error: %s\n", err)
	}

	fmt.Printf("Authorization succeed, bot is: %s\n", tgbot.user.FirstName)

	tgbot.updConfig.Timeout = 60
	tgbot.updConfig.Limit = 1
	tgbot.updConfig.Offset = 0

	tgbot.updChannel, err = tgbot.bot.GetUpdatesChan(tgbot.updConfig)
	if err != nil {
		log.Fatalf("bot getUpdatesChan error: %s", err)
	}

	for {
		tgbot.update = <- tgbot.updChannel

		if tgbot.update.Message != nil {
			if tgbot.update.Message.IsCommand() {
				comText := tgbot.update.Message.Command()
				if comText == "start" {
					msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID,
						"Привет! Я бот для записи на реснички у Ксении! Все необходимое найдешь в меню :)")
					msgConfig.ReplyMarkup = mainMenu
					tgbot.bot.Send(msgConfig)
					fmt.Printf("Chat id: %v", tgbot.update.Message.Chat.ID)
				} else if comText == "menu" {
					msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Главное меню")
					msgConfig.ReplyMarkup = mainMenu
					tgbot.bot.Send(msgConfig)
				}
				} else {
					s.Register(tgbot,clients)
					s.Prices(tgbot)
					s.CheckReg(tgbot)
				}
		} else {
			fmt.Printf("not message: %+v\n", tgbot.update)
		}
	}
	tgbot.bot.StopReceivingUpdates()
}

func(s *TgService) Register(tgbot TgBot, clients map[int]*data.Client) {
	if tgbot.update.Message.Text == mainMenu.Keyboard[0][1].Text {
		clients[tgbot.update.Message.From.ID] = new(data.Client)
		clients[tgbot.update.Message.From.ID].State = data.StateName
		fmt.Printf("message: %s\n", tgbot.update.Message.Text)
		msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Введите Ваше имя:")
		msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		tgbot.bot.Send(msgConfig)
	} else {
		cl, ok := clients[tgbot.update.Message.From.ID]
		if ok {
			if cl.State == data.StateName {
				cl.ID = tgbot.update.Message.From.ID
				cl.Name = tgbot.update.Message.Text
				msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Введите Ваш телефон:")
				tgbot.bot.Send(msgConfig)
				cl.State = 1
			} else if cl.State == data.StateNumber {
				cl.Number = tgbot.update.Message.Text
				msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Данные записаны, я Вам перезвоню для потверждения записи :)\n"+
					"Ваши данные: Ваше имя - "+cl.Name+", Ваш номер - "+cl.Number)
				msgConfig.ReplyMarkup = mainMenu
				//msg := tgbotapi.NewMessage(616237237, "К тебе записался клиент. Данные клиента: Имя - " + cl.Name + ", телефон - " + cl.Number)
				//bot.Send(msg)
				tgbot.bot.Send(msgConfig)
				delete(clients, tgbot.update.Message.From.ID)
				s.repo.SaveData(cl)
			}
			fmt.Printf("state: %+v\n", cl)
		}
	}
}

func (s *TgService) Prices(tgbot TgBot) {
	if tgbot.update.Message.Text == mainMenu.Keyboard[0][2].Text {
		msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID,
			"Классика: "+data.Classic+
				"\n2D: "+data.D2+
				"\n3D: "+data.D3+
				"\n4D: "+data.D4)
		tgbot.bot.Send(msgConfig)
	}
}

func (s *TgService) CheckReg(tgbot TgBot) {
	if tgbot.update.Message.Text == mainMenu.Keyboard[0][3].Text {
		data, err := s.repo.GetDataByTgId(tgbot.update.Message.Chat.ID)
		if err != nil {
			fmt.Printf("error in service while getting data: %s", err)
		}
		msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Вы записаны и вот Ваши данные:\n" +
			"Ваше имя: " +data.Name + ", Ваш телефон: " + data.Number)
		msgConfig.ReplyMarkup = mainMenu
		tgbot.bot.Send(msgConfig)
	}
}

//func (s *TgService) DeleteReg() {}

func (s *TgService) UpdateReg(tgbot TgBot, clients map[int]*data.Client ) {
	if tgbot.update.Message.Text == mainMenu.Keyboard[0][4].Text {
	clients[tgbot.update.Message.From.ID] = new(data.Client)
	clients[tgbot.update.Message.From.ID].State = data.StateName
	fmt.Printf("message: %s\n", tgbot.update.Message.Text)
	msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Введите Ваше имя:")
	msgConfig.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	tgbot.bot.Send(msgConfig)
} else {
	cl, ok := clients[tgbot.update.Message.From.ID]
	if ok {
		if cl.State == data.StateName {
			cl.ID = tgbot.update.Message.From.ID
			cl.Name = tgbot.update.Message.Text
			msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Введите Ваш телефон:")
			tgbot.bot.Send(msgConfig)
			cl.State = 1
		} else if cl.State == data.StateNumber {
			cl.Number = tgbot.update.Message.Text
			msgConfig := tgbotapi.NewMessage(tgbot.update.Message.Chat.ID, "Данные обновлены\n"+
				"Ваши данные: Ваше имя - "+cl.Name+", Ваш номер - "+cl.Number)
			msgConfig.ReplyMarkup = mainMenu
			//msg := tgbotapi.NewMessage(616237237, "К тебе записался клиент. Данные клиента: Имя - " + cl.Name + ", телефон - " + cl.Number)
			//tgbot.bot.Send(msg)
			tgbot.bot.Send(msgConfig)
			delete(clients, tgbot.update.Message.From.ID)
			s.repo.UpdateReg(tgbot.update.Message.Chat.ID, cl)
		}
		fmt.Printf("state: %+v\n", cl)
	}
}}
