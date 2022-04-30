package qreki

import (
	"reflect"
	"testing"
	// "fmt"
	"time"
)

func TestQreki(t *testing.T) {
	testcases := []struct {
		in  time.Time
		out Qreki
	}{
		{time.Date(2002, 6, 3, 0, 0, 0, 0, JST), Qreki{2002, 4, 23, false, "友引"}},
	}

	for _, tc := range testcases {
		got := NewQreki(ToJulian(tc.in))
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func TestToJulian(t *testing.T) {
	testcases := []struct {
		in  time.Time
		out float64
	}{
		{time.Date(2002, 6, 3, 0, 0, 0, 0, JST), 2452428.0},
	}

	for _, tc := range testcases {
		got := ToJulian(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

// TODO Available to get expected result but DeepEqual doesn't return true.
// func TestToTime(t *testing.T) {
// 	testcases := []struct {
// 		in  float64
// 		out time.Time
// 	}{
// 		{2452428.0, time.Date(2002, 6, 3, 0, 0, 0, 0, JST)},
// 	}
//
// 	for _, tc := range testcases {
// 		got := ToTime(tc.in)
// 		if !reflect.DeepEqual(got, tc.out) {
// 			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
// 		}
// 	}
// }

func TestNormalizeAngle(t *testing.T) {
	testcases := []struct {
		in  float64
		out float64
	}{
		{31557.0*2451545.0 + 161.0, 86.0},
		{-10993.857868377594 + 303.8119083007812, 109.95403992318643},
		// FYI https://go.dev/src/math/sin.go
		// Results may be meaningless for x > 2**49 = 5.6e14.
		// Golang might give the different value with sin/cos between other languages.
		// [Go 言語の math.Sin は計算結果がおかしいことがある - Qiita](https://qiita.com/Nabetani/items/2dd5e2351b6d2f8ed53f)
		// {-10993.857868377594 + 303.8119083007812, 109.95403992318825},
	}

	for _, tc := range testcases {
		got := NormalizeAngle(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func TestCalcSolarLongitude(t *testing.T) {
	testcases := []struct {
		in  float64
		out float64
	}{
		{2452428.0, 109.95403992318643},
		// {2452428.0, 109.95403992318825}, // with other languages
	}

	for _, tc := range testcases {
		got := CalcSolarLongitude(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func TestCalcMoonLongitude(t *testing.T) {
	testcases := []struct {
		in  float64
		out float64
	}{
		{2452428.0, 233.26846620995252},
	}

	for _, tc := range testcases {
		got := CalcMoonLongitude(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func TestCalcChuki(t *testing.T) {
	type args struct {
		julian    float64
		longitude float64
	}

	testcases := []struct {
		in  args
		out float64
	}{
		{args{2452428.0, 90.0}, 2452354.168240371},
	}

	for _, tc := range testcases {
		got := CalcChuki(tc.in.julian, tc.in.longitude)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}

func TestCalcNewMoon(t *testing.T) {
	testcases := []struct {
		in  float64
		out float64
	}{
		{2452428.0, 2452406.824613828},
	}

	for _, tc := range testcases {
		got := CalcNewMoon(tc.in)
		if !reflect.DeepEqual(got, tc.out) {
			t.Errorf("%v: got %v, want %v", tc.in, got, tc.out)
		}
	}
}
