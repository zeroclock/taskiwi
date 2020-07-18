package handler

import (
	"github.com/stretchr/testify/assert"

	"taskiwi/model"
	"testing"
)

func TestRouter(t *testing.T) {
	data := &model.ClockData{
		Task: "TaskA",
		Parents: "",
		Category: "",
		Start: "",
		End: "",
		Effort: "",
		Ishabit: "",
		Tags: []string{"TagA", "TagB", "TagC"},
	}

	t.Run("it returns true when tag exists", func(t *testing.T) {
		assert.True(t, hasTag(data, "TagA"))
	})

	t.Run("it returns false when tag doesn't exist", func(t *testing.T) {
		assert.False(t, hasTag(data, "tagc"))
	})
}
