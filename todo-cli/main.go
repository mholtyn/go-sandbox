package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

/*
go run main.go add "zrób kawę"
go run main.go list
go run main.go remove 1
*/

// TODO: zdefiniuj struct Task z polami ID (int), Title (string), Done (bool)
type Task struct {
	ID    int
	Title string
	Done  bool
}

// TODO: zdefiniuj struct TodoList z polem tasks (slice of Task)
type TodoList struct {
	tasks []Task
}

// TODO: napisz metode Add(title string) na TodoList ktora dodaje nowe zadanie
// ID powinno byc kolejnym numerem, Done = false
func (t *TodoList) Add(title string) {
	newTask := Task{
		ID:    len(t.tasks) + 1,
		Title: title,
		Done:  false,
	}
	t.tasks = append(t.tasks, newTask)
}

// TODO: napisz metode Remove(id int) error na TodoList
// zwroc ErrNotFound jesli nie ma zadania o tym ID
func (t *TodoList) Remove(id int) error {
	for i, task := range t.tasks {
		if task.ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

// TODO: napisz metode List() na TodoList ktora printuje wszystkie zadania
// format: "[x] 1. Zrob kawe" lub "[ ] 1. Zrob kawe"
func (t *TodoList) List() {
	for _, task := range t.tasks {
		if task.Done {
			fmt.Printf("[x] %d. %s\n", task.ID, task.Title)
		} else {
			fmt.Printf("[ ] %d. %s\n", task.ID, task.Title)
		}
	}
}

// TODO: zdefiniuj sentinel error ErrNotFound
var ErrNotFound = errors.New("Not found")

func main() {
	list := TodoList{}
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("usage: add <title> | remove <id> | list")
		os.Exit(1)
	}

	switch args[0] {
	case "add":
		// TODO: sprawdz czy args[1] istnieje, jesli nie — wypisz blad
		// TODO: wywolaj list.Add()
		if len(args) > 1 {
			list.Add(args[1])
		} else {
			os.Exit(1)
		}
	case "remove":
		// TODO: sprawdz czy args[1] istnieje
		// TODO: skonwertuj args[1] na int (uzyj strconv.Atoi)
		// TODO: wywolaj list.Remove() i obsluz ErrNotFound
		if len(args) != 2 {
			os.Exit(1)
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("failed type convertion")
			os.Exit(1)
		}
		err = list.Remove(id)
		if errors.Is(err, ErrNotFound) {
			fmt.Println(err)
			os.Exit(1)
		}
	case "list":
		// TODO: wywolaj list.List()
		list.List()
	default:
		fmt.Println("unknown command:", args[0])
		os.Exit(1)
	}
}
