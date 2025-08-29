package cases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	connection, cleanup := InitializedConnection("Databaes")
	assert.NotNil(t, connection)

	cleanup()
}