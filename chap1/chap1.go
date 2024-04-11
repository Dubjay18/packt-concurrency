package chap1

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	i    int
	max  int
	text string
}

func outputText(j *Job, goGroup *sync.WaitGroup) {
	for j.i < j.max {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(j.text)
		j.i++
	}
	goGroup.Done()
}
func Waitg() {
	goGroup := new(sync.WaitGroup) //declare a new sync.WaitGroup
	fmt.Println("Starting")        //print "Starting"
	hello := new(Job)              //declare a new Job struct
	hello.text = "hello"           //set the text field of the Job struct to "hello"
	hello.i = 0                    //set the i field of the Job struct to 0
	hello.max = 2                  //set the max field of the Job struct to 2
	world := new(Job)              //declare a new Job struct
	world.text = "world"           //set the text field of the Job struct to "world"
	world.i = 0                    //set the i field of the Job struct to 0
	world.max = 2                  //set the max field of the Job struct to 2
	go outputText(hello, goGroup)  //call the outputText function with the hello Job struct and the goGroup sync.WaitGroup
	go outputText(world, goGroup)
	goGroup.Add(2)
	goGroup.Wait()
}
