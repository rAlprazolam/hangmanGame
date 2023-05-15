package main

import "fmt"

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

func main() {
	fmt.Print(hangmanArr[0])
	fmt.Print(hangmanArr[1])
	fmt.Print(hangmanArr[2])
	fmt.Print(hangmanArr[3])
	fmt.Print(hangmanArr[4])
	fmt.Print(hangmanArr[5])
	fmt.Print(hangmanArr[6])
}
