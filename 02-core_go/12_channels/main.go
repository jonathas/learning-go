// - Channels (pub/sub?)

package main

import (
	"log"

	"github.com/jonathas/learning-go/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)

	// defer => whatever comes after defer, execute that as soon as the current function is done
	defer close(intChan)

	// a go routine runs what's called by itself, in its own routine
	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
