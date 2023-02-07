package dao

import (
	"database/sql"
	"fmt"
	"log"
	"tgbot/data"
)

type TgPostgres struct {
	db *sql.DB
}

func NewTgPostgres(db *sql.DB) *TgPostgres {
	return &TgPostgres{db: db}
}

func (q *TgPostgres) SaveData(data *data.Client) error {

	transaction, err := q.db.Begin()
	if err != nil {
		log.Printf("error with db.Begin: %s\n", err)
	}

	query := `INSERT INTO clients(user_name, user_number, user_tg_id) VALUES($1,$2,$3)`

	_, err = transaction.Exec(query, data.Name, data.Number, data.ID)
	if err != nil {
		transaction.Rollback()
		return err
	}
	fmt.Println("Данные записаны в базу")
	return transaction.Commit()
}

func (q *TgPostgres) GetDataByTgId(id int64) (data.Client, error) {
	var result data.Client
	selectValue := `SELECT user_name,user_number FROM clients WHERE user_tg_id = $1`
	get, err := q.db.Query(selectValue, id)

	if err != nil {
		log.Println("error of getting data from database: " + err.Error())
		return data.Client{}, err
	}

	for get.Next() {
		err = get.Scan(&result.Name, &result.Number)
	}
	return result, nil
}

func (q *TgPostgres) DeleteReg(id int64) error {
	transaction, err := q.db.Begin()

	if err != nil {
		log.Println("error with database: " + err.Error())
	}
	query := `DELETE FROM clients where user_tg_id = $1`

	_, err = transaction.Exec(query, id)
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit()
}

func (q *TgPostgres) UpdateReg(id int64, data *data.Client) error {
	transaction, err := q.db.Begin()

	if err != nil {
		log.Println("error with database: " + err.Error())
	}
	query := `UPDATE clients SET user_name = $1, user_number = $2 where user_tg_id = $3`

	_, err = transaction.Exec(query, data.Name, data.Number, id)
	if err != nil {
		transaction.Rollback()
		return err
	}
	return transaction.Commit()
}
