package main

import (
	"encodings-ex01/base85"
	"fmt"
)

func main() {
	fmt.Println(string(base85.Decode([]byte("9jqo^"))))
}
