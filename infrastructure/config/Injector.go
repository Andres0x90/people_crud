package config

import (
	. "people_crud/application/ports"
	. "people_crud/application/usecases"
	. "people_crud/domain"
	. "people_crud/infrastructure/driven_adapters/repository_adapters"
)

type Injector struct {
	DBConn *PostgreSQLConnection
}

func (injector *Injector) InitializeDBConnection() {
	if err := injector.DBConn.Connect(); err != nil {
		panic("Failed to connect to DB")
	}
}

func (injector *Injector) PersonRepository() PersonRepositoryPort {
	return NewPersonRepositoryAdapter(injector.DBConn)
}

func (injector *Injector) RegisterPerson() InputPort[Person, Person] {
	return &RegisterPersonUseCase{
		PersonRepository: injector.PersonRepository(),
	}
}
