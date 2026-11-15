// “Lowercase the text, except when it starts with # (it’s a code, so don’t touch it).”

package main

import (
	"fmt"
	"strings"
)

type TextFunc func(string) string

func trim(str string) string {
	return strings.TrimSpace(str)
}

func lowwercase(str string) string {
	return strings.ToLower(str)
}

func skipCode(str string) string {
	if str[:1] == "#" {
		return trim(str)
	}

	return trim(lowwercase(str))
}

func Trim(next TextFunc) TextFunc {
	return func(str string) string {
		return strings.TrimSpace(next(str))
	}
}

func SkipCode(next TextFunc) TextFunc {
	return func(x string) string {
		if x[:1] == "#" {
			return x
		}
		return next(x)
	}
}

func main() {
	// composite function
	fmt.Println(skipCode("  HELLO  "))
	fmt.Println(skipCode("#ABC123"))


	// decorator
	clear := SkipCode(Trim(lowwercase))
	fmt.Println(clear("  HELLO  "))
	fmt.Println(clear("#ABC123"))
}
