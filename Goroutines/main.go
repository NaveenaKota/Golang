package main

import "fmt"

func doTask(output chan int, done chan bool) {
	for {
		value := <-output
		output <- value * 2
		done <- true
	}
}

func main() {
	outputChan := make(chan int)
	doneChan := make(chan bool, 1)
	go doTask(outputChan, doneChan)
	outputChan <- 1
	value := <-outputChan
	fmt.Println(value)
	for i := 1; i < 10; i++ {
		outputChan <- value
		<-doneChan
		value = <-outputChan
		fmt.Println(value)
	}

	close(outputChan)
	close(doneChan)
}
