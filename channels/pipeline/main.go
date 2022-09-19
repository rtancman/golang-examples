package main

import (
	"fmt"
)

func main() {
	g := generate(2, 4, 6)
	d := divide(g)
	result := sum(d, 10)

	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()
	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)
	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()
	return channel
}

func sum(input chan int, number int) chan int {
	channel := make(chan int)
	go func() {
		for n := range input {
			channel <- n + number
		}
		close(channel)
	}()
	return channel
}
