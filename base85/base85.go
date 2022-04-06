package base85

func Encode(src []byte) (res []byte) {
	if len(src) == 0 {
		return nil
	}
	for len(src) > 0 {
		var v uint32
		switch len(src) {
		default:
			v |= uint32(src[3])
			fallthrough
		case 3:
			v |= uint32(src[2]) << 8
			fallthrough
		case 2:
			v |= uint32(src[1]) << 16
			fallthrough
		case 1:
			v |= uint32(src[0]) << 24
		}
		if v == 0 && len(src) >= 4 {
			res = append(res, 'z')
			if len(src) >= 4 {
				src = src[4:]
			} else {
				src = nil
			}
			continue
		}
		var tmp []byte
		for i := 4; i >= 0; i-- {
			tmp = append(tmp, '!'+byte(v%85))
			v /= 85
		}
		for i := len(tmp) - 1; i >= 0; i-- {
			res = append(res, tmp[i])
		}

		if len(src) >= 4 {
			src = src[4:]
		} else {
			src = nil
		}
	}
	return res
}
func Decode(src []byte) (res []byte) {

	if len(src) == 0 {
		return nil
	}
	var v uint32
	var nb int
	for _, b := range src {
		switch {
		case b <= ' ':
			continue
		case b == 'z' && nb == 0:
			nb = 5
			v = 0
		case '!' <= b && b <= 'u':
			v = v*85 + uint32(b-'!')
			nb++
		}
		if nb == 5 {
			res = append(res, byte(v>>24))
			res = append(res, byte(v>>16))
			res = append(res, byte(v>>8))
			res = append(res, byte(v))
			nb = 0
			v = 0
		}
	}
	return res
}
