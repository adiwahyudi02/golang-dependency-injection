package cases

import (
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := InitializeSimpleService()
	fmt.Println(err)
	fmt.Println(simpleService)
}