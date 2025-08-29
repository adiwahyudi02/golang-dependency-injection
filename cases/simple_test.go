package cases

import (
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := InitializeSimpleService()

	fmt.Println(simpleService.SimpleRepository)
}