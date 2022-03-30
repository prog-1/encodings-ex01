// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

import (
	"strings"
)

const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func encodeSwitch(src []byte) (dst []byte) {
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
		return nil
	}
}

// Encode encodes src using base64 encoding.
// SGVsbG8sIHdvcmxk
func Encode(s []byte) (dst []byte) {
	var src []byte
	for src = s; len(src) > 3; src = src[3:] {
		dst = append(dst,
			base64[src[0]>>2],
			base64[(src[0]&0b00000011)<<4|src[1]>>4],
			base64[(src[1]&0b00001111)<<2|src[2]>>6],
			base64[src[2]&0b00111111],
		)
	}
	return append(dst, encodeSwitch(src)...)
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

func decodeSwitch(src []byte) []byte {
	switch len(src) {
	case 4:
		return filter([]byte{
			src[0]<<2 | (src[1] >> 4),
			(src[1] << 4) | (src[2] >> 2),
			(src[2] << 6) | src[3],
		})
	case 3:
		return filter([]byte{
			src[0]<<2 | (src[1] >> 4),
			(src[1] << 4) | (src[2] >> 2),
			(src[2] << 6),
		})
	case 2:
		return filter([]byte{
			src[0]<<2 | (src[1] >> 4),
			(src[1] << 4),
		})
	case 1:
		return filter([]byte{
			src[0] << 2,
		})
	default:
		return nil
	}
}

// Decode decodes base64 encoded src.
func Decode(s []byte) (dst []byte) {
	Convert(s)
	var src []byte
	for src = s; len(src) > 4; src = src[4:] {
		dst = append(dst, filter([]byte{
			src[0]<<2 | (src[1] >> 4),
			(src[1] << 4) | (src[2] >> 2),
			(src[2] << 6) | src[3],
		})...)
	}
	return append(dst, decodeSwitch(src)...)
}
