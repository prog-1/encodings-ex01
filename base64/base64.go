// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

var base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var base64Decode = [256]uint32{
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 62, 0, 0, 0, 63,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 0, 0, 0, 0, 0, 0,
	0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0, 0, 0, 0, 0,
	0, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 0, 0, 0, 0, 0,
}

var extraLen = [3]int{0: 0, 1: 2, 2: 3}
var extraLenDecode = [4]int{0: 0, 2: 1, 3: 2}

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
	if len(src) == 0 {
		return nil
	}
	var pos int
	dst = make([]byte, len(src)/4*3+extraLenDecode[len(src)%4])
	for ; len(src) >= 4; src = src[4:] {
		dst[pos] = byte(base64Decode[src[0]]<<2 | base64Decode[src[1]]>>4)
		dst[pos+1] = byte(base64Decode[src[1]]<<4 | base64Decode[src[2]]>>2)
		dst[pos+2] = byte(base64Decode[src[2]]<<6 | base64Decode[src[3]])
		pos += 3
	}
	switch len(src) {
	case 1:
		dst[pos] = byte(base64Decode[src[0]]<<2 | base64Decode[src[1]]>>4)
	case 2:
		dst[pos] = byte(base64Decode[src[0]]<<2 | base64Decode[src[1]]>>4)
		dst[pos+1] = byte(base64Decode[src[1]]<<4 | base64Decode[src[2]]>>2)
	}
	return dst
}
