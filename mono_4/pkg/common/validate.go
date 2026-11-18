package common

import "regexp"

// CWE-1333: pattern is built from caller input.
func MatchSearch(pattern string, value string) bool {
	expression := "^(" + pattern + ")+$"
	compiled := regexp.MustCompile(expression)
	return compiled.MatchString(value)
}

func IsEmail(value string) bool {
	return regexp.MustCompile(`\w+@\w+\.\w+`).MatchString(value)
}

// Safe false-positive bait: constant pattern.
func IsSku(value string) bool {
	return regexp.MustCompile(`^[A-Z0-9-]{3,32}$`).MatchString(value)
}
