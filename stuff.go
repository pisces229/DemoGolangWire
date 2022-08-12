package main

import (
	"errors"
	"log"

	"github.com/google/wire"
)

// binding-interfaces

// Repository
type IRepository interface{}
type Repository struct{}
type FirstRepository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
func NewFirstRepository() *FirstRepository {
	return &FirstRepository{}
}

// Usecase
type ILogic interface{}
type Logic struct {
	Repository IRepository
}
type FirstLogic struct {
	Repository *FirstRepository
}

func NewLogic(repository IRepository) *Logic {
	return &Logic{Repository: repository}
}
func NewFirstLogic(repository *FirstRepository) *FirstLogic {
	return &FirstLogic{Repository: repository}
}

// Service
type IService interface{}
type Service struct {
	Logic ILogic
}
type FirstService struct {
	Logic *FirstLogic
}

func NewService(logic ILogic) *Service {
	return &Service{Logic: logic}
}
func NewFirstService(logic *FirstLogic) *FirstService {
	return &FirstService{Logic: logic}
}

var SetProvideFirstService = wire.NewSet(
	NewFirstService,
	NewFirstLogic,
	NewFirstRepository)

// var SetProvideBindService = wire.NewSet(
// 	NewRepository,
// 	wire.Bind(new(IRepository), new(*Repository)),
// 	NewLogic,
// 	wire.Bind(new(ILogic), new(*Logic)),
// 	NewService)

var SetProvideBindService = wire.NewSet(
	wire.NewSet(NewRepository, wire.Bind(new(IRepository), new(*Repository))),
	wire.NewSet(NewLogic, wire.Bind(new(ILogic), new(*Logic))),
	NewService)

type Name string
type Age int
type Message string

type First struct {
	name    Name
	age     Age
	message Message
}
type Second struct {
	name    Name
	age     Age
	message Message
}

func ProvideFirst() First {
	return First{"ProvideNameFirst", 1, "ProvideMessageFirst"}
}
func ProvideSecond() Second {
	return Second{"ProvideNameSecond", 1, "ProvideMessageSecond"}
}

type Summary struct {
	First  First
	Second Second
}

// var SetProvideStructSummary = wire.NewSet(
// 	ProvideFirst,
// 	ProvideSecond,
// 	wire.Struct(new(Summary), "First", "Second"))

var SetProvideStructSummary = wire.NewSet(
	ProvideFirst,
	ProvideSecond,
	wire.Struct(new(Summary), "*"))

var SetValueStructFirst = wire.NewSet(
	wire.Value(Name("ValueNameFirst")),
	wire.Value(Age(1)),
	wire.Value(Message("ValueMessageFirst")),
	wire.Struct(new(First), "*"))

var SetValueStructSecond = wire.NewSet(
	wire.Value(Name("ValueNameSecond")),
	wire.Value(Age(2)),
	wire.Value(Message("ValueMessageSecond")),
	wire.Struct(new(Second), "*"))

var SetValueStructSummary = wire.NewSet(
	wire.Value(Name("ValueNameSummary")),
	wire.Value(Age(3)),
	wire.Value(Message("ValueMessageSummary")),
	wire.Struct(new(First), "*"),
	wire.Struct(new(Second), "*"),
	wire.Struct(new(Summary), "*"))

type User struct {
	name    Name
	age     Age
	message Message
}

func NewUser(name Name, age Age, message Message) *User {
	return &User{name, age, message}
}

type ConnectionString string
type Database struct {
	ConnectionString ConnectionString
}

func NewDatabaseWithError(connectionString ConnectionString) (*Database, error) {
	if connectionString == "" {
		return nil, errors.New("miss ConnectionString")
	}
	return &Database{connectionString}, nil
}

func NewDatabaseWithErrorAndCleanup(connectionString ConnectionString) (*Database, func(), error) {
	if connectionString == "" {
		return nil, nil, errors.New("miss ConnectionString")
	}
	database := &Database{connectionString}
	cleanup := func() {
		log.Println("cleanup")
	}
	return database, cleanup, nil
}
