package base64

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"1", []byte{0}, []byte{65, 65}},   // 0000_0000
		{"2", []byte{1}, []byte{65, 81}},   // 0000_0001
		{"3", []byte("1"), []byte{77, 81}}, // 49 in ASCII = 0011_0001
		{"4", []byte{2}, []byte{65, 103}},
		{"5", []byte{3}, []byte{}},
		{"6", []byte{4}, []byte{66, 65}}, // 0000_0100
		{"7", []byte{63}, []byte{80, 119}},
		{"8", []byte{'A'}, []byte{81, 81}}, // 65 in ASCII = 0100_0001
		{"9", []byte("A"), []byte{81, 81}},
		{"10", []byte{'B'}, []byte{81, 103}}, // 66 in ASCII = 0100_0010
		{"11", []byte("B"), []byte{81, 103}},
		{"12", []byte{255}, []byte{}},
		{"13", []byte{0, 0}, []byte{}},
		{"14", []byte{0, 255}, []byte{}},
		{"15", []byte{'a', 'A'}, []byte{}},
		{"16", []byte{255, 255}, []byte{}},
		{"17", []byte{0, 0, 0}, []byte{}},
		{"18", []byte{0, 0, 255}, []byte{}},
		{"19", []byte{0, 'a', 'A'}, []byte{}},
		{"20", []byte("Hello"), []byte{}},
		{"21", []byte("Many hands make light work."), []byte{}},
		{"22", []byte{77, 97, 110, 121, 32, 104, 97, 110, 100, 115, 32, 109, 97, 107, 101, 32, 108, 105, 103, 104, 116, 32, 119, 111, 114, 107, 46},
			[]byte{84, 87, 70, 117, 101, 83, 66, 111, 89, 87, 53, 107, 99, 121, 66, 116, 89, 87, 116, 108, 73, 71, 120, 112, 90, 50, 104, 48, 73, 72, 100, 118, 99, 109, 115, 117}},
	} {
		if got := Encode(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERROR: Encode(%v): got = %v want = %v", tc.name, got, tc.want)
		}
	}
}
