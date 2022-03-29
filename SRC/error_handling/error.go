package main

import (
	"fmt"
	"os"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
	} else {
		fmt.Print(employee)
	}
}

func getInformation(id int) (*Employee, error) {
	employee := Employee{LastName: "Pius", FirstName: "Abish"}
	return &employee, nil
}
