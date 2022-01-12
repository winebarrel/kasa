package utils

import (
	"strings"
)

func TagsToString(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	formatted := make([]string, 0, len(tags))

	for _, t := range tags {
		formatted = append(formatted, strings.TrimLeft(t, "#"))
	}

	return "[#" + strings.Join(formatted, ",#") + "]"
}
