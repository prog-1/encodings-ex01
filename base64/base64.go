package base64

var (
	base64       = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	base64Decode = [256]uint32{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 62, 0, 0, 0, 63,
		52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 0, 0, 0, 0, 0,
		0, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
		41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 0, 0, 0, 0, 0,
	}
)

// Encode encodes src using base64 encoding.
func Encode(src []byte) []byte {
	var dst string
	if len(src) == 0 {
		return nil
	}
	for ; len(src) >= 3; src = src[3:] {
		x := uint32(src[0])<<16 | uint32(src[1])<<8 | uint32(src[2])
		dst += string([]byte{
			base64[x>>18],
			base64[x>>12&0x3f],
			base64[x>>6&0x3f],
			base64[x&0x3f],
		})
	}
	switch len(src) {
	case 1:
		dst += string([]byte{
			base64[src[0]>>2],
			base64[(src[0]&0x03)<<4],
		})
	case 2:
		x := uint16(src[0])<<8 | uint16(src[1])
		dst += string([]byte{
			base64[x>>10],
			base64[x>>4&0x3f],
			base64[(x&0x0f)<<2],
		})
	}
	return []byte(dst)
}

// Decode decodes base64 encoded src.
func Decode(src []byte) []byte {
	var dst string
	if len(src) == 0 {
		return nil
	}
	for ; len(src) >= 4; src = src[4:] {
		x := base64Decode[src[0]]<<18 | base64Decode[src[1]]<<12 | base64Decode[src[2]]<<6 | base64Decode[src[3]]
		dst += string([]byte{
			byte(x >> 16),
			byte(x >> 8),
			byte(x),
		})
	}
	switch len(src) {
	case 2:
		x := base64Decode[src[0]]<<2 | base64Decode[src[1]]>>4
		dst += string([]byte{
			byte(x),
		})
	case 3:
		x := base64Decode[src[0]]<<10 | base64Decode[src[1]]<<4 | base64Decode[src[2]]>>2
		dst += string([]byte{
			byte(x >> 8),
			byte(x),
		})
	}
	return []byte(dst)
}
