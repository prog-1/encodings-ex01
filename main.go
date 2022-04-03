package main

import (
	"encodings-ex01/base85"
	"fmt"
)

func main() {
	input := []byte("distinguished")
	encoded := base85.Encode(input)
	//decoded := base85.Decode(encoded)
	fmt.Println(string(encoded))
}
