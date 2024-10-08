package main

import (
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

var steps = []Task{
	{
		"PGS Helm Chart Values Release Update",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return nil
		},
		StatusSuccess,
	},
	{
		"Imperium Release Update",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return nil
		},
		StatusSuccess,
	},
	{
		"PGS Helm Chart Release",
		func(...any) error {
			time.Sleep(time.Second * 2)
			return nil
		},
		StatusSuccess,
	},
}

func main() {
	app := tea.NewProgram(InitTasks(steps))
	if _, err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
