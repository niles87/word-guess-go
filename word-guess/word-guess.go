package wordguess

import (
	"fmt"
)

const (
	GUESSES = 5
)

var storedWords = [][]byte{
	{'p', 'a', 'c', 'k', 'a', 'g', 'e'},
	{'i', 'o', 't', 'a'},
	{'g', 'o', 'l', 'a', 'n', 'g'},
	{'s', 't', 'r', 'u', 'c', 't'},
	{'c', 'o', 'n', 's', 't', 'a', 'n', 't'},
}

type Letter struct {
	Letter    byte
	IsGuessed bool
}

type Word struct {
	Word []Letter
}

type Player struct {
	Guesses uint
}

func (letter *Letter) DisplayLetter() byte {
	if letter.IsGuessed {
		return letter.Letter
	}
	return '_'
}

func (letter *Letter) UserGuess(char byte) bool {
	if char == letter.Letter {
		letter.IsGuessed = true
		return true
	}
	return false
}

func (word *Word) ToString() {
	for _, x := range word.Word {
		fmt.Print(x.DisplayLetter())
	}
	fmt.Println()
}

func (word *Word) CheckGuess(char byte) bool {
	var guess bool
	for _, x := range word.Word {
		guess = x.UserGuess(char)
	}
	return guess
}

func (player *Player) RemoveGuess(guess byte, word Word) {
	if !word.CheckGuess(guess) && player.Guesses > 1 {
		player.Guesses -= 1
	}
}

func (player *Player) GuessesRemaining() uint {
	return player.Guesses
}

func init() {

	fmt.Println("Welcome to Word Guess")
}
