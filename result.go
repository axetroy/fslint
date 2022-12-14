package fslint

import "sort"

type LintResult struct {
	FilePath string `json:"filepath"`
	Expect   Mode   `json:"expect"`
	Level    Level  `json:"level"`
}

type Results struct {
	values []LintResult
}

func NewResults() *Results {
	return &Results{
		values: make([]LintResult, 0),
	}
}

func (r *Results) has(result LintResult) bool {
	for _, item := range r.values {
		if result.FilePath == item.FilePath {
			return true
		}
	}

	return false
}

func (r *Results) Append(item LintResult) {
	if !r.has(item) {
		r.values = append(r.values, item)
	}
}

func (r *Results) count(targetLevel Level) int {
	count := 0

	for _, item := range r.values {
		if item.Level == targetLevel {
			count = count + 1
		}
	}

	return count
}

func (r *Results) WarnCount() int {
	return r.count(LevelWarn)
}

func (r *Results) ErrorCount() int {
	return r.count(LevelError)
}

func (r *Results) Values() []LintResult {
	sort.SliceStable(r.values, func(i, j int) bool {
		return r.values[i].FilePath < r.values[j].FilePath
	})

	return r.values
}
