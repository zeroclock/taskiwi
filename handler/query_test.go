package handler

import (
	"fmt"
	"time"

	"github.com/stretchr/testify/assert"

	"testing"

	"github.com/zeroclock/taskiwi/model"
)

func TestHandler(t *testing.T) {
	data := &model.ClockData{
		Task:     "TaskA",
		Parents:  "",
		Category: "",
		Start:    "2020-01-01 12:00",
		End:      "2020-01-01 13:00",
		Effort:   "",
		Ishabit:  "",
		Tags:     []string{"TagA", "TagB", "TagC"},
	}
	data_arr := []model.ClockData{
		{
			Task:     "TaskA",
			Parents:  "",
			Category: "",
			Start:    "2020-01-01 12:00",
			End:      "2020-01-01 13:00",
			Effort:   "",
			Ishabit:  "",
			Tags:     []string{"TagA", "TagB", "TagC"},
		},
		{
			Task:     "TaskB",
			Parents:  "",
			Category: "",
			Start:    "2020-01-01 15:00",
			End:      "2020-01-01 18:30",
			Effort:   "",
			Ishabit:  "",
			Tags:     []string{"TagA", "TagC"},
		},
		{
			Task:     "TaskC",
			Parents:  "",
			Category: "",
			Start:    "2020-01-02 9:30",
			End:      "2020-01-02 10:30",
			Effort:   "",
			Ishabit:  "",
			Tags:     []string{"TagA"},
		},
		{
			Task:     "TaskD",
			Parents:  "",
			Category: "",
			Start:    "2020-01-02 10:30",
			End:      "2020-01-02 12:00",
			Effort:   "",
			Ishabit:  "",
			Tags:     []string{"TagB", "TagC"},
		},
		{
			Task:     "TaskE",
			Parents:  "",
			Category: "",
			Start:    "2020-01-03 00:00",
			End:      "2020-01-03 23:59",
			Effort:   "",
			Ishabit:  "",
			Tags:     []string{"TagA", "TagB", "TagC"},
		},
	}

	t.Run("containsTag returns true when tag exists", func(t *testing.T) {
		assert.True(t, containsTag(data, []string{"TagA"}))
	})

	t.Run("containsTag returns false when tag doesn't exist", func(t *testing.T) {
		assert.False(t, containsTag(data, []string{"tagc"}))
	})

	t.Run("containsRange returns true when start and end contains clocked time", func(t *testing.T) {
		start, _ := time.Parse(layout, "2020-01-01 00:00")
		end, _ := time.Parse(layout, "2020-01-01 23:59")
		checkTime, _ := time.Parse(layout, data.Start)
		t.Log(start)
		t.Log(end)
		assert.True(t, containsRange(checkTime, start, end))
	})

	t.Run("containsRange returns false when start and end doesn't contain clocked time", func(t *testing.T) {
		start, _ := time.Parse(layout, "2020-01-02 00:00")
		end, _ := time.Parse(layout, "2020-01-03 23:59")
		checkTime, _ := time.Parse(layout, data.Start)
		t.Log(start)
		t.Log(end)
		assert.False(t, containsRange(checkTime, start, end))
	})

	t.Run("QueryTasks returns collect ClockDatas search by tags", func(t *testing.T) {
		expected := []model.ClockData{
			{
				Task:     "TaskA",
				Parents:  "",
				Category: "",
				Start:    "2020-01-01 12:00",
				End:      "2020-01-01 13:00",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagB", "TagC"},
			},
			{
				Task:     "TaskD",
				Parents:  "",
				Category: "",
				Start:    "2020-01-02 10:30",
				End:      "2020-01-02 12:00",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagB", "TagC"},
			},
			{
				Task:     "TaskE",
				Parents:  "",
				Category: "",
				Start:    "2020-01-03 00:00",
				End:      "2020-01-03 23:59",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagB", "TagC"},
			},
		}
		actual, _ := QueryTasks(data_arr, []string{"TagB"}, "2020-01-01", "2020-01-03")
		assert.Equal(t, expected, actual)
	})

	t.Run("QueryTasks returns all ClockDatas search by empty tags", func(t *testing.T) {
		expected := []model.ClockData{
			{
				Task:     "TaskA",
				Parents:  "",
				Category: "",
				Start:    "2020-01-01 12:00",
				End:      "2020-01-01 13:00",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagB", "TagC"},
			},
			{
				Task:     "TaskB",
				Parents:  "",
				Category: "",
				Start:    "2020-01-01 15:00",
				End:      "2020-01-01 18:30",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagC"},
			},
			{
				Task:     "TaskC",
				Parents:  "",
				Category: "",
				Start:    "2020-01-02 9:30",
				End:      "2020-01-02 10:30",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA"},
			},
			{
				Task:     "TaskD",
				Parents:  "",
				Category: "",
				Start:    "2020-01-02 10:30",
				End:      "2020-01-02 12:00",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagB", "TagC"},
			},
			{
				Task:     "TaskE",
				Parents:  "",
				Category: "",
				Start:    "2020-01-03 00:00",
				End:      "2020-01-03 23:59",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagB", "TagC"},
			},
		}
		actual, _ := QueryTasks(data_arr, []string{}, "2020-01-01", "2020-01-05")
		assert.Equal(t, expected, actual)
	})

	t.Run("QueryTasks returns collect ClockDatas search by range", func(t *testing.T) {
		expected := []model.ClockData{
			{
				Task:     "TaskA",
				Parents:  "",
				Category: "",
				Start:    "2020-01-01 12:00",
				End:      "2020-01-01 13:00",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagB", "TagC"},
			},
			{
				Task:     "TaskB",
				Parents:  "",
				Category: "",
				Start:    "2020-01-01 15:00",
				End:      "2020-01-01 18:30",
				Effort:   "",
				Ishabit:  "",
				Tags:     []string{"TagA", "TagC"},
			},
		}
		actual, _ := QueryTasks(data_arr, []string{"TagA", "TagB", "TagC"}, "2020-01-01", "2020-01-01")
		assert.Equal(t, expected, actual)
	})

	t.Run("QueryTasks returns empty slice search by range out of start time", func(t *testing.T) {
		expected := []model.ClockData{}
		actual, _ := QueryTasks(data_arr, []string{}, "2019-12-01", "2019-12-31")
		assert.Equal(t, expected, actual)
	})

	t.Run("GetTagsFromClockDatas returns collect slice when tags exist", func(t *testing.T) {
		expected := []string{"TagA", "TagB", "TagC"}
		actual := GetTagsFromClockDatas(data_arr)
		assert.Equal(t, expected, actual)
	})

	t.Run("GetTagsFromClockDatas returns empty when ClockData is empty", func(t *testing.T) {
		expected := []string{}
		actual := GetTagsFromClockDatas([]model.ClockData{})
		assert.Equal(t, expected, actual)
	})

	t.Run("AggregateClockDatasByTags returns correct results", func(t *testing.T) {
		total := 1769.0 + 1589.0 + 1799.0
		ex_workTimes := model.WorkTimes{
			{
				Tag: "TagA",
				Time: "1769",
				Percent: fmt.Sprint(1769.0 / total),
			},
			{
				Tag: "TagB",
				Time: "1589",
				Percent: fmt.Sprint(1589.0 / total),
			},
			{
				Tag: "TagC",
				Time: "1799",
				Percent: fmt.Sprint(1799.0 / total),
			},
		}
		ex_taskDatas := data_arr
		workTimes, taskDatas, _ := AggregateClockDatasByTags(data_arr, []string{
			"TagA",
			"TagB",
			"TagC",
		})
		assert.Equal(t, ex_workTimes, workTimes)
		assert.Equal(t, ex_taskDatas, taskDatas)
	})

}
