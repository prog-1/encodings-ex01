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
		{[]byte{0b1101001, 0b1101101, 0b1110000, 0b1101100, 0b1100101, 0b1101101, 0b1100101, 0b1101110, 0b1110100, 0b100000, 0b1101101, 0b1100101}, []byte{97, 87, 49, 119, 98, 71, 86, 116, 90, 87, 53, 48, 73, 71, 49, 108}},
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
		{[]byte{97, 87, 49, 119, 98, 71, 86, 116, 90, 87, 53, 48, 73, 71, 49, 108}, []byte{0b1101001, 0b1101101, 0b1110000, 0b1101100, 0b1100101, 0b1101101, 0b1100101, 0b1101110, 0b1110100, 0b100000, 0b1101101, 0b1100101}},
	} {
		t.Run("", func(t *testing.T) {
			if got := Decode(tc.src); !reflect.DeepEqual(tc.want, got) {
				t.Errorf("Decode(%v) = %v, want %v", tc.src, got, tc.want)
			}
		})
	}
}
