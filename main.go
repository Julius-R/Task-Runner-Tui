package main

import (
	"errors"
	"log"
	"time"
	
	tea "github.com/charmbracelet/bubbletea"
)

var steps = []Task{
	{
		"Task 1",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return nil
		},
		StatusSuccess,
	},
	{
		"Task 2",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return nil
		},
		StatusSuccess,
	},
	{
		"Task 3",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return nil
		},
		StatusSuccess,
	},
	{
		"Task 4",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return errors.New("fail")
		},
		StatusError,
	},
}

func main() {
	app := tea.NewProgram(InitTasks(steps))
	if _, err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
