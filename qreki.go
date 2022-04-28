package qreki

import (
	"time"
)

/* vars */
// var JST
// var timezoneOffsetOfJapan
// var minutesOf24hours
// var tz

/* types */
// type Rokuyo
// type Qreki
// func NewQreki

/* funcs */
// func ToJulian
// func NormalizeAngle

/***** vars *****/
var JST, _ = time.LoadLocation("Asia/Tokyo")
var timezoneOffsetOfJapan = -540.0
var minutesOf24hours = 1440.0
var tz = timezoneOffsetOfJapan / minutesOf24hours

/***** types *****/

type Julian float64

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
func ToJulian(t time.Time) Julian {
	return Julian(2440587.0 + float64(t.UnixMilli())/864e5 - tz)
}

func NormalizeAngle(angle float64) float64 {
	return angle - 360.0*float64(int(angle/360.0))
}
