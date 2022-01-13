package utils

import (
	"regexp"
	"strconv"
)

func GetPostNum(team, searchPath string) (int, error) {
	r := regexp.MustCompile(`(?:https://` + team + `\.esa\.io/posts/|//)(\d+)$`)
	m := r.FindStringSubmatch(searchPath)

	if len(m) != 2 {
		return 0, nil
	}

	return strconv.Atoi(m[1])
}
