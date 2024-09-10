package app

import (
	"github.com/urfave/cli"
	"fmt"
	//"log"
)

//add task []
//update task []
//delete task []
//mark task as in progress []
//list tasks []
//list by status []

const fileName = "tasks.json"

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "GoTask"
	app.Usage = "Command line app made using Go to track your tasks"
	flags := []cli.Flag{
		cli.StringFlag{
			Name: "add",
			Usage: "add a task to the list",
			Value: "",
		},
		cli.StringFlag{
			Name: "update",
			Usage: "update a task in the list",
			Value: "",
		},
		cli.StringFlag{
			Name: "delete",
			Usage: "delete a task to the list",
			Value: "",
		},
		cli.StringFlag{
			Name: "mark-in-progress",
			Usage: "mark a task as 'in progress'",
			Value: "",
		},
		cli.StringFlag{
			Name: "mark-done",
			Usage: "mark a task as 'done'",
			Value: "",
		},
		cli.StringFlag{
			Name: "list",
			Usage: "show the task list",
			Value: "all",
		},
		
	}
	app.Commands = []cli.Command{
		
	}
	return app
}

func listTasks(filter string) {
	taskFilter := []struct{
		Status string
	} {
		{"todo"},
		{"in-progress"},
		{"done"},
		{"all"},
	}
	fmt.Println("Listing:")

	for _, task := range taskFilter {
		if filter == "all" || task.Status == filter {
			fmt.Printf("%s:\n", task.Status)
		}
	}
 }
