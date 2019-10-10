package main

import "log"

var ch = make(chan int)

func routines() {
	for i := 0; i < 10; i++ {
		log.Printf("I: %d", i)
	}
}

//MyMain func
func MyMain() {
	go routines()
	for i := 0; i < 1; i++ {
		log.Printf("J: %d", i)
	}
}
