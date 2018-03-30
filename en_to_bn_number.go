package ponjika

import "fmt"

// enToBnNumber covert english number to bengali number String
func enToBnNumber(d int) string {
	var o string
	ds := fmt.Sprintf("%v", d)
	bdmap := map[string]string{
		"1": "১",
		"2": "২",
		"3": "৩",
		"4": "৪",
		"5": "৫",
		"6": "৬",
		"7": "৭",
		"8": "৮",
		"9": "৯",
		"0": "০",
	}
	for i := 0; i < len(ds); i++ {
		o += bdmap[string(ds[i])]
	}
	return o
}
