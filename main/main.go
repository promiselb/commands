package main

import (
	"fmt"

	"github.com/promiselb/commands"
)

func main() {
	fmt.Println(commands.StdExecuter.Execute("$help"))
}
