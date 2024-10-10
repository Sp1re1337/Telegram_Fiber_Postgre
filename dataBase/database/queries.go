package database

import (
	"context"
	"log"
)

// Структура для завдань
type Task struct {
	ID          int
	Task        string
	IsCompleted bool
}

// Створюємо таблицю для завдань
func CreateTasksTable() error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        user_id BIGINT,
        task TEXT,
        is_completed BOOLEAN DEFAULT FALSE
    );`
		_, err := conn.Exec(context.Background(), createTableSQL)
      if err != nil {
				log.Printf("Помилка виконання запиту на створення таблиці: %v", err)
			}
		return err
}

// Додавання завдання
func AddTask(userID int64, task string) error {
	query := `INSERT INTO tasks (user_id, task) VALUES ($1, $2)`
	_, err := conn.Exec(context.Background(), query, userID, task)
		if err != nil {
			log.Printf("Помилка додавання завдання для користувача %d: %v", userID, err)
		}
	return err
}

// Отримання списку завдань
func GetTasks(userID int64) ([]Task, error) {
	query := `SELECT id, task, is_completed FROM tasks WHERE user_id = $1`
	rows, err := conn.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("Помилка виконання запиту на отримання завдань для користувача %d: %v", userID, err)
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err = rows.Scan(&task.ID, &task.Task, &task.IsCompleted)
		if err != nil {
			log.Printf("Помилка сканування рядка для завдання користувача %d: %v", userID, err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Помилка під час ітерації по результатам запиту для користувача %d: %v", userID, err)
		return nil, err
}

	return tasks, nil
}

// Видалення завдання за ID
func DeleteTask(userID int64, taskID int) error {
	query := `DELETE FROM tasks WHERE user_id = $1 AND id = $2`
	_, err := conn.Exec(context.Background(), query, userID, taskID)
	if err != nil {
		log.Printf("Помилка видалення завдання з ID %d для користувача %d: %v", taskID, userID, err)
    }
		return err
	}


// Позначення завдання як виконаного
func CompleteTask(userID int64, taskID int) error {
	query := `UPDATE tasks SET is_completed = TRUE WHERE user_id = $1 AND id = $2`
	_, err := conn.Exec(context.Background(), query, userID, taskID)
	if err != nil {
		log.Printf("Помилка позначення завдання з ID %d як виконаного для користувача %d: %v", taskID, userID, err)
	}
	return err
}