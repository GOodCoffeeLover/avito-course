package guesses

import (
	"log"
	"math/rand"
)

var guessesClient *Guesses

type Guesses struct {
	number int
}

func New() *Guesses {
	if guessesClient != nil {
		return guessesClient
	}

	guessesClient = &Guesses{
		number: rand.Intn(100),
	}
	log.Printf("init guesses with value %v", guessesClient.number)
	return guessesClient
}

func (g *Guesses) Reset() {
	if g == nil {
		panic("using nil guesses object")
	}
	g.number = rand.Intn(100)
	log.Printf("reset guesses with value %v", guessesClient.number)

}

func (g *Guesses) Guess(guess int) int {
	if g == nil {
		panic("using nil guesses object")
	}
	if g.number < guess {
		return 1
	}

	if g.number > guess {
		return -1
	}

	return 0
}
