package internal

import (
	"fmt"
	"strings"
)

type TodoApp struct {
	tasks  []Task
	nextID int
}

func NewTodoApp() *TodoApp {
	return &TodoApp{
		tasks:  make([]Task, 0),
		nextID: 1,
	}
}

func (app *TodoApp) AddTask(title string) {
	task := Task{
		ID:        app.nextID,
		Title:     title,
		Completed: false,
	}
	app.tasks = append(app.tasks, task)
	app.nextID++
}

func (app *TodoApp) ListTasks() []Task {
	return app.tasks
}

func (app *TodoApp) CompleteTask(id int) error {
	for i := range app.tasks {
		if app.tasks[i].ID == id {
			app.tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (app *TodoApp) DeleteTask(id int) error {
	for i, task := range app.tasks {
		if task.ID == id {
			app.tasks = append(app.tasks[:i], app.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}

func (app *TodoApp) Search(query string) []Task {
	var results []Task
	query = strings.ToLower(query)
	for _, task := range app.tasks {
		if strings.Contains(strings.ToLower(task.Title), query) {
			results = append(results, task)
		}
	}
	return results
}
