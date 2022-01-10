package postname

import "strings"

func Split(fullName string) (string, string) {
	if fullName == "" {
		return "", ""
	}

	if strings.HasSuffix(fullName, "/") {
		return fullName, ""
	}

	names := strings.Split(fullName, "/")

	return strings.Join(names[0:len(names)-1], "/"), names[len(names)-1]
}
