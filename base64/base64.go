// Package base64 implements base64 encoding as specified by RFC 4648.
package base64

const base64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// Encode encodes src using base64 encoding.
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

// Decode decodes base64 encoded src.
func Decode(src []byte) (dst []byte) {
	return nil
}
