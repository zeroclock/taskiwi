package utils_test

import (
	"github.com/stretchr/testify/assert"
	"time"

	"taskiwi/utils"
	"testing"
)

func TestUnique_exist_duplicate(t *testing.T) {
	tags := []string{"tagA", "tagB", "tagC", "tagC", "tagB", ""}
	expected := []string{"tagA", "tagB", "tagC"}
	result := utils.Unique(tags)
	assert.Equal(t, expected, result)
}

func TestUnique_no_duplicate(t *testing.T) {
	tags := []string{"tagA", "tagB", "tagC"}
	expected := []string{"tagA", "tagB", "tagC"}
	result := utils.Unique(tags)
	assert.Equal(t, expected, result)
}

func TestTruncateDateHMS(t *testing.T) {
	const layout = "2006-01-02 15:04"
	date, _ := time.Parse(layout, "2020-01-01 15:24")
	expected, _ := time.Parse(layout, "2020-01-01 00:00")
	result := utils.TruncateDateHMS(date)
	assert.True(t, result.Equal(expected))
}
