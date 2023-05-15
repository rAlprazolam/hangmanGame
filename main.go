package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

/*
 +----+
 0    |
/|\   |
/ \   |
     ===

Secret word : M_____
Incorret Guesses : A

Guess a letter : Y

Sorry Your Dead! The word was MONKEY
Yes the secret word is MONKEY

Please Enter Only One Letter
Please Enter a Letter
Please Enter a Letter You Haven´t Guessed

*/

var hangmanArr = [7]string{
	" +----+\n" +
		" 0    |\n" +
		"/|\\   |\n" +
		"/ \\   |\n" +
		"     ===\n",
	" +----+\n" +
		" 0    |\n" +
		"/|\\   |\n" +
		"/     |\n" +
		"     ===\n",
	" +----+\n" +
		" 0    |\n" +
		"/|\\   |\n" +
		"      |\n" +
		"     ===\n",
	" +----+\n" +
		" 0    |\n" +
		"/|    |\n" +
		"      |\n" +
		"     ===\n",
	" +----+\n" +
		" 0    |\n" +
		" |    |\n" +
		"      |\n" +
		"     ===\n",
	" +----+\n" +
		" 0    |\n" +
		"      |\n" +
		"      |\n" +
		"     ===\n",
	" +----+\n" +
		"      |\n" +
		"      |\n" +
		"      |\n" +
		"     ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZODIAC", "ZIPPER", "FLUFF", "ZOMBIE", "APPLE", "TREE",
}

var randWord string
var guessedLetters string
var correctLetter []string
var wrongGuesses []string

func getLetterInput() rune {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	rInputSlice := []rune(input)
	if len(rInputSlice) > 1 {
		fmt.Printf("Please Enter Only One Letter: ")
		return getLetterInput()
	}
	if !unicode.IsLetter(rInputSlice[0]) {
		fmt.Printf("Please Enter a Letter: ")
		return getLetterInput()
	}
	return rInputSlice[0]
}

func main() {
	var lives int = 6
	// Show Game Board

	// Get a letter from the user

	// A. If they guesses letter in word
	// Add to correctLetter
	// 1. Are there more letters to guess?
	// 2. If no more letters to guess (You Win)
	// B. If they guessed letter not in word
	// 1. Add new letter to guessedLetters,
	// wrongGuesses
	// Check if they died

	for lives > 0 {
		fmt.Println(hangmanArr[lives])
		fmt.Printf("Please Guess a Letter: ")
		letter := getLetterInput()
	}
}
