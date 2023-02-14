package commands

// Command type for: [commandPrefix][command] <arguments> => <value>
//
// (...) are not always asked for.
type Command struct {
	// The Name of the command used in the chat with the bot.
	Name string

	// The Handler returned from the bot to the user.
	Handler func(s string, arg ...any) string

	// The Description of the command.
	Description string

	// To determine the type of the command's parameters.
	ParamsType commandsParametersType
}
