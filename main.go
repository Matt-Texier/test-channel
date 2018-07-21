package main

import (
	"fmt"
	"time"
)

type Task struct {
	ToBePrinted string
	TimeOut     time.Duration
	Quit        chan bool
}

var myTasks = []*Task{
	{
		ToBePrinted: "coucou1",
		TimeOut:     time.Second * 2,
	},
	{
		ToBePrinted: "coucou2",
		TimeOut:     time.Second * 1,
	},
	{
		ToBePrinted: "coucou3",
		TimeOut:     time.Second * 3,
	},
}

func (myTask *Task) workerTask() {
	for {
		select {
		case <-myTask.Quit:
			{
				fmt.Printf("%s stop working\n", myTask.ToBePrinted)
				return
			}
		default:
			{
				fmt.Println(myTask.ToBePrinted)
				time.Sleep(myTask.TimeOut)
			}
		}
	}
}

func main() {
	for i := 0; i < len(myTasks); i++ {
		myTasks[i].Quit = make(chan bool, 1)
		go myTasks[i].workerTask()
	}

	for {
		var i int
		fmt.Scanf("%d", &i)
		fmt.Println("wants to stop: ", i)
		myTasks[i].Quit <- true
	}
}
