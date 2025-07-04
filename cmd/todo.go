package cmd

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) list() {
	table := table.New(os.Stdout)
	table.SetRowLines(true)

	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, row := range *todos {
		completed := "❌"
		completedAt := ""

		if row.Completed {
			completed = "✅"
			if row.CompletedAt != nil {
				completedAt = row.CompletedAt.Format(time.RFC1123) //time standard
			}
		}

		table.AddRow(strconv.Itoa(index), row.Title, completed, row.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}

func (todos *Todos) add(content string) {
	item := Todo{
		Title:       content,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	*todos = append(*todos, item)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := *todos
	todo := &t[index]

	if todo.Completed {
		completedTime := time.Now()
		todo.CompletedAt = &completedTime
	} else {
		todo.CompletedAt = nil
	}

	todo.Completed = !todo.Completed
	return nil
}

func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)

	return nil
}
