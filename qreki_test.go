package qreki

import (
	"reflect"
	"testing"
	// "fmt"
	"time"
)

func TestQreki(t *testing.T) {
	testcases := []struct {
		in  string
		out Qreki
	}{
		{"2015-12-31", Qreki{2015, 11, 21, false, "大安", 0.0, 0.0, 0.0, 0}},
	}

	for _, tc := range testcases {
		in, _ := time.Parse("2002-06-03", tc.in)
		got := NewQreki(ToJulian(in))
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func TestToJulian(t *testing.T) {
	testcases := []struct {
		in  time.Time
		out Julian
	}{
		{time.Date(2002, 6, 3, 0, 0, 0, 0, JST), Julian(2452428)},
	}

	for _, tc := range testcases {
		got := ToJulian(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}
