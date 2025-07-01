package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["srimanta"] = 1
	m["rick"] = 2

	value, exists := m["gopal"]

	fmt.Println(value, exists)

}
