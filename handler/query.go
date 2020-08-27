package handler

import (
	"fmt"
	"time"

	"github.com/zeroclock/taskiwi/model"
	"github.com/zeroclock/taskiwi/utils"
)

const layout = "2006-01-02 15:04"

// received date's format
const layout_req = "2006-01-02"

func QueryTasks(taskData []model.ClockData, tags []string, start string, end string) ([]model.ClockData, error) {
	result := []model.ClockData{}
	conditionStart, err := time.Parse(layout_req, start)
	if err != nil {
		return result, err
	}
	conditionEnd, err := time.Parse(layout_req, end)
	if err != nil {
		return result, err
	}
	// set start's time to 00:00:00
	conditionStart = utils.TruncateDateHMS(conditionStart)
	// set end's time to 23:59:50
	conditionEnd = utils.SetDateHMSToEndTime(conditionEnd)
	for _, v := range taskData {
		// check only task's start datetime
		clockedDate, err := time.Parse(layout, v.Start)
		if err != nil {
			return result, err
		}
		if len(tags) > 0 {
			if containsTag(&v, tags) && containsRange(clockedDate, conditionStart, conditionEnd) {
				result = append(result, v)
			}
		} else {
			if containsRange(clockedDate, conditionStart, conditionEnd) {
				result = append(result, v)
			}
		}
	}

	return result, nil
}

func GetTagsFromClockDatas(taskData []model.ClockData) []string {
	tags := []string{}
	for _, v := range taskData {
		for _, tag := range v.Tags {
			exists := false
			for _, t := range tags {
				if t == tag {
					exists = true
				}
			}
			if !exists {
				tags = append(tags, tag)
			}
		}
	}
	return tags
}

func AggregateClockDatasByTags(taskDatas []model.ClockData, tags []string) (model.WorkTimes, []model.ClockData, error) {
	var total float64 = 0.0
	m := map[string]float64{}
	var workTimes model.WorkTimes

	for _, tag := range tags {
		m[tag] = 0
		for _, data := range taskDatas {
			if containsTag(&data, []string{tag}) {
				start, err := time.Parse(layout, data.Start)
				if err != nil {
					return model.WorkTimes{}, []model.ClockData{}, err
				}
				end, err := time.Parse(layout, data.End)
				if err != nil {
					return model.WorkTimes{}, []model.ClockData{}, err
				}
				sub := end.Sub(start).Minutes()
				m[tag] += sub
				total += sub
			}
		}
	}

	for k, v := range m {
		workTimes = append(workTimes, &model.WorkTime{
			Tag:     k,
			Time:    fmt.Sprint(v),
			Percent: fmt.Sprint(v / total),
		})
	}

	return workTimes, taskDatas, nil
}

func containsTag(task *model.ClockData, tags []string) bool {
	for _, v := range task.Tags {
		for _, tag := range tags {
			if v == tag {
				return true
			}
		}
	}
	return false
}

func containsRange(targetDatetime time.Time, start time.Time, end time.Time) bool {
	return !start.After(targetDatetime) && !end.Before(targetDatetime)
}
