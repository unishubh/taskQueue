package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Starting the dispatcher")
	StartDispatcher()
	for i:=0;i<400;i++ {
		Collector("shubham"+strconv.Itoa(i),50)
	}
	time.Sleep(10*time.Second)
}
