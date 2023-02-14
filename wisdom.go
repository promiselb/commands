package commands

type Wisdom struct {
	Sentence string `json:"sentence"`
	Teller   string `json:"teller"`
}

func (w Wisdom) Print() string {
	return w.Sentence + " -" + w.Teller
}

// Very clear.
func RandomWisdom() *Wisdom {
	return RandomValue(Wisdoms...)
}

// I use it on my personal machine.
func GetWisdoms() ([]*Wisdom, error) {
	return getJsonObjects[*Wisdom]("wisdoms.json")
}

var Wisdoms, _ = GetWisdoms()
