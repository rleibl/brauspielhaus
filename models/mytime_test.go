package models

import (
	"fmt"
	"testing"
)

func TestMyTimeParse(t *testing.T) {

	time_str := "01.01.1970"
	myt := ParseTime(time_str)

	if myt.Day != 1 {
		t.Error("Day")
	}
	if myt.Month != 1 {
		t.Error("Month")
	}
	if myt.Year != 1970 {
		t.Error("Year")
	}

	res := myt.ToString()
	fmt.Println(res)
	fmt.Println(time_str)
	if res != time_str {
		t.Error("ToString()")
	}
}

type TimeDiffT struct {
	t1 string // time 1
	t2 string // time 2
	d  int    // expected duration
}

func TestDuration(t *testing.T) {

	testcases := []TimeDiffT{
		TimeDiffT{"01.01.1970", "02.01.1970", 1},
		TimeDiffT{"01.01.1970", "01.01.1970", 0},
		TimeDiffT{"10.01.1970", "01.01.1970", 9},
		TimeDiffT{"01.01.1970", "01.01.1971", 365},
		TimeDiffT{"01.01.1970", "01.02.1970", 31},
		TimeDiffT{"01.01.1970", "01.02.1971", 396},
	}

	for _, tc := range testcases {

		t1 := ParseTime(tc.t1)
		t2 := ParseTime(tc.t2)

		d := t1.DurationBetween(t2)
		if d != tc.d {
			t.Errorf("Got: %d, expected: %d, from %+v\n", d, tc.d, tc)
		}
	}
}

func TestDurationStr(t *testing.T) {

	testcases := []TimeDiffT{
		TimeDiffT{"01.01.1970", "02.01.1970", 1},
		TimeDiffT{"01.01.1970", "01.01.1970", 0},
		TimeDiffT{"10.01.1970", "01.01.1970", 9},
		TimeDiffT{"01.01.1970", "01.01.1971", 365},
		TimeDiffT{"01.01.1970", "01.02.1970", 31},
		TimeDiffT{"01.01.1970", "01.02.1971", 396},
	}

	for _, tc := range testcases {

		d := Duration(tc.t1, tc.t2)
		if d != tc.d {
			t.Errorf("Got: %d, expected: %d, from %+v\n", d, tc.d, tc)
		}
	}

}
