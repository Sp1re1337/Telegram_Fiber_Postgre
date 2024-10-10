package bot

import (
	"github.com/gofiber/fiber/v2"
  "github.com/mymmrac/telego"
)


// Створюємо нового бота (це глобальний об'єкт)
var bot, _ = telego.NewBot("ToKeN")


// Обробка вебхука
func HandleWebhook(c *fiber.Ctx) error {
	var update telego.Update
	if err := c.BodyParser(&update); err != nil {
		  return c.Status(fiber.StatusBadRequest).SendString("Неправильний формат")
	}


	// Перевіряємо, чи є повідомлення
	if update.Message != nil {
		userID := update.Message.Chat.ID
		text := update.Message.Text


		 // Обробляємо команди
		if text == "/start" {
			HandleStartCommand(userID)
	} else if text == "/tasks" {
			HandleTasksCommand(userID)
	} else if text == "/add_task" {
			HandleAddTaskCommand(userID)
	} else if text == "/delete_task" {
			HandleDeleteTaskCommand(userID)
	} else if text == "/complete_task" {
			HandleCompleteTaskCommand(userID)
	} else {
			// Інші дії, такі як додавання, видалення та завершення завдань
			HandleOtherCommands(userID, text)
	  }
	}

	return c.SendStatus(fiber.StatusOK)
}