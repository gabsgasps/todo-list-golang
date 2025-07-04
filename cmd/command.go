package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func ParseFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo with title")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.list()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)

		if len(parts) != 2 {
			fmt.Println("Error, invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		todos.edit(index, parts[1])
	default:
		fmt.Println("Invalid command")
	}
}
