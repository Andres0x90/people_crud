package repository_adapters

import (
	. "people_crud/domain"
)

type PersonModel struct {
	Type           string
	Identification string `gorm:"primaryKey"`
	Name           string
	Age            int8
	CompanyNit     string
	Company        CompanyModel   `gorm:"foreignKey:CompanyNit; references:nit"`
	Skills         []SkillModel   `gorm:"many2many:person_skills;"`
	Payrolls       []PayrollModel `gorm:"foreignKey:Document; references:Identification"`
}

type PersonSkillsModel struct {
	Document  string      `gorm:"primaryKey"`
	PersonRef PersonModel `gorm:"foreignKey:Document; references:Identification"`
	SkillId   string      `gorm:"primaryKey"`
	SkillRef  SkillModel  `gorm:"foreignKey:SkillId; references:ID"`
}

func (p *PersonSkillsModel) TableName() string {
	return "person_skills"
}

func (p *PersonModel) TableName() string {
	return "people"
}

type PersonRepositoryAdapter struct {
	*DBConnection
}

func (personRepository *PersonRepositoryAdapter) CreatePerson(person *Person) (*Person, error) {
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

	person.Skills = make([]Skill, 0)

	for _, skill := range person.Skills {
		personMapped.Skills = append(personMapped.Skills, SkillModel{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
		})
	}

	result := personRepository.DB.Create(&personMapped)

	if result.Error != nil {
		return &Person{}, result.Error
	}

	return person, nil
}
