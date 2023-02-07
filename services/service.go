package services

import (
	"tgbot/dao"
	"tgbot/data"
)

type TgApp interface {
	TgBotInit(api string)
	Register(tgbot TgBot, clients map[int]*data.Client)
	Prices(tgbot TgBot)
	CheckReg(tgbot TgBot)
}

type Service struct {
	TgApp
}

func NewService(rep *dao.Repository) *Service {
	return &Service{
		NewTgService(*rep),
	}
}