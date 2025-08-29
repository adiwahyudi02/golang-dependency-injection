//go:build wireinject
// +build wireinject

package cases

import "github.com/google/wire"

func InitializeSimpleService() *SimpleService {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil
}