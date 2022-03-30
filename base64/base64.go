// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

import "strings"

// Encode encodes src using base64 encoding.
const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func Encode(src []byte) (dst []byte) {
	switch len(src) {
	case 3:
		return []byte{
			base64[src[0]>>2],
			base64[(src[0]&0b00000011)<<4|src[1]>>4],
			base64[(src[1]&0b00001111)<<2|src[2]>>6],
			base64[src[2]&0b00111111],
		}

	case 2:
		return []byte{
			base64[src[0]>>2],
			base64[(src[0]&0b00000011)<<4|src[1]>>4],
			base64[(src[1]&0b00001111)<<2],
			'=',
		}
	case 1:
		return []byte{
			base64[src[0]>>2],
			base64[(src[0]&0b00000011)<<4],
			'=',
			'=',
		}
	default:
		if len(src) != 0 {
			return append([]byte{
				base64[src[0]>>2],
				base64[(src[0]&0b00000011)<<4|src[1]>>4],
				base64[(src[1]&0b00001111)<<2|src[2]>>6],
				base64[src[2]&0b00111111],
			}, Encode(src[3:])...)
		}
		return nil
	}
}

func Convert(a []byte) []byte {
	for i, v := range a {
		if a[i] != '=' {
			a[i] = byte(strings.Index(base64, string(v)))
		} else {
			a[i] = 0
		}
	}
	return a
}

func filter(a []byte) []byte {
	n := 0
	for _, x := range a {
		if x != 0 {
			a[n] = x
			n++
		}
	}
	a = a[:n]
	return a
}

func Decode(src []byte) []byte {
	return decode(Convert(src))
}

func decode(src []byte) (dst []byte) {
	switch len(src) {
	case 4:
		return filter([]byte{
			src[0]<<2 | (src[1]>>4)&0b00000011,
			((src[1] & 0b00001111) << 4) | (src[2] >> 2),
			((src[2] & 0b00000011) << 6) | src[3],
		})
	case 3:
		return filter([]byte{
			src[0]<<2 | (src[1]>>4)&0b00000011,
			((src[1] & 0b00001111) << 4) | (src[2] >> 2),
			((src[2] & 0b00000011) << 6),
		})
	case 2:
		return filter([]byte{
			src[0]<<2 | (src[1]>>4)&0b00000011,
			((src[1] & 0b00001111) << 4),
		})
	case 1:
		return filter([]byte{
			src[0] << 2,
		})
	default:
		if len(src) != 0 {
			return filter(
				append(
					[]byte{src[0]<<2 | (src[1]>>4)&0b00000011,
						((src[1] & 0b00001111) << 4) | (src[2] >> 2),
						((src[2] & 0b00000011) << 6) | src[3]},
					decode(src[4:])...))
		}
	}
	return
}
