// utils/extract_parts.go

package utils

import (
	"strings"
)

func ExtractParts(input string) []string {
	parts := strings.Split(input, ",")
	for i, part := range parts {
		if len(part) >= 4 {
			parts[i] = part[:4]
		} else {
			parts[i] = part
		}
	}
	return parts
}
