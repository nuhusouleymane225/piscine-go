package piscine

import"unicode/utf8"

func FirstRune(s string) rune{
	ss,_:=utf8.DecodeRune([]byte(s))
	return ss
}
