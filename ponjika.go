// Package ponjika provide basic functionalities for working with bengali calendar
// Algorithm are directly ported from Nuhil Mehdy's [bangla-calendar] https://github.com/nuhil/bangla-calendar
package ponjika

import (
	"fmt"
	"time"
)

// BN describes both bengali font and phonetic meaning
type BN struct {
	Bengali, Phonetic string
}

// BanglaMonthsList contains bengali month names both bengali font and phonetic
var BanglaMonthsList = [...]BN{
	{"পৌষ", "Poush"},
	{"মাঘ", "Maagh"},
	{"ফাল্গুন", "Falgun"},
	{"চৈত্র", "Chaitra"},
	{"বৈশাখ", "Boisakh"},
	{"জ্যৈষ্ঠ", "Joistho"},
	{"আষাঢ়", "Ashar"},
	{"শ্রাবণ", "Shraban"},
	{"ভাদ্র", "Vadro"},
	{"আশ্বিন", "Ashin"},
	{"কার্তিক", "Kartik"},
	{"অগ্রহায়ণ", "Agrahan"},
}

// WeekDayList contains list of day weeks both bengali font and phonetic
var WeekDayList = [...]BN{
	{"রবিবার", "Robibar"},
	{"সোমবার", "Sombar"},
	{"মঙ্গলবার", "Mongolbar"},
	{"বুধবার", "Budhbar"},
	{"বৃহস্পতিবার", "Brihospotibar"},
	{"শুক্রবার", "Shukrobar"},
	{"শনিবার", "Shonibar"},
}

// BanglaSeasonsList contains bengali seasons
var BanglaSeasonsList = [...]BN{
	{"শীত", "Sheet"},
	{"বসন্ত", "Bosonto"},
	{"গ্রীষ্ম", "Grismo"},
	{"বর্ষা", "Borsha"},
	{"শরৎ", "Sorot"},
	{"হেমন্ত", "Hemonto"}}

var midMonthDate = [...]int{13, 12, 14, 13, 14, 14, 15, 15, 15, 15, 14, 14}

var totalMonthDays = [...]int{30, 30, 30, 30, 31, 31, 31, 31, 31, 30, 30, 30}
var leapYearIndex = 2  //Leap Year will affect only the day count in 'Falgun'
var lastMonthIndex = 3 //'Chaitro' is the last month and it's index is 3 in banglaMonthsList

func isLeapYear(year int) bool {
	return ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0)
}

//EnToBnYear calculate bengali year respect to the Gregorian Year, month and date
func EnToBnYear(year, month, date int) int {
	var banglaYear = year - 594 //2017(Gregorian Year) - 594 = 1423(Bangla Year)
	//if the month is after 'chaitro' then it is a bangla new year, hence the year count will be one more
	if (month > lastMonthIndex) || (month == lastMonthIndex && date > 13) {
		banglaYear = banglaYear + 1
	}
	return banglaYear
}

// Ponjika describes basic information of bengali panjika
type Ponjika struct {
	Year, Date, MonthIndex, TotalDays                                 int
	BengaliYear, BengaliMonth, BengaliDate, BengaliSeason, BengaliDay BN
}

// String satisfy stringer inteface
func (p Ponjika) String() string {
	return fmt.Sprintf("%s %s %s রোজ %s",
		p.BengaliDate.Bengali, p.BengaliMonth.Bengali, p.BengaliYear.Bengali, p.BengaliDay.Bengali)
}

// Phonetic return bengali date string using phonetic
func (p Ponjika) Phonetic() string {
	return fmt.Sprintf("%s %s %s Roj %s",
		p.BengaliDate.Phonetic, p.BengaliMonth.Phonetic, p.BengaliYear.Phonetic, p.BengaliDay.Phonetic)
}

// New return a new instance of Ponjika
func New(t time.Time) Ponjika {
	localTime := t.Add(-6) // align time to Asia/Dhaka local time
	gYear, gMMonth, gDay := localTime.Date()
	gMonth := int(gMMonth) - 1 // considering the month index 0
	var banglaYear, banglaDate, banglaMonthIndex int
	var banglaMonth, banglaSeason BN

	banglaYear = EnToBnYear(gYear, gMonth, gDay)

	monthDays := totalMonthDays[gMonth] //In a leap year, for 'Falgun' month total number of Month Days will be 31 instead of 30
	if gDay <= midMonthDate[gMonth] {
		if gMonth == leapYearIndex && isLeapYear(gYear) {
			monthDays = totalMonthDays[gMonth] + 1
		}
		banglaDate = monthDays + gDay - midMonthDate[gMonth]
		banglaMonthIndex = gMonth
		banglaMonth = BanglaMonthsList[banglaMonthIndex]
	} else {
		banglaDate = gDay - midMonthDate[gMonth]
		banglaMonthIndex = (gMonth + 1) % 12 //banglaMonthsList is 0-based indexed
		banglaMonth = BanglaMonthsList[banglaMonthIndex]
	}

	banglaSeason = BanglaSeasonsList[banglaMonthIndex/2]

	return Ponjika{
		Year:          banglaYear,
		BengaliYear:   BN{enToBnNumber(banglaYear), fmt.Sprintf("%v", banglaYear)},
		BengaliMonth:  banglaMonth,
		BengaliDate:   BN{enToBnNumber(banglaDate), fmt.Sprintf("%v", banglaDate)},
		TotalDays:     totalMonthDays[banglaMonthIndex],
		Date:          banglaDate,
		BengaliDay:    WeekDayList[int(localTime.Weekday())],
		BengaliSeason: banglaSeason,
		MonthIndex:    banglaMonthIndex,
	}
}
