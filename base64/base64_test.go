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
		{"5", []byte{3}, []byte{65, 119}},
		{"6", []byte{4}, []byte{66, 65}}, // 0000_0100
		{"7", []byte{63}, []byte{80, 119}},
		{"8", []byte{'A'}, []byte{81, 81}}, // 65 in ASCII = 0100_0001
		{"9", []byte("A"), []byte{81, 81}},
		{"10", []byte{'B'}, []byte{81, 103}}, // 66 in ASCII = 0100_0010
		{"11", []byte("B"), []byte{81, 103}},
		{"12", []byte{255}, []byte{47, 119}},
		{"13", []byte{0, 0}, []byte{65, 65, 65}},
		{"14", []byte{0, 255}, []byte{65, 80, 56}},
		{"15", []byte{'a', 'A'}, []byte{89, 85, 69}},
		{"16", []byte{255, 255}, []byte{47, 47, 56}},
		{"17", []byte{0, 0, 0}, []byte{65, 65, 65, 65}},
		{"18", []byte{0, 0, 255}, []byte{65, 65, 68, 47}},
		{"19", []byte{0, 'a', 'A'}, []byte{65, 71, 70, 66}},
		{"20", []byte("Hello"), []byte{83, 71, 86, 115, 98, 71, 56}},
		{"21", []byte("Many hands make light work."),
			[]byte{84, 87, 70, 117, 101, 83, 66, 111, 89, 87, 53, 107, 99, 121, 66, 116, 89, 87, 116, 108, 73, 71, 120, 112, 90, 50, 104, 48, 73, 72, 100, 118, 99, 109, 115, 117}},
	} {
		if got := Encode(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERROR: Encode(%v): got = %v want = %v", tc.name, got, tc.want)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"1", []byte{65, 65}, []byte{0}},
		{"2", []byte{65, 81}, []byte{1}},
		{"3", []byte{77, 81}, []byte("1")},
		{"4", []byte{65, 103}, []byte{2}},
		{"5", []byte{65, 119}, []byte{3}},
		{"6", []byte{66, 65}, []byte{4}},
		{"7", []byte{80, 119}, []byte{63}},
		{"8", []byte{81, 81}, []byte{'A'}},
		{"9", []byte{81, 81}, []byte("A")},
		{"10", []byte{81, 103}, []byte{'B'}},
		{"11", []byte{81, 103}, []byte("B")},
		{"12", []byte{47, 119}, []byte{255}},
		{"13", []byte{65, 65, 65}, []byte{0, 0}},
		{"14", []byte{65, 80, 56}, []byte{0, 255}},
		{"15", []byte{89, 85, 69}, []byte{'a', 'A'}},
		{"16", []byte{47, 47, 56}, []byte{255, 255}},
		{"17", []byte{65, 65, 65, 65}, []byte{0, 0, 0}},
		{"18", []byte{65, 65, 68, 47}, []byte{0, 0, 255}},
		{"19", []byte{65, 71, 70, 66}, []byte{0, 'a', 'A'}},
		{"20", []byte{83, 71, 86, 115, 98, 71, 56}, []byte("Hello")},
		{"21", []byte{84, 87, 70, 117, 101, 83, 66, 111, 89, 87, 53, 107, 99, 121, 66, 116, 89, 87, 116, 108, 73, 71, 120, 112, 90, 50, 104, 48, 73, 72, 100, 118, 99, 109, 115, 117},
			[]byte("Many hands make light work.")},
	} {
		if got := Decode(tc.input); !reflect.DeepEqual(got, tc.want) {
			t.Errorf("ERROR: Decode(%v): got = %v want = %v", tc.name, got, tc.want)
		}
	}
}
