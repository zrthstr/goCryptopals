/*
 * Fixed XOR
 * Write a function that takes two equal-length buffers and produces
 * their XOR combination.
 *
 * If your function works properly, then when you feed it the string:
 *
 * 1c0111001f010100061a024b53535009181c
 * ... after hex decoding, and when XOR'd against:
 *
 * 686974207468652062756c6c277320657965
 * ... should produce:
 *
 * 746865206b696420646f6e277420706c6179
 *
 */

package main

import (
	"encoding/hex"
	"errors"
	"fmt"
)

func sXOR(a []byte, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("Byte stices not same lenght...")
	}
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c, nil
}

func main() {
	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	cShould := "746865206b696420646f6e277420706c6179"

	c, _ := sXOR(a, b)

	fmt.Println("cIs:", hex.EncodeToString(c))
	fmt.Println("cShould:", cShould)
}
