package hw02unpackstring

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(tc string) (string, error) {
	// Place your code here.

	runes := []rune(tc)
	var builder strings.Builder
	var escape bool
	var prev rune
	var printPrev bool

	for i, r := range runes {
		switch {
		case escape:
			{
				if r != '\\' && !unicode.IsDigit(r) {
					return "", ErrInvalidString
				}
				escape = false
				if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
					prev = r
					printPrev = true
				} else {
					builder.WriteString(string(r))
				}
			}
		case r == '\\':
			{
				escape = true
			}
		case unicode.IsDigit(r):
			{
				if i == 0 {
					//все пропало
					return "", ErrInvalidString
				}
				if printPrev || !unicode.IsDigit(runes[i-1]) {
					builder.WriteString(strings.Repeat(string(prev), int(r-'0')))
					printPrev = false
				} else {
					//до этого тоже была цифра но неэкранированная
					return "", ErrInvalidString
				}
			}
		case i+1 < len(runes) && unicode.IsDigit(runes[i+1]):
			{
				prev = r
			}
		default:
			{
				builder.WriteString(string(r))
			}
		}

	}

	fmt.Println(builder.String())
	return builder.String(), nil
}
