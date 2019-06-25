package piscine

import"unicode/utf8"

func LastRune(s string) rune{
	ss,_:=utf8.DecodeLastRune([]byte(s))
	return ss
}
