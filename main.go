package main

import (
	"bytes"
	"encodings-ex01/base64"
	"fmt"
)

func main() {

	matched := true
	for _, input := range [][]byte{
		nil,
		{0}, {63}, {'A'}, {255},
		{0, 0}, {0, 255}, {'a', 'A'}, {255, 255},
		{0, 0, 0}, {0, 0, 255}, {0, 255, 255}, {0, 'a', 'A'}, {'a', 'b', 'Z'}, {255, 255, 0}, {255, 255, 255},
		[]byte("Many hands make light work."), // Example from https://en.wikipedia.org/wiki/Base64#Examples
	} {
		encoded := base64.Encode(input)
		decoded := base64.Decode(encoded)

		if bytes.Equal(input, decoded) {
			continue
		}

		matched = false
		fmt.Println("FAILED ", input)
		fmt.Print("input  : ", input)

		fmt.Println()
		fmt.Print("encoded: ", encoded)

		fmt.Println()
		fmt.Printf("decoded: %d\n", decoded)

	}
	if matched {
		fmt.Println("ALL MATCHED")
	}
}
