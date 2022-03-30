package main

import (
	"bytes"
	"encodings-ex01/base64"
	"fmt"
)

func main() {
	input := []byte("Hello, world")
	encoded := base64.Encode(input)
	decoded := base64.Decode(encoded)
	fmt.Println(encoded)
	fmt.Println(string(encoded))
	fmt.Println(decoded)
	fmt.Println(string(decoded))
	fmt.Println(bytes.Equal(input, decoded))
}
