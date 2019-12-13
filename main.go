package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	flag "github.com/spf13/pflag"
)

var filepath string

type TaskFile struct {
	Tasks []struct {
		Name    string `json:"name"`
		Limit   int    `json:"limit"`
		Current int    `json:"current"`
	} `json:"tasks"`
}

func listTasks(user TaskFile) {
	for index, currTask := range user.Tasks {
		if currTask.Limit != currTask.Current {
			fmt.Printf("%d: %s [%d/%d]\n", index, currTask.Name, currTask.Current, currTask.Limit)
		}
	}
}

func resetTasks(user TaskFile) TaskFile {
	fmt.Printf("Resetting tasks")
	for index, _ := range user.Tasks {
		user.Tasks[index].Current = 0
	}
	return user
}

func startTask(user TaskFile, id int) TaskFile {
	if user.Tasks[id].Current != user.Tasks[id].Limit {
		fmt.Printf("Working on %s\n", user.Tasks[id].Name)
		user.Tasks[id].Current = user.Tasks[id].Current + 1
	} else {
		fmt.Println("You have already met the limit for that task")
	}
	return user
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

//May only definitively prove it doesn't exist.
//Beware false positives
func exists(filepath string) bool {
	_, err := os.Stat(filepath)
	return os.IsNotExist(err)
}

func loadTasks() (TaskFile, error) {
	var ret TaskFile
	filepath = "tasks.json"
	if !exists(filepath) {
		dir := os.Getenv("HOME") + "/.config/tracker"
		filepath = dir + "/data.json"
        if err := os.MkdirAll(dir, 0755); err != nil {
			if !exists(filepath) {
				return ret, ioutil.WriteFile(filepath, []byte("{tasks:[]}"), 0644)
			}
		} else {
			return ret, err
		}
	}

    data, err := ioutil.ReadFile(filepath)
	if err == nil {
		err = json.Unmarshal(data, &ret)
	}
	return ret, err
}

func main() {
	var id int
	var op string

	flag.IntVarP(&id, "task", "t", 999, "The task you want to start.")
	flag.Parse()
	op = flag.Arg(0)

	tfile, err := loadTasks()
	panicIfErr(err)

	newTFile := tfile

	switch op {
	case "list":
		listTasks(tfile)
	case "reset":
		newTFile = resetTasks(tfile)
	case "start":
		if id != 999 {
			newTFile = startTask(tfile, id)
			break
		}
		fallthrough
	default:
		fmt.Printf("Measure weekly tasks in 2.5 hour blocks\n\nUsage:\n  tracker <command> [options]\nCommands:\n  list\n  reset\n  start\nOptions:\n")
		flag.PrintDefaults()
	}

	if !reflect.DeepEqual(newTFile, tfile) {
		output, err := json.Marshal(newTFile)
		panicIfErr(err)
		err = ioutil.WriteFile(filepath, output, 0644)
		panicIfErr(err)
	}
}
