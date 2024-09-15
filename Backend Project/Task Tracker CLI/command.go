package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add       string
	Edit      string
	Status    string 
	Del       int
	Toggle    int
	List      bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task description")
	flag.StringVar(&cf.Status, "status", "", "Set a task status (id:new_status)")
	flag.IntVar(&cf.Del, "del", -1, "Delete a task")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task status")
	flag.BoolVar(&cf.List, "list", false, "List all tasks")

	flag.Parse()

	return &cf
}

func (cf *CmdFlags) Execute(trackers *Trackers) {
	switch {
	case cf.List:
		trackers.print()
	case cf.Add != "":
		trackers.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit format. Please use id:new_description")
			os.Exit(1)
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid task id for edit.")
			os.Exit(1)
		}
		trackers.edit(id, parts[1])
	case cf.Status != "":
		parts := strings.SplitN(cf.Status, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid status format. Please use id:new_status")
			os.Exit(1)
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid task id for status.")
			os.Exit(1)
		}
		status := Status(parts[1])
		if status != Todo && status != InProgress && status != Done {
			fmt.Println("Invalid status. Allowed values are 'todo', 'in-progress', 'done'")
			os.Exit(1)
		}
		err = trackers.setStatus(id, status)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case cf.Toggle != -1:
		trackers.toggle(cf.Toggle)
	case cf.Del != -1:
		trackers.delete(cf.Del)
	default:
		fmt.Println("No valid command provided")
	}
}
