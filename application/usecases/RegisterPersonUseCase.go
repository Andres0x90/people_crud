package usecases

import (
	. "people_crud/application/ports"
	. "people_crud/domain"
)

type RegisterPersonUseCase struct {
	PersonRepository PersonRepositoryPort
}

func (useCase *RegisterPersonUseCase) Execute(person *Person) (*Person, error) {
	return useCase.PersonRepository.CreatePerson(person)
}
