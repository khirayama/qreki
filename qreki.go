package qreki

import (
	"math"
	"time"
)

/* vars */
// var JST
// var timezoneOffsetOfJapan
// var minutesOf24hours
// var tz
// var k

/* types */
// type Rokuyo
// type Qreki
// func NewQreki

/* funcs */
// func ToJulian
// func CalcChuki
// func NormalizeAngle
// func CalcSolarLongitude

/***** vars *****/
var JST, _ = time.LoadLocation("Asia/Tokyo")
var timezoneOffsetOfJapan = -540.0
var minutesOf24hours = 1440.0
var tz = timezoneOffsetOfJapan / minutesOf24hours
var k = math.Pi / 180

/***** types *****/
// TODO Enumにする？
type Rokuyo string

type Qreki struct {
	Year     int
	Month    int
	Day      int
	Uruu     bool
	Rokuyo   Rokuyo
	Mage     float64
	Magenoon float64
	Illumi   float64
	Mphase   int
}

func NewQreki(julian Julian) Qreki {
	qreki := Qreki{
		Year:     2015,
		Month:    11,
		Day:      21,
		Uruu:     false,
		Rokuyo:   "大安",
		Mage:     0.0,
		Magenoon: 0.0,
		Illumi:   0.0,
		Mphase:   0,
	}
	return qreki
}

/***** funcs *****/
func ToJulian(t time.Time) float64 {
	return 2440587.0 + float64(t.UnixMilli())/864e5 - tz
}

func calcEclipticLongitude(sl float64, longitude float64) float64 {
	return longitude * math.Floor(sl/longitude)
}

func getNishinibun(julian float64) (float64, float64, float64) {
	julianIntegerPart := math.Floor(float64(julian))
	julianDecimalPart := float64(julian) - float64(julianIntegerPart) + tz
	nnj := (julianDecimalPart+0.5)/36525.0 + (julianIntegerPart-2451545.0)/36525.0
	return julianIntegerPart, julianDecimalPart, nnj
}

func CalcChuki(julian float64, longitude float64) float64 {
	julianIntegerPart, julianDecimalPart, nnj := getNishinibun(julian)
	sl := CalcSolarLongitude(nnj)
	el := calcEclipticLongitude(sl, longitude)

	dt1 := 0.0
	dt2 := 1.0
	for math.Abs(dt1+dt2) > 1/86400.0 {
		t := (julianDecimalPart+0.5)/36525.0 + (julianIntegerPart-2451545.0)/36525.0
		sl = CalcSolarLongitude(t)
		ds := sl - el
		if ds > 180.0 {
			ds -= 360.0
		} else if ds < -180.0 {
			ds += 360.0
		}
		dt1 = math.Floor(ds * 365.2 / 360.0)
		dt2 = ds*365.2/360.0 - dt1

		julianIntegerPart = julianIntegerPart - dt1
		julianDecimalPart = julianDecimalPart - dt2
		if julianDecimalPart < 0.0 {
			julianIntegerPart -= 1.0
			julianDecimalPart += 1.0
		}
	}

	return julianDecimalPart + julianIntegerPart - tz
}

func NormalizeAngle(angle float64) float64 {
	return angle - 360.0*math.Floor(angle/360.0)
}

func CalcSolarLongitude(julian float64) float64 {
	var t = julian
	var th float64 = 0.0
	th += 0.0004 * math.Cos(k*NormalizeAngle(31557.0*t+161.0))
	th += 0.0004 * math.Cos(k*NormalizeAngle(29930.0*t+48.0))
	th += 0.0005 * math.Cos(k*NormalizeAngle(2281.0*t+221.0))
	th += 0.0005 * math.Cos(k*NormalizeAngle(155.0*t+118.0))
	th += 0.0006 * math.Cos(k*NormalizeAngle(33718.0*t+316.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(9038.0*t+64.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(3035.0*t+110.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(65929.0*t+45.0))
	th += 0.0013 * math.Cos(k*NormalizeAngle(22519.0*t+352.0))
	th += 0.0015 * math.Cos(k*NormalizeAngle(45038.0*t+254.0))
	th += 0.0018 * math.Cos(k*NormalizeAngle(445267.0*t+208.0))
	th += 0.0018 * math.Cos(k*NormalizeAngle(19.0*t+159.0))
	th += 0.0020 * math.Cos(k*NormalizeAngle(32964.0*t+158.0))
	th += 0.0200 * math.Cos(k*NormalizeAngle(71997.1*t+265.1))

	var ang = NormalizeAngle(35999.05*t + 267.52)
	th = th - 0.0048*t*math.Cos(k*ang)
	th += 1.9147 * math.Cos(k*ang)

	ang = NormalizeAngle(36000.7695 * t)
	ang = NormalizeAngle(ang + 280.4659)
	th = NormalizeAngle(th + ang)
	return th
}

func CalcMoonLongitude(julian float64) float64 {
	// TODO
	return julian
}

func CalcNewMoon(julian float64) float64 {
	count := 1
	julianIntegerPart, julianDecimalPart, _ := getNishinibun(julian)

	dt1 := 0.0
	dt2 := 1.0
	for math.Abs(dt1+dt2) > 1.0/86400.0 {
		t := (julianDecimalPart+0.5)/36525.0 + (julianIntegerPart-2451545.0)/36525.0

		sl := CalcSolarLongitude(t)
		ml := CalcMoonLongitude(t)
		d := ml - sl
		if count == 1 && d < 0.0 {
			d = NormalizeAngle(d)
		} else if sl >= 0 && sl <= 20 && ml >= 300 {
			d = NormalizeAngle(d)
			d = 360.0 - d
		} else if math.Abs(d) > 40 {
			d = NormalizeAngle(d)
		}
		dIntegerPart := math.Floor(d * 29.530589 / 360.0)
		dDecimalPart := d*29.530589/360.0 - dIntegerPart
		julianIntegerPart = julianIntegerPart - dIntegerPart
		julianDecimalPart = julianDecimalPart - dDecimalPart
		if julianDecimalPart < 0 {
			julianIntegerPart -= 1.0
			julianDecimalPart += 1.0
		}

		if count == 15 && math.Abs(dIntegerPart+dDecimalPart) > (1.0/86400.0) {
			julianIntegerPart = math.Floor(sl - 26)
			julianDecimalPart = 0.0
		} else if count > 30 && math.Abs(dIntegerPart+dDecimalPart) > (1.0/86400.0) {
			julianIntegerPart = sl
			julianDecimalPart = 0.0
			break
		}

		count += 1
	}

	return julianDecimalPart + julianIntegerPart - tz
}
