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

var fooSet = wire.NewSet(
	NewFooRepository,
	NewFooService,
)

var barSet = wire.NewSet(
	NewBarRepository,
	NewBarService,
)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// this is wrong implementation,
// got inject InitializeHelloService: no provider found for golang-dependency-injection/cases.SayHello
// func InitializeHelloService() *HelloService {
// 	wire.Build(NewSayHelloImpl, NewHelloService)
// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}