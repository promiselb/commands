package commands

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
)

const (
	// a command that doesn't need parameters when to execute it.
	NoParams commandsParametersType = iota

	// a command that parameters are optional
	OptParams

	// a command that parameters are necessary.
	MustParams
)

type commandsParametersType int

// commandPrefix = "$", logger = log.Default()
var StdExecuter = NewExecuter("$", log.Default())

// Errors.
var (
	errMissingPrefix  = errors.New("error: Missing prefix?")
	errArg            = errors.New("error: Not enough arguments or missing key literal?")
	errInvalidCommand = errors.New("error: invalid command")
	errNoSuchCommand  = errors.New("error: No such command?")
)

var (
	Closecmd = &Command{ // 1
		Name: "close",
		Handler: func(s string, a ...any) string {
			return "Alex offline"
		},
		Description: "Returns Alex to sleep.",
	}

	Hicmd = &Command{ // 2
		Name: "hi",
		Handler: func(s string, a ...any) string {
			if s == "" {
				s = "there"
			}
			return "Hello " + s + "!"
		},
		Description: "Greeting.",
		ParamsType:  OptParams,
	}

	Breakfastcmd = &Command{ // 3
		Name: "breakfast",
		Handler: func(s string, a ...any) string {
			return "Let's eat some fried eggs 🍳🥚 and cheese 🧀!"
		},
		Description: "Breakfast of cheese and fried eggs!",
	}

	Lunchcmd = &Command{ // 4
		Name: "lunch",
		Handler: func(s string, a ...any) string {
			return "Let's eat pizza 🍕!"
		},
		Description: "Pizza time!",
	}

	Dinnercmd = &Command{ //5
		Name: "dinner",
		Handler: func(s string, a ...any) string {
			return "Let's eat some meat 🥩 and drink something 🍸!"
		},
		Description: "Dinner time!",
	}

	Partycmd = &Command{ // 6
		Name: "party",
		Handler: func(s string, a ...any) string {
			return "Let's enjoy some gâteau 🎂 and live our life!"
		},
		Description: "Party time!",
	}

	Pingcmd = &Command{ // 7
		Name: "ping",
		Handler: func(s string, a ...any) string {
			return "pong"
		},
		Description: "Ping pong!",
	}

	Pongcmd = &Command{ // 9
		Name: "pong",
		Handler: func(s string, a ...any) string {
			return "ping"
		},
		Description: "Ping pong!",
	}

	Shrugcmd = &Command{ // 10
		Name: "shrug",
		Handler: func(s string, a ...any) string {
			return `¯\_(ツ)_/¯`
		},
		Description: `Adds ¯\_(ツ)_/¯ to the end of the message.`,
	}

	Comparecmd = &Command{ // 11
		Name: "compare",
		Handler: func(s string, a ...any) string {

			slice := SplitArgs(s)
			if len(slice) != 2 {
				return errArg.Error()
			}
			return fmt.Sprint(slice[0] == slice[1])
		},
		Description: `Compare two arguments like: [prefix]compare <arg[0]> <arg[1]>`,
		ParamsType:  MustParams,
	}

	Testcmd = &Command{ // 12
		Name: "test",
		Handler: func(s string, a ...any) string {
			return fmt.Sprint(RandomBool())
		},
		Description: "$test <sentence>?",
		ParamsType:  OptParams,
	}

	Helpcmd = &Command{ // 13
		Name: "help",
		Handler: func(s string, a ...any) string {
			return Help()
		},
		Description: "Lists all commands and theirs descriptions.",
	}

	SetTimercmd = &Command{ // 14
		Name: "timer",
		Handler: func(s string, a ...any) string {
			slice := SplitArgs(s)
			if len(slice) < 3 {
				return errArg.Error()
			}

			duration, err := strconv.Atoi(slice[1][1 : len(slice[1])-1])
			if err != nil {
				return err.Error()
			}

			nTimes, err := strconv.Atoi(slice[2][1 : len(slice[2])-1])
			if err != nil {
				return err.Error()
			}

			secs := time.Duration(duration)
			Ticker := time.NewTicker(secs * time.Second)

			mychannel := make(chan bool)
			go func() {
				for {
					select {
					case <-mychannel:
						return

					case <-Ticker.C:
						fmt.Println(slice[0][1 : len(slice[0])-1])
					}
				}
			}()

			time.Sleep((secs*time.Duration(nTimes))*time.Second + 1)

			Ticker.Stop()

			mychannel <- true

			return "Ticker is turned off!"
		},
		Description: "set a timer to send a message when the time is done. Like:\n$timer <msg> <d> <n>\nmsg: the message to be sent each after each time the intervalle is finished.\nd: the intervalle in seconds.\nn: how many times to send the message.",
		ParamsType:  MustParams,
	}

	Jokecmd = &Command{ // 15
		Name: "joke",
		Handler: func(s string, a ...any) string {
			return RandomJoke().Print()
		},
		Description: "Says a joke.",
		ParamsType:  NoParams,
	}

	Wisdomcmd = &Command{ // 16
		Name: "wisdom",
		Handler: func(s string, arg ...any) string {
			return RandomWisdom().Print()
		},
		Description: "Says a wisdom.",
		ParamsType:  NoParams,
	}

	Proverbcmd = &Command{ // 17
		Name: "proverb",
		Handler: func(s string, arg ...any) string {
			return RandomProverb().Print()
		},
		Description: "Says a golang proverb.",
		ParamsType:  NoParams,
	}
)

// Commands stored in a map by their names except for help command.
var Commands = map[string]*Command{
	Closecmd.Name: Closecmd,

	Hicmd.Name: Hicmd,

	Breakfastcmd.Name: Breakfastcmd,

	Lunchcmd.Name: Lunchcmd,

	Dinnercmd.Name: Dinnercmd,

	Partycmd.Name: Partycmd,

	Shrugcmd.Name: Shrugcmd,

	Pingcmd.Name: Pingcmd,

	Pongcmd.Name: Pongcmd,

	Testcmd.Name: Testcmd,

	Comparecmd.Name: Comparecmd,

	SetTimercmd.Name: SetTimercmd,

	Jokecmd.Name: Jokecmd,

	Wisdomcmd.Name: Wisdomcmd,

	Proverbcmd.Name: Proverbcmd,
}
