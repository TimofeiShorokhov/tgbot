package dao

import (
	"database/sql"
	"tgbot/data"
)

type Repository struct {
	TgRep
}

type TgRep interface {
	SaveData(data *data.Client) error
	GetDataByTgId(id int64) (data.Client, error)
	UpdateReg(id int64, data *data.Client) error
	DeleteReg(id int64) error
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		NewTgPostgres(db),
	}
}
