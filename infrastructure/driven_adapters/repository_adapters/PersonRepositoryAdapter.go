package repository_adapters

import (
	"errors"
	. "people_crud/domain"
)

type PersonModel struct {
	Type           string
	Identification string `gorm:"primaryKey"`
	Name           string
	Age            int8
	CompanyNit     string
	Company        CompanyModel   `gorm:"foreignKey:CompanyNit; references:nit"`
	Skills         []SkillModel   `gorm:"many2many:person_skills; foreignKey:Identification; joinForeignKey: Document; references:ID;  joinReferences: SkillId;"`
	Payrolls       []PayrollModel `gorm:"foreignKey:Document; references:Identification"`
}

func (p *PersonModel) TableName() string {
	return "people"
}

type PersonRepositoryAdapter struct {
	*PostgreSQLConnection
}

func NewPersonRepositoryAdapter(db *PostgreSQLConnection) *PersonRepositoryAdapter {
	return &PersonRepositoryAdapter{PostgreSQLConnection: db}
}

func (personRepository *PersonRepositoryAdapter) CreatePerson(person *Person) (*Person, error) {
	personMapped := mapEntityToPersonModel(person)

	result := personRepository.DB.Create(personMapped)

	if result.Error != nil {
		return &Person{}, result.Error
	}

	return person, nil
}

func mapEntityToPersonModel(person *Person) *PersonModel {
	personMapped := PersonModel{
		Type:           string(person.Type),
		Identification: person.Identification,
		Name:           person.Name,
		Age:            person.Age,
		Company: CompanyModel{
			NIT:         person.Company.NIT,
			Name:        person.Company.Name,
			Description: person.Company.Description,
		},
	}

	personMapped.Skills = make([]SkillModel, 0)

	for _, skill := range person.Skills {
		personMapped.Skills = append(personMapped.Skills, SkillModel{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
		})
	}
	return &personMapped
}

func mapModelToPersonEntity(personModel *PersonModel) *Person {
	personEntity := Person{
		Type:           IDType(personModel.Type),
		Identification: personModel.Identification,
		Name:           personModel.Name,
		Age:            personModel.Age,
		Company: Company{
			NIT:         personModel.Company.NIT,
			Name:        personModel.Company.Name,
			Description: personModel.Company.Description,
		},
	}

	for _, skill := range personModel.Skills {
		personEntity.Skills = append(personEntity.Skills, Skill{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
		})
	}

	return &personEntity
}

func (personRepository *PersonRepositoryAdapter) FindPersonById(id string, userFoundChannel chan *Person) chan *Person {
	person := PersonModel{Company: CompanyModel{}}
	personRepository.DB.Preload("Skills").
		Joins("Company").
		First(&person, id)

	if person.Identification == "" {
		userFoundChannel <- &Person{}
		return userFoundChannel
	}

	userFoundChannel <- mapModelToPersonEntity(&person)

	return userFoundChannel
}

func (personRepository *PersonRepositoryAdapter) UpdatePersonById(person *Person) (*Person, error) {
	userFoundChannel := make(chan *Person)
	go personRepository.FindPersonById(person.Identification, userFoundChannel)
	personFound := <-userFoundChannel

	if personFound.Identification == "" {
		return nil, errors.New("person not found")
	}

	personRepository.DB.Updates(mapEntityToPersonModel(person))
	return person, nil
}

func (personRepository *PersonRepositoryAdapter) DeletePersonById(id string) error {
	//TODO implement me
	panic("implement me")
}

func (personRepository *PersonRepositoryAdapter) CountPeople() (int, error) {
	//TODO implement me
	panic("implement me")
}
