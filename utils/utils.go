package utils

import "strings"

func Capitalize(s string) string {
	capitalized := ""
	for i, c := range s {
		sym := string(c)
		if i == 0 {
			sym = strings.ToUpper(string(c))
		}
		capitalized = capitalized + sym
	}
	return capitalized
}
