package internal

type App struct {
	tasks []Task
}

func NewApp() *App {
	return &App{
		tasks: make([]Task, 0),
	}
}
