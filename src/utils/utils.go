package utils

import (
	"time"
	
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

func TruncateDateHMS(datetime time.Time) time.Time {
	hour := time.Duration(datetime.Hour())
	minute := time.Duration(datetime.Minute())
	second := time.Duration(datetime.Second())
	datetime = datetime.Add(-hour * time.Hour)
	datetime = datetime.Add(-minute * time.Minute)
	datetime = datetime.Add(-second * time.Second)
	return datetime
}
