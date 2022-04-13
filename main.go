package main

import (
	"bytes"
	"encodings-ex01/base64"
	"fmt"
)

func main() {
	input := []byte("implement me")
	encodedBase64 := base64.Encode(input)
	decodedBase64 := base64.Decode(encodedBase64)
	fmt.Println(bytes.Equal(input, decodedBase64))
}
