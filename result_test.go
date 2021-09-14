package fslint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	results := NewResults()

	r := LintResult{
		FilePath: "/demo",
		Expect:   ModeCamelCase,
		Level:    LevelWarn,
	}

	results.Append(r)

	assert.True(t, results.has(r))

	assert.Equal(t, 1, results.WarnCount())
	assert.Equal(t, 0, results.ErrorCount())
}
