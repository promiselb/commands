package commands

import (
	"log"
	"strings"
)

type Executer struct {
	commandPrefix string
	logger        *log.Logger
}

// NewExecuter returns a new Executer that executes the messages
// passed only prefix with prefix and use the logger to log them.
func NewExecuter(commandPrefix string, logger *log.Logger) *Executer {
	return &Executer{
		commandPrefix: commandPrefix,
		logger:        logger,
	}
}

// Execute execute takes a string as input and slices it so the first
// part is the command name and the 2nd part is the parameters passed.
func (executer *Executer) Execute(msg string, a ...any) string {
	if !strings.HasPrefix(msg, executer.commandPrefix) {
		return errMissingPrefix.Error()
	}
	msg = msg[1:] // remove the prefix

	if msg == "" {
		return errArg.Error()
	}

	commandName, arguments, _ := strings.Cut(msg, " ")

	if commandName == Helpcmd.Name {
		if arguments != "" {
			if cmd, ok := Commands[arguments[1:]]; ok {
				return cmd.Description
			} else {
				return errNoSuchCommand.Error()
			}
		}
		return Helpcmd.Handler("")
	}

	if cmd, ok := Commands[commandName]; ok {
		if cmd.ParamsType == MustParams && arguments == "" {
			return errArg.Error()
		}
		executer.Log(cmd, arguments)
		return cmd.Handler(arguments, a)
	} else {
		return errInvalidCommand.Error()
	}
}

// Log logs the command with the given message to indicate when its executed. This is used with error tracking.
func (executer *Executer) Log(cmd *Command, msg string) {
	executer.logger.Printf("%s Executed. %s\n", cmd.Name, msg)
}
