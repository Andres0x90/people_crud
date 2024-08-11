package usecases

import (
	"errors"
	"people_crud/application/ports"
	"people_crud/domain"
)

type FindPersonUseCase struct {
	PersonRepository ports.PersonRepositoryPort
}

func (useCase *FindPersonUseCase) Execute(document *string) (*domain.Person, error) {
	personFoundChannel := make(chan *domain.Person)
	go useCase.PersonRepository.FindPersonById(*document, personFoundChannel)
	person := <-personFoundChannel

	if person.Identification == "" {
		return nil, errors.New("Person not found")
	}
	return person, nil
}
