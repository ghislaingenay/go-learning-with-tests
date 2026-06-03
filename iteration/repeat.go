package iteration

import "strings"

const (
	repeatCount = 5
)

func Repeat(character string, count int) string {
	var repeated strings.Builder
	for i := 0; i < 5; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}