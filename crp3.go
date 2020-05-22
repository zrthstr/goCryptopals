package main

/*
 *
 * Single-byte XOR cipher
 * The hex encoded string:
 *
 * 1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
 * ... has been XOR'd against a single character. Find the key, decrypt the message.
 *
 * You can do this by hand. But don't: write code to do it for you.
 *
 * How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.
 *
 */

import (
	"fmt"
)

func score(str string) int {
	// char freq in percent * 1000, english
	//freq := map[string]int{
	freq := map[byte]int{
		'e': 11162,
		't': 9356,
		'a': 8497,
		'r': 7587,
		'i': 7546,
		'o': 7507,
		'n': 6749,
		's': 6327,
		'h': 6094,
		'd': 4253,
		'l': 4025,
		'u': 2758,
		'w': 2560,
		'm': 2406,
		'f': 2228,
	}
	fmt.Println(freq)
	return len(str) // div by score
}

func main() {
	fmt.Println("..")
}
