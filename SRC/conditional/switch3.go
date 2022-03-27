package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn something new!")
	default:
		fmt.Println("It's weekend, time to relax!")
	}

	fmt.Println(time.Now().Weekday().String())
}
