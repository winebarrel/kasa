package utils

import (
	"strings"
)

func TagsToString(tags []string) string {
	if len(tags) == 0 {
		return ""
	}

	return "[#" + strings.Join(tags, ",#") + "]"
}
