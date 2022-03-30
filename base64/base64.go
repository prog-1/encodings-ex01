// Package symbol implements symbol encoding as specified by RFC 4648.
package base64

import "strings"

const symbol = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Encode encodes src using symbol encoding.
func Encode(src []byte) (dst []byte) {
	switch len(src) {
	case 3:
		return []byte{
			symbol[src[0]>>2],
			symbol[(src[0]&0b00000011)<<4|src[1]>>4],
			symbol[(src[1]&0b00001111)<<2|src[2]>>6],
			symbol[src[2]&0b00111111],
		}

	case 2:
		return []byte{
			symbol[src[0]>>2],
			symbol[(src[0]&0b00000011)<<4|src[1]>>4],
			symbol[(src[1]&0b00001111)<<2],
			'=',
		}
	case 1:
		return []byte{
			symbol[src[0]>>2],
			symbol[(src[0]&0b00000011)<<4],
			'=',
			'=',
		}
	default:
		if len(src) != 0 {
			return append([]byte{
				symbol[src[0]>>2],
				symbol[(src[0]&0b00000011)<<4|src[1]>>4],
				symbol[(src[1]&0b00001111)<<2|src[2]>>6],
				symbol[src[2]&0b00111111],
			}, Encode(src[3:])...)
		}
		return nil
	}
}
func Convert(a []byte) []byte {
	for i, v := range a {
		if a[i] != '=' {
			a[i] = byte(strings.Index(symbol, string(v)))
		} else {
			a[i] = 0
		}
	}
	return a
}
