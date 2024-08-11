package ports

import (
	. "people_crud/domain"
)

type PersonRepositoryPort interface {
	CreatePerson(person *Person) (*Person, error)
	FindPersonById(id string, userFoundChannel chan *Person) chan *Person
	UpdatePersonById(person *Person) (*Person, error)
	DeletePersonById(id string) error
	CountPeople() (int, error)
}
