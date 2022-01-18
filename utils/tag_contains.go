package utils

import (
	"strings"
)

func TagContains(tags []string, tag string) bool {
	tag = strings.TrimLeft(tag, "#")

	for _, t := range tags {
		t = strings.TrimLeft(t, "#")

		if t == tag {
			return true
		}
	}

	return false
}
