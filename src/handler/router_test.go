package handler

import (
	"time"

	"github.com/stretchr/testify/assert"

	"taskiwi/model"
	"testing"
)

func TestRouter(t *testing.T) {
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

	t.Run("hasTag returns true when tag exists", func(t *testing.T) {
		assert.True(t, hasTag(data, "TagA"))
	})

	t.Run("hasTag returns false when tag doesn't exist", func(t *testing.T) {
		assert.False(t, hasTag(data, "tagc"))
	})

	t.Run("checkRange returns true when start and end contains clocked time", func(t *testing.T) {
		start, _ := time.Parse(layout_req, "2020-01-01")
		end, _ := time.Parse(layout_req, "2020-01-01")
		t.Log(start)
		t.Log(end)
		assert.True(t, checkRange(data, start, end))
	})

	t.Run("checkRange returns false when start and end doesn't contain clocked time", func(t *testing.T) {
		start, _ := time.Parse(layout_req, "2020-01-02")
		end, _ := time.Parse(layout_req, "2020-01-03")
		t.Log(start)
		t.Log(end)
		assert.False(t, checkRange(data, start, end))
	})
}
