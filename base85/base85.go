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
			} else if len(src) == 3 {
				src = src[3:]
			} else if len(src) == 2 {
				src = src[2:]
			} else if len(src) == 1 {
				src = src[1:]
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
