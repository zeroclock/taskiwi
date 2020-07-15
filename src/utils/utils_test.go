package utils_test

import (
	"github.com/stretchr/testify/assert"

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
