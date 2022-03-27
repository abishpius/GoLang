package main

import "fmt"

func main() {
	firstName := "Abish"
	updateName(&firstName)
	fmt.Println(firstName)
}

func updateName(name *string) {
	*name = "Abish Pius"
}
