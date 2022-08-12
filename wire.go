//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/google/wire"
)

func InitializeSetProvideFirstService() *FirstService {
	panic(wire.Build(SetProvideFirstService))
}

func InitializeSetProvideBindService() *Service {
	panic(wire.Build(SetProvideBindService))
}

func InitializeSetProvideStructSummary() *Summary {
	panic(wire.Build(SetProvideStructSummary))
}

func InitializeSetValueStructFirst() *First {
	panic(wire.Build(SetValueStructFirst))
}

func InitializeSetValueStructSecond() *Second {
	panic(wire.Build(SetValueStructSecond))
}

func InitializeSetValueStructSummary() *Summary {
	panic(wire.Build(SetValueStructSummary))
}

func InitializeUser(name Name, age Age, message Message) *User {
	panic(wire.Build(NewUser))
}

func InitializeDatabaseWithError(connectionString ConnectionString) (*Database, error) {
	panic(wire.Build(NewDatabaseWithError))
}

func InitializeDatabaseWithErrorAndCleanup(connectionString ConnectionString) (*Database, func(), error) {
	panic(wire.Build(NewDatabaseWithErrorAndCleanup))
}
