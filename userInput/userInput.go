package userinput

import (
	"bufio"
	"os"
)

func GetUserInputAndPrint() {
	reader := bufio.NewReader(os.Stdin)
	println("Please enter some input:")
	input, _ := reader.ReadString('\n')
	println("You entered:", input)
}
