package utils

import (
	"time"
	
	"github.com/zeroclock/taskiwi/model"
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

func SetDateHMSToEndTime(datetime time.Time) time.Time {
	datetime = TruncateDateHMS(datetime)
	datetime = datetime.Add(time.Hour * 23)
	datetime = datetime.Add(time.Minute * 59)
	datetime = datetime.Add(time.Second * 59)
	return datetime
}
