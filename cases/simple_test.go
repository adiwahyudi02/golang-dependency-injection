package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := InitializeSimpleService(false)
	assert.Nil(t, err)
	assert.NotNil(t, simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := InitializeSimpleService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}