package qreki

import (
	"time"
)

/* vars */
// var JST

/* utils */

/* types */
// type Rokuyo
// type Qreki
// func NewQreki

/* funcs */
// func ToJulian

/***** vars *****/
var JST, _ = time.LoadLocation("Asia/Tokyo")

/***** utils *****/

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

func NewQreki(t time.Time) Qreki {
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
	tz := -540.0 / 1440.0
	return Julian(2440587.0 + float64(t.UnixMilli())/864e5 - tz)
}
