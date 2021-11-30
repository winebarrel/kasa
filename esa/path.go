package esa

import (
	"strings"
)

type Path struct {
	Category string
	Name     string
}

func NewPath(path string) *Path {
	if strings.HasSuffix(path, "/") {
		return &Path{
			Category: path,
		}
	}

	names := strings.Split(path, "/")

	if len(names) == 0 {
		return &Path{}
	} else if len(names) == 1 {
		return &Path{
			Name: names[0],
		}
	}

	return &Path{
		Category: strings.Join(names[0:len(names)-1], "/"),
		Name:     names[len(names)-1],
	}
}

func (path *Path) IsCategory() bool {
	return path.Name == ""
}
