package ponjika

import "testing"

func Test_enToBnNumber(t *testing.T) {
	testCases := map[int]string{
		1:    "১",
		2:    "২",
		3:    "৩",
		4:    "৪",
		5:    "৫",
		6:    "৬",
		7:    "৭",
		8:    "৮",
		9:    "৯",
		0:    "০",
		2018: "২০১৮",
	}
	for en, bn := range testCases {
		if r := enToBnNumber(en); r != bn {
			t.Error("e2dDigit failed", "expected: ", bn, " got: ", r)
		}
	}
}
