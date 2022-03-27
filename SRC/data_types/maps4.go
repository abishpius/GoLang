package main

import "fmt"

func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	
	age, exist := studentsAge["christy"]
	if exist {
		fmt.Println("Christy age is", age)
	} else {
		fmt.Println("age couldnt be found")
	}
}
