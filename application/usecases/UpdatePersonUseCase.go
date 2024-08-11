package usecases

import (
	"people_crud/application/ports"
	"people_crud/domain"
)

type UpdatePersonUseCase struct {
	PersonRepository ports.PersonRepositoryPort
}

func (useCase *UpdatePersonUseCase) Execute(person *domain.Person) (*domain.Person, error) {
	return useCase.PersonRepository.UpdatePersonById(person)
}
