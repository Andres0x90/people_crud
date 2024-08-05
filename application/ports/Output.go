package ports

import (
	. "people_crud/domain"
)

type PersonRepositoryPort interface {
	CreatePerson(person *Person) (*Person, error)
	FindPersonById(id string) (Person, error)
	UpdatePersonById(id string) (Person, error)
	DeletePersonById(id string) error
	CountPeople() (int, error)
}
