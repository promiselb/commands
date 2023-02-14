package commands

// Very clear.
func RandomProverb() *Wisdom {
	return RandomValue(Proverbs...)
}

// I use it on my personal machine.
func GetProverbs() ([]*Wisdom, error) {
	return getJsonObjects[*Wisdom]("proverbs.json")
}

var Proverbs, _ = GetProverbs()
