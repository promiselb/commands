package commands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
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

// I use it on my personal machine.
func getJsonObjects[t interface{}](fileName string) ([]t, error) {
	if !strings.HasSuffix(fileName, ".json") {
		fileName = fileName + ".json"
	}
	objs := []t{}
	file, err := os.Open(`C:\Go\src\golang-book\1.20\commands\json\` + fileName)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(file).Decode(&objs)
	return objs, err
}

// Returns a random value from a slice.
func RandomValue[t interface{}](a ...t) t {
	return a[rand.Intn(len(a))]
}
