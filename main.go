package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
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
var guessedLetters []string
var correctLetters []string
var wrongGuesses []string

func getLetterInput() string {
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
	for _, value := range guessedLetters {
		if strings.ToUpper(input) == value {
			fmt.Print("Please Enter a Letter You Haven`t Guessed: ")
			return getLetterInput()
		}
	}
	guessedLetters = append(guessedLetters, input)
	return strings.ToUpper(input)
}

func getRandomWord() string {
	seed := time.Now().Unix()
	rand.Seed(seed)
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	fmt.Println(hangmanArr[6-len(wrongGuesses)])
	fmt.Print("Secret Word : ")
	for _, value := range correctLetters {
		if value == "" {
			fmt.Print("_")
		} else {
			fmt.Print(value)
		}
	}
	fmt.Print("\nIncorrect Guesses : ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	fmt.Println()
}

func getAllIndexes(theStr, subStr string) (indices []int) {
	if (len(subStr) == 0) || (len(theStr) == 0) {
		return indices
	}
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

func main() {
	getRandomWord()
	for true {
		showBoard()
		guess := getLetterInput()
		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
		}
	}
}
