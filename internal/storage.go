package internal

import (
	"encoding/json"
	"os"
)

func (app *TodoApp) SaveToFile(filename string) error {
	data, err := json.Marshal(app.tasks)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (app *TodoApp) LoadFromFile(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	if len(data) == 0 {
		return nil
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return err
	}

	app.tasks = tasks
	if len(tasks) > 0 {
		app.nextID = tasks[len(tasks)-1].ID + 1
	}
	return nil
}
