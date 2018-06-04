package cuprum2018

import (
	"sort"
)

// BeautifulPassword is used to create John's new password which is easier to remember.
// Therefore, it must fulfill certain requirements: he wants each character (digit or letter)
// to appear in the new password an even number of times (possibly zero).
// Also, since he is so proud of his original password, he wants the new password to be a contiguous substring of the original password.
func BeautifulPassword(S string) int {
	for i := len(S); i > 0; i-- {
		if i%2 != 0 {
			continue
		}
		for j := 0; j <= len(S)-i; j++ {
			substring := S[j : i+j]
			if isStringCharactersEven(substring) {
				return len(substring)
			}
		}
	}
	return 0
}

func isStringCharactersEven(s string) bool {
	// Golang 1.8 or later:
	// r := []rune(s)
	// sort.Slice(r, func(i, j int) bool {
	// 	return r[i] < r[j]
	// })

	var r ByRune = []rune(s)
	sort.Sort(r)

	prev := 0
	for i := 0; i < len(r); i++ {
		if i%2 == 1 {
			if int(r[i]) != prev {
				return false
			}
		} else {
			prev = int(r[i])
		}
	}
	return true
}

// ByRune is Type of rune array
type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }
