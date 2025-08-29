//go:build wireinject
// +build wireinject

package cases

import "github.com/google/wire"

func InitializeSimpleService(isErorr bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabaseRepository() (*DatabaseRepository, error) {
	wire.Build(NewDatabasePostgreSQL, NewDatabaseMongoDB, NewDatabaseRepository)
	return nil, nil
}