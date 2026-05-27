package main

import (
	"errors"
	"fmt"
	"os"
)

/*
go run main.go add "zrób kawę"
go run main.go list
go run main.go remove 1
*/

// TODO: zdefiniuj struct Task z polami ID (int), Title (string), Done (bool)
type Task struct {
	ID int
	Title string
	Done bool
}
// TODO: zdefiniuj struct TodoList z polem tasks (slice of Task)
type TodoList struct {
	tasks []Task
}

// TODO: napisz metode Add(title string) na TodoList ktora dodaje nowe zadanie
// ID powinno byc kolejnym numerem, Done = false

// TODO: napisz metode Remove(id int) error na TodoList
// zwroc ErrNotFound jesli nie ma zadania o tym ID

// TODO: napisz metode List() na TodoList ktora printuje wszystkie zadania
// format: "[x] 1. Zrob kawe" lub "[ ] 1. Zrob kawe"

// TODO: zdefiniuj sentinel error ErrNotFound

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
	case "remove":
		// TODO: sprawdz czy args[1] istnieje
		// TODO: skonwertuj args[1] na int (uzyj strconv.Atoi)
		// TODO: wywolaj list.Remove() i obsluz ErrNotFound
	case "list":
		// TODO: wywolaj list.List()
	default:
		fmt.Println("unknown command:", args[0])
		os.Exit(1)
	}

	_ = errors // zeby kompilator nie krzyczal zanim uzyjesz
}