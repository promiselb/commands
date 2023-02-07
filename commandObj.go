package commands

import (
	"log"
	"os"
	"time"
)

// Command type for: [commandPrefix][command] <arguments> => <value>
//
// (...) are not always asked for.
type Command struct {
	// The Name of the command used in the chat with the bot.
	Name string

	// The Value returned from the bot to the user.
	Value func(arg ...any) string

	// The Description of the command.
	Description string

	// To determine the type of the command's parameters.
	ParamsType commandsParametersType

	// To determine the type of the command.
	CommandType commandType
}

// Log logs the command to indicate when its executed. This is used with error tracking.
func (cmd *Command) Log(s string, logger *log.Logger) {
	if cmd.Name == Closecmd.Name {
		logger.Println("Alex offline.")
	}

	logger.Printf("%s Executed. %s\n", cmd.Name, s)
}

func (cmd *Command) Close(s string, logger *log.Logger) {

	if cmd.Name == Closecmd.Name {
		logger.Println("Alex offline.")
		time.Sleep(time.Second * 2)
		logger.Printf("%s Executed. %s\n", cmd.Name, s)
		os.Exit(0)
	}
}
