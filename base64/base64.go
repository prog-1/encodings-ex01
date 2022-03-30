// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Encode encodes src using base64 encoding.
func Encode(src []byte) (dst []byte) {
	//for len(src)%3 != 0 {
	//src = append(src, 61)
	//}
	for a, i := 0, 1; a < len(src); a++ {
		switch i {
		case 1:
			dst = append(dst, alphabet[src[a]>>2])
			i++
		case 2:
			dst = append(dst, alphabet[(src[a-1]&0b00000011)<<4|src[a]>>4])
			i++
		case 3:
			dst = append(dst, alphabet[(src[a-1]&0b00001111)<<2|src[a]>>6])
			i++
		case 4:
			a--
			dst = append(dst, alphabet[src[a]&0b00111111])
			i = 1
		}
	}
	return dst
}

// Decode decodes base64 encoded src.
func Decode(src []byte) (dst []byte) {
	if len(src) < 1 {
		return nil
	}
	return dst
}
