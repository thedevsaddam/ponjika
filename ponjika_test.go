package ponjika

import (
	"testing"
	"time"
)

func Test_isLeapYear(t *testing.T) {
	testCases := []struct {
		Tag        string
		Year       int
		IsLeapYear bool
	}{
		{
			"Year 1900 is NOT leap year",
			1900,
			false,
		},
		{
			"Year 2000 is leap year",
			2000,
			true,
		},
		{
			"Year 2004 is leap year",
			2004,
			true,
		},
	}

	for _, tc := range testCases {
		if isLeapYear(tc.Year) != tc.IsLeapYear {
			t.Error(tc.Tag)
		}
	}
}

func Test_EnToBnYear(t *testing.T) {
	testCases := []struct {
		Month, Date, Year, BYear int
	}{
		{
			Month: 3, Date: 13, Year: 2015, BYear: 1421,
		},
		{
			Month: 5, Date: 10, Year: 2016, BYear: 1423,
		},
	}
	// TODO: Need more test cases and accurate data
	for _, tc := range testCases {
		if r := EnToBnYear(tc.Year, tc.Month, tc.Date); r != tc.BYear {
			t.Error("expected: ", tc.BYear, "go: ", r)
		}
	}
}

func Test_New(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	testCases := []struct {
		Tag                      string
		EnDate                   string
		ExptdBnYear, ExptdBnDate int
		ExptdBnMonth, ExptdBnDay string
	}{
		{
			Tag:          "Expeted result for 2018-03-28 15:30:00 is 1424 bengali year and 14 bengali date",
			EnDate:       "2018-03-28 15:30:00",
			ExptdBnYear:  1424,
			ExptdBnDate:  14,
			ExptdBnMonth: "চৈত্র",
			ExptdBnDay:   "বুধবার",
		},
		{
			Tag:          "Expeted result for 2018-03-29 15:30:00 is 1424 bengali year and 15 bengali date",
			EnDate:       "2018-03-29 15:30:00",
			ExptdBnYear:  1424,
			ExptdBnDate:  15,
			ExptdBnMonth: "চৈত্র",
			ExptdBnDay:   "বৃহস্পতিবার",
		},
		{
			Tag:          "Expeted result for 2018-04-14 14:18:00 is 1425 bengali year and 1 bengali date. Pohela boisakh",
			EnDate:       "2018-04-14 14:18:00",
			ExptdBnYear:  1425,
			ExptdBnDate:  1,
			ExptdBnMonth: "বৈশাখ",
			ExptdBnDay:   "শনিবার",
		},
		{
			Tag:          "Expeted result for 2001-03-14 14:18:00 is expected as non leap year and falun should be 30 days long",
			EnDate:       "2001-03-14 14:18:00",
			ExptdBnYear:  1407,
			ExptdBnDate:  30,
			ExptdBnMonth: "ফাল্গুন",
			ExptdBnDay:   "বুধবার",
		},
		{
			Tag:          "Expeted result for 2004-03-14 14:18:00 is expected as leap year and falun should be 31 days long",
			EnDate:       "2004-03-14 14:18:00",
			ExptdBnYear:  1410,
			ExptdBnDate:  31,
			ExptdBnMonth: "ফাল্গুন",
			ExptdBnDay:   "রবিবার",
		},
	}

	for _, tc := range testCases {
		d, err := time.Parse(layout, tc.EnDate)
		if err != nil {
			t.Error("failed to parse english date time: ", err)
		} else {
			p := New(d)
			if p.Year != tc.ExptdBnYear ||
				p.Date != tc.ExptdBnDate ||
				p.BengaliMonth.Bengali != tc.ExptdBnMonth ||
				p.BengaliDay.Bengali != tc.ExptdBnDay {
				t.Errorf("%s\nExpected: year %d date %d month %s day %s \nGot: year %d date %d month %s day %s\n",
					tc.Tag, tc.ExptdBnYear, tc.ExptdBnDate, tc.ExptdBnMonth, tc.ExptdBnDay,
					p.Year, p.Date, p.BengaliMonth.Bengali, p.BengaliDay.Bengali)
			}
		}
	}
}

func TestPonjika_String(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	testCases := []struct {
		EnDate, ExpectedBnDate string
	}{
		{
			EnDate:         "2001-03-14 14:18:00",
			ExpectedBnDate: "৩০ ফাল্গুন ১৪০৭ রোজ বুধবার",
		},
		{
			EnDate:         "2001-03-14 14:18:00",
			ExpectedBnDate: "৩০ ফাল্গুন ১৪০৭ রোজ বুধবার",
		},
		{
			EnDate:         "2018-04-14 14:18:00",
			ExpectedBnDate: "১ বৈশাখ ১৪২৫ রোজ শনিবার",
		},
	}
	for _, tc := range testCases {
		d, err := time.Parse(layout, tc.EnDate)
		if err != nil {
			t.Error("failed to parse english date time: ", err)
		} else {
			p := New(d)
			if p.String() != tc.ExpectedBnDate {
				t.Errorf("Expected %s \nGot: %s", p.String(), tc.ExpectedBnDate)
			}
		}
	}
}

func TestPonjika_Phonetic(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	testCases := []struct {
		EnDate, ExpectedBnDate string
	}{
		{
			EnDate:         "2001-03-14 14:18:00",
			ExpectedBnDate: "30 Falgun 1407 Roj Budhbar",
		},
		{
			EnDate:         "2001-03-14 14:18:00",
			ExpectedBnDate: "30 Falgun 1407 Roj Budhbar",
		},
		{
			EnDate:         "2018-04-14 14:18:00",
			ExpectedBnDate: "1 Boisakh 1425 Roj Shonibar",
		},
	}
	for _, tc := range testCases {
		d, err := time.Parse(layout, tc.EnDate)
		if err != nil {
			t.Error("failed to parse english date time: ", err)
		} else {
			p := New(d)
			if p.Phonetic() != tc.ExpectedBnDate {
				t.Errorf("Expected %s \nGot: %s", p.Phonetic(), tc.ExpectedBnDate)
			}
		}
	}
}

func Test_BanglaMonthTotalDays(t *testing.T) {
	layout := "2006-01-02 15:04:05"
	testCases := []struct {
		EnDate        string
		ExpectedIndex int
	}{
		{
			EnDate:        "2018-04-01 14:18:00",
			ExpectedIndex: 30,
		},
		{
			EnDate:        "2018-04-15 14:18:00",
			ExpectedIndex: 31,
		},
	}
	for _, tc := range testCases {
		d, err := time.Parse(layout, tc.EnDate)
		if err != nil {
			t.Error("failed to parse english date time: ", err)
		} else {
			p := New(d)
			if p.TotalDays != tc.ExpectedIndex {
				t.Errorf("Expected %d \nGot: %d", tc.ExpectedIndex, p.TotalDays)
			}
		}
	}
}
