// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

var extraLen = [3]int{0: 0, 1: 2, 2: 3}

// Encode encodes src using base64 encoding.
func Encode(src []byte) (dst []byte) {
	if len(src) == 0 {
		return nil
	}
	var pos int
	dst = make([]byte, len(src)/3*4+extraLen[len(src)%3])
	for ; len(src) >= 3; src = src[3:] {
		dst[pos] = base64[src[0]>>2]
		dst[pos+1] = base64[(src[0]&0b00000011)<<4|src[1]>>4]
		dst[pos+2] = base64[(src[1]&0b00001111)<<2|src[2]>>6]
		dst[pos+3] = base64[src[2]&0b00111111]
		pos += 4
	}
	switch len(src) {
	case 1:
		dst[pos] = base64[src[0]>>2]
		dst[pos+1] = base64[(src[0]&0b00000011)<<4]
	case 2:
		dst[pos] = base64[src[0]>>2]
		dst[pos+1] = base64[(src[0]&0b00000011)<<4|src[1]>>4]
		dst[pos+2] = base64[(src[1]&0b00001111)<<2]
	}
	return dst
}

// Decode decodes base64 encoded src.
func Decode(src []byte) (dst []byte) {
	return nil
}
