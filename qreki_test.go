package qreki

import (
	"reflect"
	"testing"
	"time"
)

func TestQreki(t *testing.T) {
	testcases := []struct {
		in  time.Time
		out Qreki
	}{
		{time.Date(2002, 6, 3, 0, 0, 0, 0, JST), Qreki{2002, 4, 23, false, "友引"}},
		// https://github.com/shogo82148/go-qreki/blob/33f500c22366f2c5796db35c29693263a909c473/qreki_test.go
		{time.Date(2015, 12, 31, 0, 0, 0, 0, JST), Qreki{2015, 11, 21, false, "先勝"}},
		{time.Date(2016, 1, 1, 0, 0, 0, 0, JST), Qreki{2015, 11, 22, false, "友引"}},
		{time.Date(2016, 2, 7, 0, 0, 0, 0, JST), Qreki{2015, 12, 29, false, "仏滅"}},
		{time.Date(2016, 2, 8, 0, 0, 0, 0, JST), Qreki{2016, 1, 1, false, "先勝"}},
		{time.Date(2016, 3, 8, 0, 0, 0, 0, JST), Qreki{2016, 1, 30, false, "赤口"}},
		{time.Date(2016, 3, 9, 0, 0, 0, 0, JST), Qreki{2016, 2, 1, false, "友引"}},
		{time.Date(2016, 4, 6, 0, 0, 0, 0, JST), Qreki{2016, 2, 29, false, "赤口"}},
		{time.Date(2016, 4, 7, 0, 0, 0, 0, JST), Qreki{2016, 3, 1, false, "先負"}},
		{time.Date(2016, 5, 6, 0, 0, 0, 0, JST), Qreki{2016, 3, 30, false, "友引"}},
		{time.Date(2016, 5, 7, 0, 0, 0, 0, JST), Qreki{2016, 4, 1, false, "仏滅"}},
		{time.Date(2016, 6, 4, 0, 0, 0, 0, JST), Qreki{2016, 4, 29, false, "友引"}},
		{time.Date(2016, 6, 5, 0, 0, 0, 0, JST), Qreki{2016, 5, 1, false, "大安"}},
		{time.Date(2016, 7, 3, 0, 0, 0, 0, JST), Qreki{2016, 5, 29, false, "先負"}},
		{time.Date(2016, 7, 4, 0, 0, 0, 0, JST), Qreki{2016, 6, 1, false, "赤口"}},
		{time.Date(2016, 8, 2, 0, 0, 0, 0, JST), Qreki{2016, 6, 30, false, "大安"}},
		{time.Date(2016, 8, 3, 0, 0, 0, 0, JST), Qreki{2016, 7, 1, false, "先勝"}},
		{time.Date(2016, 8, 31, 0, 0, 0, 0, JST), Qreki{2016, 7, 29, false, "大安"}},
		{time.Date(2016, 9, 1, 0, 0, 0, 0, JST), Qreki{2016, 8, 1, false, "友引"}},
		{time.Date(2016, 9, 30, 0, 0, 0, 0, JST), Qreki{2016, 8, 30, false, "先勝"}},
		{time.Date(2016, 10, 1, 0, 0, 0, 0, JST), Qreki{2016, 9, 1, false, "先負"}},
		{time.Date(2016, 10, 30, 0, 0, 0, 0, JST), Qreki{2016, 9, 30, false, "友引"}},
		{time.Date(2016, 10, 31, 0, 0, 0, 0, JST), Qreki{2016, 10, 1, false, "仏滅"}},
		{time.Date(2016, 11, 28, 0, 0, 0, 0, JST), Qreki{2016, 10, 29, false, "友引"}},
		{time.Date(2016, 11, 29, 0, 0, 0, 0, JST), Qreki{2016, 11, 1, false, "大安"}},
		{time.Date(2016, 12, 28, 0, 0, 0, 0, JST), Qreki{2016, 11, 30, false, "仏滅"}},
		{time.Date(2016, 12, 29, 0, 0, 0, 0, JST), Qreki{2016, 12, 1, false, "赤口"}},
		{time.Date(2017, 1, 1, 0, 0, 0, 0, JST), Qreki{2016, 12, 4, false, "先負"}},
		{time.Date(2017, 1, 27, 0, 0, 0, 0, JST), Qreki{2016, 12, 30, false, "大安"}},
		{time.Date(2017, 1, 28, 0, 0, 0, 0, JST), Qreki{2017, 1, 1, false, "先勝"}},
		{time.Date(2017, 6, 23, 0, 0, 0, 0, JST), Qreki{2017, 5, 29, false, "先負"}},
		{time.Date(2017, 6, 24, 0, 0, 0, 0, JST), Qreki{2017, 5, 1, true, "大安"}},
		{time.Date(2017, 7, 22, 0, 0, 0, 0, JST), Qreki{2017, 5, 29, true, "先負"}},
		{time.Date(2017, 7, 23, 0, 0, 0, 0, JST), Qreki{2017, 6, 1, false, "赤口"}},
		{time.Date(2033, 12, 21, 0, 0, 0, 0, JST), Qreki{2033, 11, 30, false, "仏滅"}},
		{time.Date(2033, 12, 22, 0, 0, 0, 0, JST), Qreki{2033, 11, 1, true, "大安"}},
		{time.Date(2034, 1, 19, 0, 0, 0, 0, JST), Qreki{2033, 11, 29, true, "先負"}},
		{time.Date(2034, 1, 20, 0, 0, 0, 0, JST), Qreki{2033, 12, 1, false, "赤口"}},
		{time.Date(2020, 1, 1, 0, 0, 0, 0, JST), Qreki{2019, 12, 7, false, "赤口"}},
		{time.Date(2030, 1, 1, 0, 0, 0, 0, JST), Qreki{2029, 11, 28, false, "友引"}},
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
		// {args{2452428.0, 90.0}, 2452354.168240371},
		{args{2463952.0, 90.0}, 2463861.675270711},
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
