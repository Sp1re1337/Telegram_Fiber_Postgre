package database

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
)

// Ініціалізація бази даних
var conn *pgx.Conn

func InitDB() error {
	var err error


	// Підключення до бази даних PostgreSQL
	conn, err = pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/telegram-db")
	if err != nil {
		log.Printf("Помилка підключення до бази даних PostgreSQL: %v", err)
        return errors.New("неможливо підключитися до бази даних")
	}

	err = CreateTasksTable()
	if err != nil {
		log.Printf("Помилка створення таблиці tasks: %v", err)
		return err
	}

	log.Println("База даних успішно ініціалізована.")
	return nil
}