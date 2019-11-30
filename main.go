package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	flag "github.com/spf13/pflag"
)

var saveFile string = "tasks.json"

type Task struct {
	Tasks []struct {
		Name    string `json:"name"`
		Limit   int    `json:"limit"`
		Current int    `json:"current"`
	} `json:"tasks"`
}

func save(user Task) {
	output, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(saveFile, output, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func listTasks(user Task) {
	for index, currTask := range user.Tasks {
		if currTask.Limit != currTask.Current {
			fmt.Printf("%d: %s [%d/%d]\n", index, currTask.Name, currTask.Current, currTask.Limit)
		}
	}
}

func resetTasks(user Task) {
	fmt.Printf("Resetting tasks in %s\n", saveFile)
	for index, _ := range user.Tasks {
		user.Tasks[index].Current = 0
	}
	save(user)
}

func startTask(user Task, id int) {
	if user.Tasks[id].Current != user.Tasks[id].Limit {
		fmt.Printf("Working on %s\n", user.Tasks[id].Name)
		user.Tasks[id].Current = user.Tasks[id].Current + 1
	} else {
		fmt.Println("You have already met the limit for that task")
	}
	save(user)
}
func main() {
	var id int
	var op string
	var user Task

	flag.IntVarP(&id, "task", "t", 999, "The task you want to start.")
	flag.Parse()
	op = flag.Arg(0)

	data, err := ioutil.ReadFile(saveFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
	}

	switch op {
	case "list":
		listTasks(user)
	case "reset":
		resetTasks(user)
	case "start":
		if id != 999 {
			startTask(user, id)
			break
		}
		fallthrough
	default:
		fmt.Printf("Measure weekly tasks in 2.5 hour blocks\n\nUsage:\n  tracker <command> [options]\nCommands:\n  list\n  reset\n  start\nOptions:\n")
		flag.PrintDefaults()
	}
}
