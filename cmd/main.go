package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aosmankin/todoapp/internal"
)

const dataFile = "tasks.json"

func main() {
	app := internal.NewTodoApp()
	if err := app.LoadFromFile(dataFile); err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
	}

	defer func() {
		if err := app.SaveToFile(dataFile); err != nil {
			fmt.Printf("Error saving tasks: %v\n", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			addTask(app, scanner)
		case "2":
			listTasks(app)
		case "3":
			completeTask(app, scanner)
		case "4":
			deleteTask(app, scanner)
		case "5":
			searchTasks(app, scanner)
		case "6":
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func printMenu() {
	fmt.Println("\nMenu")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Search Tasks")
	fmt.Println("6. Exit")
	fmt.Print("Enter your choice: ")
}

func addTask(app *internal.TodoApp, scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())
	if title == "" {
		fmt.Println("Title cannot be empty")
		return
	}
	app.AddTask(title)
	fmt.Println("Task added successfully!")
}

func listTasks(app *internal.TodoApp) {
	tasks := app.ListTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	fmt.Println("\nTask List:")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "OK"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
	}
}

func completeTask(app *internal.TodoApp, scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to complete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	if err := app.CompleteTask(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task marked as completed")
}

func deleteTask(app *internal.TodoApp, scanner *bufio.Scanner) {
	fmt.Print("Enter task ID to delete: ")
	scanner.Scan()
	id, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	if err := app.DeleteTask(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Task deleted successfully")
}

func searchTasks(app *internal.TodoApp, scanner *bufio.Scanner) {
	fmt.Print("Enter search query: ")
	scanner.Scan()
	query := scanner.Text()

	tasks := app.Search(query)
	if len(tasks) == 0 {
		fmt.Println("No matching tasks found")
		return
	}

	fmt.Println("\nSearch Results:")
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", task.ID, status, task.Title)
	}
}
