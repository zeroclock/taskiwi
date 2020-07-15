package utils

import (
	"taskiwi/model"
)

func Unique(tags model.Tags) []string {
	m := make(map[string]bool)
	uniq := [] string{}
	
	for _, ele := range tags {
		if !m[ele] && ele != "" {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}

	return uniq
}
