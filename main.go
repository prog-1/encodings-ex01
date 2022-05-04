package main

import (
	"bytes"
	"encodings-ex01/base64"
	"fmt"
)

func main() {
	for _, input := range [][]byte{
		nil,
		{0}, {63}, {'A'}, {255},
		{0, 0}, {0, 255}, {'a', 'A'}, {255, 255},
		{0, 0, 0}, {0, 0, 255}, {0, 255, 255}, {0, 'a', 'A'}, {'a', 'b', 'Z'}, {255, 255, 0}, {255, 255, 255},
		[]byte("Many hands make light work."), // Example from https://en.wikipedia.org/wiki/Base64#Examples
	} {
		encoded := base64.Encode(input)
		decoded := base64.Decode(encoded)
		fmt.Println(bytes.Equal(input, decoded))
	}
}
