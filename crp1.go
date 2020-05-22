package main

/*
 * Convert hex to base64
 * The string:
 *
 * 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d
 * Should produce:
 *
 * SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
 * So go ahead and make that happen. You'll need to use this code for the rest of the exercises.
 *
 * Cryptopals Rule
 * Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
 *
 */

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HToB64(hs string) (string, error) {
	val, err := hex.DecodeString(hs)
	if err != nil {
		return "", err
	}
	fmt.Printf("%s", val)
	return base64.StdEncoding.EncodeToString(val), nil
}

func main() {
	hexStr := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	b64, _ := HToB64(hexStr)
	fmt.Println("")
	fmt.Println("hexStr:", hexStr)
	fmt.Println("b64:", b64)
}
