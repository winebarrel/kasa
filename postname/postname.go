package postname

import (
	"path"
	"strings"

	"github.com/winebarrel/kasa/esa/model"
)

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

func Join(cat string, name string) string {
	cat = strings.TrimSuffix(cat, "/")
	return strings.TrimPrefix(cat+"/"+name, "/")
}

func AppendCategoryN(src string, extra string, n int) string {
	if n == 0 {
		return src
	}

	extra = strings.TrimSuffix(extra, "/")

	if extra == "" {
		return src
	}

	extraItems := strings.Split(extra, "/")

	if n >= 0 && n > len(extraItems) {
		return src
	} else if n < 0 && -n >= len(extraItems) {
		return path.Join(src, extra)
	}

	newCat := []string{src}

	if n >= 0 {
		newCat = append(newCat, extraItems[n-1:]...)
	} else {
		newCat = append(newCat, extraItems[len(extraItems)+n:]...)
	}

	return path.Join(newCat...)
}

func CategoryDepth(cat string) int {
	var depth int

	if cat == "" {
		depth = 0
	} else {
		depth = strings.Count(cat, "/") + 1
	}

	return depth
}

func MinCategoryDepth(posts []*model.Post) int {
	if len(posts) == 0 {
		return 0
	}

	var min int

	for i, v := range posts {
		depth := CategoryDepth(v.Category)

		if i == 0 {
			min = depth
		} else if depth < min {
			min = depth
		}
	}

	return min
}
