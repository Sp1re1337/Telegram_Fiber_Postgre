package main

import (
	"log"
	  "telegram-bot-fiber-example/bot"
	  "telegram-bot-fiber-example/database"

	  "github.com/gofiber/fiber/v2"
)

func main() {
	// Ініціалізація бази даних
	err := database.InitDB()
	if err != nil {
		log.Fatalf("Помилка ініціалізації бази даних %v", err)
	}

	// Створюємо Fiber додаток
	app := fiber.New()


	// Налаштовуємо маршрут для вебхуків
	app.Post("/webhook", bot.HandleWebhook)


	// Запуск серверу
	err = app.Listen(":Port")
	if err != nil {
		log.Fatalf("Помилка запуску сервера: %v", err)
	}
}
