package main

import (
	"fmt"
	
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TaskStatus string

type Task struct {
	Name   string
	Run    func(...any) error
	Status TaskStatus
}

const (
	StatusSuccess = "success"
	StatusError   = "error"
)

var (
	updateStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#DB4203"))
	doneStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
)

var StatusIcons = map[TaskStatus]string{
	StatusSuccess: "✨",
	StatusError:   "❌",
}

func (t *Task) GetStatus() string {
	return fmt.Sprintf("%s %s", StatusIcons[t.Status], t.Name)
}

func (t *Task) SetStatus(status TaskStatus) {
	t.Status = status
}

func (t *Task) GetName() string {
	return t.Name
}

type updateCmd struct{}
type startCmd struct{}

type TaskTui struct {
	tasks         []Task
	currTask      int
	tasksComplete bool
	spinner       spinner.Model
}

func (tt *TaskTui) RunCurrentTask() tea.Msg {
	err := tt.tasks[tt.currTask].Run()
	tt.tasks[tt.currTask].SetStatus(
		TaskStatus(
			chooseBetween(
				err != nil, StatusError, StatusSuccess,
			),
		),
	)
	return updateCmd{}
}

func (tt *TaskTui) Init() tea.Cmd {
	return tea.Batch(
		tt.spinner.Tick, func() tea.Msg {
			return startCmd{}
		},
	)
}

func (tt *TaskTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return tt, tea.Quit
		}
	
	case startCmd:
		return tt, tt.RunCurrentTask
	
	case updateCmd:
		if tt.currTask >= len(tt.tasks)-1 {
			tt.tasksComplete = true
			ms := tt.tasks[tt.currTask].GetStatus()
			return tt, tea.Sequence(
				tea.Printf("%s completed", ms),
				tea.Printf(doneStyle.Render("Updates completed")),
				tea.Quit,
			)
		}
		ms := tt.tasks[tt.currTask].GetStatus()
		tt.currTask++
		return tt, tea.Batch(tea.Printf("%s completed", ms), tt.RunCurrentTask)
	}
	
	var cmd tea.Cmd
	tt.spinner, cmd = tt.spinner.Update(msg)
	return tt, cmd
	
}

func (tt *TaskTui) View() string {
	return fmt.Sprintf(
		"%s Updating %s",
		tt.spinner.View(),
		updateStyle.Render(
			tt.tasks[tt.currTask].GetName(),
		),
	)
}

func chooseBetween[T any](comparator bool, item1, item2 T) T {
	if comparator {
		return item1
	}
	return item2
}

func InitTasks(items []Task) *TaskTui {
	spin := spinner.New()
	spin.Spinner = spinner.Dot
	return &TaskTui{
		spinner:  spin,
		currTask: 0,
		tasks:    items,
	}
}
