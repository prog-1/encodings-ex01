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
		{[]byte{0b1010001, 0b1010111, 0b1001010, 0b1101010, 0b1011010, 0b1000001, 0b111101, 0b111101}, []byte{85, 86, 100, 75, 97, 108, 112, 66, 80, 84, 48, 61}},
	} {
		got := Encode(tc.str)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERR: Encode(%v): got = %v, want = %v", tc.str, got, tc.want)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range []struct {
		str  []byte
		want []byte
	}{
		{[]byte{85, 86, 100, 75, 97, 108, 112, 66, 80, 84, 48, 61}, []byte{0b1010001, 0b1010111, 0b1001010, 0b1101010, 0b1011010, 0b1000001, 0b111101, 0b111101}},
	} {
		got := Decode(tc.str)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERR: Decode(%v): got = %v, want = %v", tc.str, got, tc.want)
		}
	}
}
