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
// func ToTime
// func calcEclipticLongitude
// func getNishinibun
// func CalcChuki
// func NormalizeAngle
// func CalcSolarLongitude
// func CalcMoonLongitude
// func CalcNewMoon

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

func ToTime(julian float64) time.Time {
	return time.UnixMilli(int64((julian + tz - 2440587.0) * 864e5))
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
	var t = julian
	var th float64 = 0.0
	th += 0.0003 * math.Cos(k*NormalizeAngle(2322131.0*t+191.0))
	th += 0.0003 * math.Cos(k*NormalizeAngle(4067.0*t+70.0))
	th += 0.0003 * math.Cos(k*NormalizeAngle(549197.0*t+220.0))
	th += 0.0003 * math.Cos(k*NormalizeAngle(1808933.0*t+58.0))
	th += 0.0003 * math.Cos(k*NormalizeAngle(349472.0*t+337.0))
	th += 0.0003 * math.Cos(k*NormalizeAngle(381404.0*t+354.0))
	th += 0.0003 * math.Cos(k*NormalizeAngle(958465.0*t+340.0))
	th += 0.0004 * math.Cos(k*NormalizeAngle(12006.0*t+187.0))
	th += 0.0004 * math.Cos(k*NormalizeAngle(39871.0*t+223.0))
	th += 0.0005 * math.Cos(k*NormalizeAngle(509131.0*t+242.0))
	th += 0.0005 * math.Cos(k*NormalizeAngle(1745069.0*t+24.0))
	th += 0.0005 * math.Cos(k*NormalizeAngle(1908795.0*t+90.0))
	th += 0.0006 * math.Cos(k*NormalizeAngle(2258267.0*t+156.0))
	th += 0.0006 * math.Cos(k*NormalizeAngle(111869.0*t+38.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(27864.0*t+127.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(485333.0*t+186.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(405201.0*t+50.0))
	th += 0.0007 * math.Cos(k*NormalizeAngle(790672.0*t+114.0))
	th += 0.0008 * math.Cos(k*NormalizeAngle(1403732.0*t+98.0))
	th += 0.0009 * math.Cos(k*NormalizeAngle(858602.0*t+129.0))
	th += 0.0011 * math.Cos(k*NormalizeAngle(1920802.0*t+186.0))
	th += 0.0012 * math.Cos(k*NormalizeAngle(1267871.0*t+249.0))
	th += 0.0016 * math.Cos(k*NormalizeAngle(1856938.0*t+152.0))
	th += 0.0018 * math.Cos(k*NormalizeAngle(401329.0*t+274.0))
	th += 0.0021 * math.Cos(k*NormalizeAngle(341337.0*t+16.0))
	th += 0.0021 * math.Cos(k*NormalizeAngle(71998.0*t+85.0))
	th += 0.0021 * math.Cos(k*NormalizeAngle(990397.0*t+357.0))
	th += 0.0022 * math.Cos(k*NormalizeAngle(818536.0*t+151.0))
	th += 0.0023 * math.Cos(k*NormalizeAngle(922466.0*t+163.0))
	th += 0.0024 * math.Cos(k*NormalizeAngle(99863.0*t+122.0))
	th += 0.0026 * math.Cos(k*NormalizeAngle(1379739.0*t+17.0))
	th += 0.0027 * math.Cos(k*NormalizeAngle(918399.0*t+182.0))
	th += 0.0028 * math.Cos(k*NormalizeAngle(1934.0*t+145.0))
	th += 0.0037 * math.Cos(k*NormalizeAngle(541062.0*t+259.0))
	th += 0.0038 * math.Cos(k*NormalizeAngle(1781068.0*t+21.0))
	th += 0.0040 * math.Cos(k*NormalizeAngle(133.0*t+29.0))
	th += 0.0040 * math.Cos(k*NormalizeAngle(1844932.0*t+56.0))
	th += 0.0040 * math.Cos(k*NormalizeAngle(1331734.0*t+283.0))
	th += 0.0050 * math.Cos(k*NormalizeAngle(481266.0*t+205.0))
	th += 0.0052 * math.Cos(k*NormalizeAngle(31932.0*t+107.0))
	th += 0.0068 * math.Cos(k*NormalizeAngle(926533.0*t+323.0))
	th += 0.0079 * math.Cos(k*NormalizeAngle(449334.0*t+188.0))
	th += 0.0085 * math.Cos(k*NormalizeAngle(826671.0*t+111.0))
	th += 0.0100 * math.Cos(k*NormalizeAngle(1431597.0*t+315.0))
	th += 0.0107 * math.Cos(k*NormalizeAngle(1303870.0*t+246.0))
	th += 0.0110 * math.Cos(k*NormalizeAngle(489205.0*t+142.0))
	th += 0.0125 * math.Cos(k*NormalizeAngle(1443603.0*t+52.0))
	th += 0.0154 * math.Cos(k*NormalizeAngle(75870.0*t+41.0))
	th += 0.0304 * math.Cos(k*NormalizeAngle(513197.9*t+222.5))
	th += 0.0347 * math.Cos(k*NormalizeAngle(445267.1*t+27.9))
	th += 0.0409 * math.Cos(k*NormalizeAngle(441199.8*t+47.4))
	th += 0.0458 * math.Cos(k*NormalizeAngle(854535.2*t+148.2))
	th += 0.0533 * math.Cos(k*NormalizeAngle(1367733.1*t+280.7))
	th += 0.0571 * math.Cos(k*NormalizeAngle(377336.3*t+13.2))
	th += 0.0588 * math.Cos(k*NormalizeAngle(63863.5*t+124.2))
	th += 0.1144 * math.Cos(k*NormalizeAngle(966404.0*t+276.5))
	th += 0.1851 * math.Cos(k*NormalizeAngle(35999.05*t+87.53))
	th += 0.2136 * math.Cos(k*NormalizeAngle(954397.74*t+179.93))
	th += 0.6583 * math.Cos(k*NormalizeAngle(890534.22*t+145.7))
	th += 1.2740 * math.Cos(k*NormalizeAngle(413335.35*t+10.74))
	th += 6.2888 * math.Cos(k*NormalizeAngle(477198.868*t+44.963))

	var ang = NormalizeAngle(481267.8809 * t)
	ang = NormalizeAngle(ang + 218.3162)
	th = NormalizeAngle(th + ang)
	return th
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
		if count == 1 && d < 0 {
			d = NormalizeAngle(d)
		} else if sl >= 0 && sl <= 20 && ml >= 300 {
			d = NormalizeAngle(d)
			d = 360.0 - d
		} else if math.Abs(d) > 40 {
			d = NormalizeAngle(d)
		}

		dt1 = math.Floor(d * 29.530589 / 360.0)
		dt2 = d*29.530589/360.0 - dt1
		julianIntegerPart = julianIntegerPart - dt1
		julianDecimalPart = julianDecimalPart - dt2
		if julianDecimalPart < 0 {
			julianIntegerPart -= 1.0
			julianDecimalPart += 1.0
		}

		if count == 15 && math.Abs(dt1+dt2) > (1.0/86400.0) {
			julianIntegerPart = math.Floor(sl - 26)
			julianDecimalPart = 0.0
		} else if count > 30 && math.Abs(dt1+dt2) > (1.0/86400.0) {
			julianIntegerPart = sl
			julianDecimalPart = 0.0
			break
		}

		count += 1
	}

	return julianDecimalPart + julianIntegerPart - tz
}
