package utils

import (
	"testing"

	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestIs200Success(test_framework *testing.T) {
	assert.Equal(test_framework, Is2xxSuccessful(200), true)
}

func TestIs201Success(test_framework *testing.T) {
	assert.Equal(test_framework, Is2xxSuccessful(201), true)
}

func TestIs202Success(test_framework *testing.T) {
	assert.Equal(test_framework, Is2xxSuccessful(202), true)
}

func TestIs500Fail(test_framework *testing.T) {
	assert.Equal(test_framework, Is2xxSuccessful(500), false)
}
