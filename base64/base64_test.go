package base64

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tc := range []struct {
		str  []byte
		want []byte
	}{
		{[]byte{0b1010100, 0b1110111, 0b1100101, 0b1101100, 0b1110110, 0b1100101, 0b100000, 0b1111001, 0b1100101, 0b1100001, 0b1110010, 0b1110011, 0b100000, 0b1101100, 0b1100001, 0b1110100, 0b1100101, 0b1110010}, []byte{84, 119, 101, 108, 118, 101, 32, 121, 101, 97, 114, 115, 32, 108, 97, 116, 101, 114}},
	} {
		t.Run("", func(t *testing.T) {
			if got := Encode(tc.str); !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Encode(%v): got = %v, want %v", tc.str, got, tc.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range []struct {
		src  []byte
		want []byte
	}{
		{[]byte{84, 119, 101, 108, 118, 101, 32, 121, 101, 97, 114, 115, 32, 108, 97, 116, 101, 114}, []byte{0b1010100, 0b1110111, 0b1100101, 0b1101100, 0b1110110, 0b1100101, 0b100000, 0b1111001, 0b1100101, 0b1100001, 0b1110010, 0b1110011, 0b100000, 0b1101100, 0b1100001, 0b1110100, 0b1100101, 0b1110010}},
	} {
		t.Run("", func(t *testing.T) {
			if got := Decode(tc.src); !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Decode(%v) = %v, want %v", tc.src, got, tc.want)
			}
		})
	}
}
