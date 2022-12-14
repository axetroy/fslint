package fslint

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {

	b, err := json.Marshal(Config{
		Include: []Selector{},
	})

	assert.Equal(t, err, nil)

	config, err := NewConfig(b)

	assert.Equal(t, err, nil)

	assert.Equal(t, config.Exclude, &DefaultExclude)
}
