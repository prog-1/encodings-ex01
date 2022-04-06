// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

import (
	"fmt"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var antialphabet = [256]uint32{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 62, 0, 0, 0, 63,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0, 0, 0, 0, 0,
	0, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 0, 0, 0, 0, 0,
}

// Encode encodes src using base64 encoding.
func Encode(src []byte) (dst []byte) {
	fmt.Println(string(dst))
	if len(src) == 0 {
		return nil
	}
	//for i := 1; len(src) > 0; {
	//switch i {
	//case 1:
	//dst = append(dst, alphabet[src[0]>>2])
	//i++
	//case 2:
	//if len(src) == 1 {
	//	dst = append(dst, alphabet[(src[0]&0b00000011)<<4])
	//	return
	//} else {
	//	dst = append(dst, alphabet[(src[0]&0b00000011)<<4|src[1]>>4])
	//	i++
	//}
	//case 3:
	//if len(src) == 2 {
	//dst = append(dst, alphabet[(src[1]&0b00001111)<<2])
	//return
	//} else {
	//	dst = append(dst, alphabet[(src[1]&0b00001111)<<2|src[2]>>6])
	//i++
	//}
	//case 4:
	//dst = append(dst, alphabet[src[2]&0b00111111])
	//i = 1
	//src = src[3:]
	//}
	//}

	for len(src) > 0 { //there is bug somewhere that eats last number
		if len(src) >= 3 {
			dst = append(dst, alphabet[src[0]>>2])
			dst = append(dst, alphabet[(src[0]&0b00000011)<<4|src[1]>>4])
			dst = append(dst, alphabet[(src[1]&0b00001111)<<2|src[2]>>6])
			dst = append(dst, alphabet[src[2]&0b00111111])
			src = src[3:]
		} else {
			switch len(src) {
			case 1:
				dst = append(dst, alphabet[src[0]>>2])
				dst = append(dst, alphabet[(src[0]&0b00000011)<<4])
				return dst
			case 2:
				dst = append(dst, alphabet[src[0]>>2])
				dst = append(dst, alphabet[(src[0]&0b00000011)<<4|src[1]>>4])
				dst = append(dst, alphabet[(src[1]&0b00001111)<<2])
				return dst
			}
		}
	}
	return dst
}

// Decode decodes base64 encoded src.
func Decode(src []byte) (dst []byte) {
	if len(src) == 0 {
		return nil
	}
	for len(src) > 0 {
		if len(src) >= 4 {
			x := antialphabet[src[0]]<<18 | antialphabet[src[1]]<<12 | antialphabet[src[2]]<<6 | antialphabet[src[3]]
			dst = append(dst, byte(x>>16))
			dst = append(dst, byte(x>>8))
			dst = append(dst, byte(x))
			src = src[4:]
		} else {
			switch len(src) {
			case 2:
				x := antialphabet[src[0]]<<2 | antialphabet[src[1]]>>4
				dst = append(dst, byte(x))
				return dst
			case 3:
				x := antialphabet[src[0]]<<10 | antialphabet[src[1]]<<4 | antialphabet[src[2]]>>2
				dst = append(dst, byte(x>>8))
				dst = append(dst, byte(x))
				return dst
			}
		}
	}
	return dst
}
