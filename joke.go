package commands

import "math/rand"

type Joke []string

func RandomJoke() *Joke {
	i := rand.Intn(len(Jokes))
	return Jokes[i]
}

var Jokes = []*Joke{
	{"Why was 6 afraid of 7?", "Because 7 ate 9!"},
	{"I'm reading an antigravity book.", "It's impossible to put down!"},
	{"What kind of cheese doesn't belong to you?", "Nacho cheese!"},
	{"You can't trust atoms.", "They make up everything!"},
	{"Why do mushrooms get invited to all the best parties?", "Because they are such fungis!"},
	{"Why did the dog cross the road?", "To get to the barking lot!"},
	{"Why do potatoes argue?", "Because they can't see eye to eye!"},
	{"Can February March?", "No, but April May!"},
	{"What's the loneliest cheese?", "ProvAlone!"},
	{"What kind of dog does Dracula have?", "A bloodhound!"},
	{"Did you get a haircut?", "No, I got them all cut!"},
	{"What do you call a cold dog?", "A chili dog!"},
	{"Why did the cat run away from the tree?", "It was afraid of the bark!"},
	{"Why don't cannibals eat clowns?", "Because they taste funny!"},
	{"Why don't eggs tell each other jokes?", "They'd crack each other up!"},
	{"How do you fix a broken tomato?", "Tomato paste!"},
	{"Why did the hipster burn his tongue on coffee?", "Because he drank it before it was cool!"},
	{"Why did the farmer win an award?", "He was out standing in his field!"},
	{"What do sea monsters eat?", "Fish and ships!"},
	{"Why do melons have weddings?", "Because they cantaloupe!"},
	{"What did the daddy tomato say to the baby tomato?", "Ketchup!"},
	{"What did the ocean say to the shore?", "Nothing, it just waved!"},
	{"Why do birds fly south?", "It's easier than walking!"},
	{"What's a plant's favorite drink?", "Root beer!"},
	{"Do you think...", "Earth makes fun of other planets for having no life?!"},
	{"Who cleans the ocean?", "Mer-maids!"},
	{"Velcro...", "Is such a rip-off!"},
	{"What kind of songs do tortillas write?", "Wraps!"},
	{"Why don't people like Russian dolls?"},
	{"Because they are full of themselves!"},
	{"What did the hungry clock do?", "It went back 4 seconds!"},
	{"What's the best season for trampolines?", "Spring time!"},
	{"Why did the two 4s skip dinner?", "They already 8!"},
	{"I used to hate facial hair."},
	{"But then it grew on me!"},
	{"Who shaves 10 times a day but still has a beard?", "A barber!"},
	{"Why did the students eat their homework?", "Because the teacher said it was a piece of cake!"},
	{"Why can't you trust a burrito?", "Because they tend to spill the beans!"},
	{"What do bees brush their hair with?", "A honeycomb!"},
	{"Where do snowmen keep their money?", "In snowbanks!"},
	{"Why is a carrot the best detective?", "They get to the root of every case!"},
	{"What do you call a meditating wolf?", "Aware wolf!"},
	{"What do you call an American bee?", "A USB!"},
	{"The past, present, and future walk into a war.", "It was tense!"},
	{"Not to brag,", "but I defeated our local chess champion in less than five moves.", "Finally,", "my high school karate lessons paid off."},
	{"Air used to be free at the gas station, now it's $1.50. You know why?", "Inflation."},
	{"I spent a lot of time, money, and effort childproofing my house...", "but the kids still get in."},
	{"When I was a kid,", "my mother told me I could be anyone I wanted to be.", "Turns out,", "identity theft is a crime."},
}
