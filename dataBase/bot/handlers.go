package bot

import (
	"fmt"
	"log"
	"strconv"
	"telegram-bot-fiber-example/database"

	"github.com/mymmrac/telego/telegoutil"
)

// Обробка команди /start
func HandleStartCommand(userID int64) {
	msg := telegoutil.Message(telegoutil.ID(userID), "Привіт! Ви можете додати завдання, подивитися свої завдання, видалити їх або завершити їх.")
	bot.SendMessage(msg)
}

// Обробка команди /tasks
func HandleTasksCommand(userID int64) {
	tasks, err := database.GetTasks(userID)
	if err != nil {
		log.Printf("Помилка отримання завдань: %v", err)
			return
	}

	response := "Ваші завдання:\n"
	for _, task := range tasks {
		status := "Не виконано"
		if task.IsCompleted {
			status = "Виконано"
		}
		response += fmt.Sprintf("ID: %d, Завдання: %s, Статус: %s\n", task.ID, task.Task, status)
	}

	msg := telegoutil.Message(telegoutil.ID(userID), response)
	bot.SendMessage(msg)
}

// Обробка команди /add_task
func HandleAddTaskCommand(userID int64) {
	msg := telegoutil.Message(telegoutil.ID(userID), "Введіть текст вашого нового завдання:")
	bot.SendMessage(msg)
}

// Обробка команди /delete_task
func HandleDeleteTaskCommand(userID int64) {
	msg := telegoutil.Message(telegoutil.ID(userID), "Введіть ID завдання, яке потрібно видалити:")
	bot.SendMessage(msg)
}

// Обробка команди /complete_task
func HandleCompleteTaskCommand(userID int64) {
	msg := telegoutil.Message(telegoutil.ID(userID), "Введіть ID завдання, яке потрібно завершити:")
	bot.SendMessage(msg)
}

// Обробка інших команд (додавання, видалення, завершення завдань)
func HandleOtherCommands(userID int64, text string) {
	if len(text) > 12 && text[:12] == "нове завдання: "{
	// Додавання завдання
  taskText := text[12:]
	err := database.AddTask(userID, taskText)
	if err != nil {
		log.Printf("Помилка додавання завдання: %v", err)
		return
	}
	msg := telegoutil.Message(telegoutil.ID(userID), "Завдання додано!")
	bot.SendMessage(msg)
} else if len(text) > 16 && text[:16] == "видалити завдання: " {
	// Видалення завдання
	taskIDStr := text[16:]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		msg := telegoutil.Message(telegoutil.ID(userID), "Неправильний формат ID.")
		bot.SendMessage(msg)
		return
	}

	err = database.DeleteTask(userID, taskID)
	if err != nil {
		log.Printf("Помилка видалення завдання: %v", err)
    return
	}
	msg := telegoutil.Message(telegoutil.ID(userID), "Завдання видалено!")
	bot.SendMessage(msg)
} else if len(text) > 18 && text[:18] == "завершити завдання: " {
	// Позначення завдання як завершеного
	taskIDStr := text[18:]
	taskID, err := strconv.Atoi(taskIDStr)
	if err != nil {
		msg := telegoutil.Message(telegoutil.ID(userID), "Неправильний формат ID.")
		bot.SendMessage(msg)
		return
	}

	err = database.CompleteTask(userID, taskID)
	if err != nil {
		log.Printf("Помилка завершення завдання: %v", err)
    return
	  }
		msg := telegoutil.Message(telegoutil.ID(userID), "Завдання завершено!")
		bot.SendMessage(msg)
  }
}