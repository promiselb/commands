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
func (executer *Executer) Execute(msg string, anyThing ...any) string {
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
		return Helpcmd.Value()
	}

	if cmd, ok := Commands[commandName]; ok {
		if cmd.ParamsType == MustParams && arguments == "" {
			return errArg.Error()
		}
		cmd.Log(arguments, executer.logger)
		return cmd.Value(arguments, anyThing)
	} else {
		return errInvalidCommand.Error()
	}
}
