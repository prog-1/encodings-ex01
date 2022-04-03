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

func Encode(src []byte) (dst []byte) {
	if len(src) == 0 {
		return nil
	}
	i := 0
	for ; i <= len(src)-3; i = i + 3 {
		tmp := src[i:]
		x := uint32(tmp[0])<<16 | uint32(tmp[1])<<8 | uint32(tmp[2])
		dst = append(dst, base64[x>>18])
		dst = append(dst, base64[(x>>12)&0x3f])
		dst = append(dst, base64[(x>>6)&0x3f])
		dst = append(dst, base64[x&0x3f])
	}
	switch len(src) % 3 {
	case 1:
		x := src[i]
		dst = append(dst, base64[x>>2])
		dst = append(dst, base64[(x&0x03)<<4])

	case 2:
		x := uint16(src[i])<<8 | uint16(src[i+1])
		dst = append(dst, base64[x>>10])
		dst = append(dst, base64[x>>4&0x3f])
		dst = append(dst, base64[(x&0x0f)<<2])

	}
	return dst
}
func Decode(src []byte) (dst []byte) {
	if len(src) == 0 {
		return nil
	}
	i := 0
	for ; i <= len(src)-4; i = i + 4 {
		tmp := src[i:]
		x := base64Decode[tmp[0]]<<18 | base64Decode[tmp[1]]<<12 | base64Decode[tmp[2]]<<6 | base64Decode[tmp[3]]
		dst = append(dst, byte(x>>16))
		dst = append(dst, byte(x>>8))
		dst = append(dst, byte(x))

	}

	if len(src) == 2 || len(src)%4 == 2 {
		x := base64Decode[src[i]]<<2 | base64Decode[src[i+1]]>>4
		dst = append(dst, byte(x))
	}

	if len(src) == 3 || len(src)%4 == 3 {
		x := base64Decode[src[i]]<<10 | base64Decode[src[i+1]]<<4 | base64Decode[src[i+2]]>>2
		dst = append(dst, byte(x>>8))
		dst = append(dst, byte(x))
	}
	return dst
}
