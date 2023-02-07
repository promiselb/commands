package commands

import (
	"fmt"
	"math/rand"
)

// Generates a random bool.
func RandomBool() bool {
	return rand.Intn(100) > 50
}

// Prints the commands.
func Help() string {
	var s = "Command: Description\n\n"

	for _, v := range Commands {
		s += fmt.Sprintf("%s: %s\n\n", v.Name, v.Description)
	}
	return s
}

// Used to split the args with <...> for some commands.
func SplitArgs(arg string) []string {
	slice := []string{}
	j := 0
	arg = " " + arg
	for i := 0; i < len(arg); i++ {
		x := string(arg[i])
		if x == "<" {
			j = i
		} else if x == ">" && j != 0 {
			slice = append(slice, arg[j:i+1])
			j = 0
		}
	}
	return slice
}
