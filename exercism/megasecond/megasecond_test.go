/*
This is written to practice writing tets in go.
This test is written based off of gigasecond/gigasecond_test.go
*/
package megasecond

import (
	"strings"
	"testing"
	"time"
)

var cases = []struct {
	in   string
	want string
}{
	{
		"2011-04-25",
		"2011-05-06T13:46:40",
	},
	{
		"1977-06-13",
		"1977-06-24T13:46:40",
	},
}

func TestAddMegasecond(t *testing.T) {
	for _, tc := range cases {
		in := parse(tc.in, t)
		expected := parse(tc.want, t)
		got := AddMegasecond(in)
		if !got.Equal(expected) {
			t.Fatalf(`FAIL: AddMegasecond(%s) = %s want %s`, tc.in, got, expected)
		}
	}
	t.Log("PASS")
}

func parse(ts string, t *testing.T) time.Time {
	if strings.Contains(ts, "T") {
		tc, err := time.Parse("2006-01-02T15:04:05", ts)
		if err != nil {
			t.Fatal(err)
		}
		return tc
	}
	tc, err := time.Parse("2006-01-02", ts)
	if err != nil {
		t.Fatal(err)
	}
	return tc
}
