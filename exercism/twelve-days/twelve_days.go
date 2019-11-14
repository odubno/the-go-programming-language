package twelve

import (
	"strings"
)

var lyrics = []struct {
	day, lyric string
}{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song returns the full song
func Song() string {
	// get all lines
	return strings.Join(getLyrics(), "\n")
}

// Verse returns the specific verse
func Verse(i int) string {
	return getLyrics()[i-1]
}

// getLyrics creates full song
func getLyrics() []string {
	// use song to store the full song
	var song []string
	// Uses verses to track all used lyrics
	var verses []string
	for idx := 0; idx <= 11; idx++ {
		line := lyrics[idx]
		if idx == 0 {
			verses = prepend(verses, "and "+line.lyric+".")
			verse := "On the " + line.day + " day of Christmas my true love gave to me: " + line.lyric + "."
			song = append(song, verse)
		} else {
			verses = prepend(verses, line.lyric+",")
			verse := "On the " + line.day + " day of Christmas my true love gave to me: " + strings.Join(verses, " ")
			song = append(song, verse)
		}
	}
	return song
}

func prepend(arr []string, item string) []string {
	return append([]string{item}, arr...)
}
