package custom_string

import (
	"fmt"
	"strings"
)

// Generate Course Code
func GenerateCodeCourse(name string, count int) string {
	words := strings.Split(name, " ")
	abbreviation := ""
	for _, word := range words {
		abbreviation += string(word[0])
	}

	uppercaseAbbreviation := strings.ToUpper(abbreviation)
	return fmt.Sprintf("%s-%d", uppercaseAbbreviation, count)
}
