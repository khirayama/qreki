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
		in  Julian
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

func TestCalcChuki(t *testing.T) {
	type args struct {
		julian    Julian
		longitude float64
	}

	testcases := []struct {
		in  args
		out Julian
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
