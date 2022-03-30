package main

import (
	"bytes"
	"encodings-ex01/base64"
	"fmt"
)

func main() {
	input := []byte("Abcbdjalk")
	encoded := base64.Encode(input)
	decoded := base64.Decode(encoded)
	fmt.Println(bytes.Equal(input, decoded))
}
