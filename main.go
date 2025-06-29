package main

import (
	config "first_project/config"
	util "first_project/utils"
	"fmt"
)

func main() {
	fmt.Println(util.AddNumbers(3, 4))
	fmt.Println(config.GetConfig())
}
