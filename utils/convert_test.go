package utils

import "testing"

func TestStr(t *testing.T) {
	cases := [...][2]string{
		{Str(32), "32"},
		{Str(true), "true"},
		{Str(56.12), "56.12"},
	}

	for i := 0; i < len(cases); i++ {
		result := cases[i][0]
		expect := cases[i][1]

		if result != expect {
			t.Fail()
		}
	}
}
