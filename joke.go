package commands

type Joke struct {
	SingleJoke []string `json:"singleJoke"`
}

func (j Joke) Print() string {
	s := ""
	for _, v := range j.SingleJoke {
		s += v + "\n"
	}
	return s
}

// Very clear.
func RandomJoke() *Joke {
	return RandomValue(Jokes...)
}

// I use it on my personal machine.
func GetJokes() ([]*Joke, error) {
	return getJsonObjects[*Joke]("jokes.json")
}

var Jokes, _ = GetJokes()
