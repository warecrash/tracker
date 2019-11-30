package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

var saveFile string = "tasks.json"

type Task struct {
	Tasks []struct {
		Name  string `json:"name"`
		Limit int    `json:"limit"`
		Current int  `json:"current"`
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

func getTasks(user Task) {

	for index,currTask := range user.Tasks {

		if currTask.Limit != currTask.Current {
			fmt.Printf("%d: %s [%d/%d]\n", index, currTask.Name, currTask.Current, currTask.Limit)
		}
	}
}

func resetTasks(user Task) {
	for index,_ := range user.Tasks{
		user.Tasks[index].Current = 0
	}
	save(user)
}

func startTask(user Task, id int) {
	if user.Tasks[id].Current != user.Tasks[id].Limit {
		fmt.Printf("Working on: %s\n",user.Tasks[id].Name)
		user.Tasks[id].Current = user.Tasks[id].Current+1
	} else {
		fmt.Println("You have already met the limit for that task")
	}
	save(user)
}
func main() {
	//id,err := strconv.Atoi(os.Args[1])
	//if err != nil {
	//	fmt.Println(err)
	//}
	data, err := ioutil.ReadFile(saveFile)
	if err != nil {
		fmt.Println(err)
	}


	var user Task

	err = json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
	}
	//startTask(user, id)
	getTasks(user)
	//resetTasks(user)

}
