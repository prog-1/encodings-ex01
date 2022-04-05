package main

import (
	"bytes"
	"encodings-ex01/base64"
	"encodings-ex01/base85"
	"fmt"
)

func main() {
	input := []byte("implement me")
	encodedBase64 := base64.Encode(input)
	decodedBase64 := base64.Decode(encodedBase64)
	encodedBase85 := base85.Encode(input)
	decodedBase85 := base85.Decode(encodedBase85)
	fmt.Println(bytes.Equal(input, decodedBase64))
	fmt.Println(bytes.Equal(input, decodedBase85))
}
