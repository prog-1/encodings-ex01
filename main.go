package main

import (
	"encodings-ex01/base85"
	"fmt"
)

func main() {
	fmt.Println(string(base85.Encode([]byte("Man is distinguished, not only by his reason, but by this singular passion from other animals, which is a lust of the mind, that by a perseverance of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of any carnal pleasure."))))
}
