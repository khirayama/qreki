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
		in, _ := time.Parse("2006-01-02", tc.in)
		got := NewQreki(in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}
