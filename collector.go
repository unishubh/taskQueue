package main

type Task struct {
	delay int
	name string
}
 var WorkQueue = make(chan Task,1000)
func Collector(name string, delay int) {

	currentTask := Task{
		name:name,
		delay:delay,
	}

	WorkQueue <- currentTask
	return
}
