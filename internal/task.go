package internal

var nextID = 1

type Task struct {
	id          int
	title       string
	isCompleted bool
}

func NewTask(title string) Task {
	task := Task{
		id:          nextID,
		title:       title,
		isCompleted: false,
	}
	nextID++
	return task
}

func (t *Task) ChangeStatus(string, error) {

}
