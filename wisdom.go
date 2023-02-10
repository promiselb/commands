package commands

import "math/rand"

type Wisdom struct {
	Sentence string
	Teller   string
}

func RandomWisdom() *Wisdom {
	i := rand.Intn(len(Wisdoms))
	return Wisdoms[i]
}

func (w Wisdom) Print() string {
	return w.Sentence + " -" + w.Teller
}

var Wisdoms = []*Wisdom{
	{"The fool does think he is wise, but the wise man knows himself to be a fool.", "William Shakespeare"},
	{"Knowing yourself is the beginning of all wisdom.", "Aristotle"},
	{"The only true wisdom is in knowing you know nothing.", "Socrates"},
	{"Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.", "Rumi"},
	{"Show all your love on your friend, but do not show all your trust on him.", "Imam Ali"},
	{"Whoever touches his heart the exaggerated love for life, his heart will be surrounded by three things: worry that he does not lose, a desire that he does not leave, and hope that he does not realize.", "Imam Ali"},
	{"Too much laughter kills the heart.", "Prophet Mohamad"},
	{"Speak only when your words are more beneficial than your silence.", "Imam Ali"},
}
