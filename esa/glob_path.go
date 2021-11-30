package esa

import (
	"strings"

	"github.com/gobwas/glob"
)

type GlobPath struct {
	category string
	filename string
	pattern  glob.Glob
}

func newGlobPath(path string) *GlobPath {
	catNames := strings.Split(path, "/")
	nonGlobs := []string{}
	border := 0

	for i, cat := range catNames {
		if !strings.HasSuffix(path, "/") && i >= len(catNames)-1 {
			break
		}

		if cat != glob.QuoteMeta(cat) {
			break
		}

		nonGlobs = append(nonGlobs, cat)
		border = i + 1
	}

	rest := catNames[border:]
	globPath := strings.Join(rest, "/")
	filename := ""

	if globPath == "" {
		globPath = "**"
	} else if globPath != glob.QuoteMeta(globPath) {
		filename = globPath
	}

	return &GlobPath{
		category: strings.Join(nonGlobs, "/") + "/",
		filename: filename,
		pattern:  glob.MustCompile(globPath, rune('/')),
	}
}

func (path *GlobPath) match(fullName string) bool {
	return path.pattern.Match(path.base(fullName))
}

func (path *GlobPath) base(fullName string) string {
	return strings.TrimPrefix(fullName[len(path.category)-1:], "/")
}
