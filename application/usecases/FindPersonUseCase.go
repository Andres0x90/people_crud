package usecases

import (
	"people_crud/application/ports"
	"people_crud/domain"
)

type FindPersonUseCase struct {
	PersonRepository ports.PersonRepositoryPort
}

func (useCase *FindPersonUseCase) Execute(document *string) (*domain.Person, error) {
	return useCase.PersonRepository.FindPersonById(*document)
}
