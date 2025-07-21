package internal

var nextID = 1

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Task struct {
	id          int
	title       string
	isCompleted bool
	priority    Priority
}

func NewTask(title string, priority Priority) Task {
	task := Task{
		id:          nextID,
		title:       title,
		isCompleted: false,
		priority:    priority,
	}
	nextID++
	return task
}

func (t *Task) ChangeStatus(string, error) {

}
