package dtos

import "people_crud/domain"

type PersonDTO struct {
	Type           string     `json:"type" binding:"required"`
	Identification string     `json:"identification" binding:"required"`
	Name           string     `json:"name" binding:"required"`
	Age            int8       `json:"age" binding:"required"`
	Company        CompanyDTO `json:"company" binding:"required"`
	Skills         []SkillDTO `json:"skills" binding:"required"`
}

type CompanyDTO struct {
	NIT         string `json:"nit" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type SkillDTO struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func MapDtoToPerson(dto *PersonDTO) *domain.Person {
	personMapped := domain.Person{
		Type:           domain.IDType(dto.Type),
		Identification: dto.Identification,
		Name:           dto.Name,
		Age:            dto.Age,
		Company: domain.Company{
			NIT:         dto.Company.NIT,
			Name:        dto.Company.Name,
			Description: dto.Company.Description,
		},
	}

	for _, skill := range dto.Skills {
		personMapped.Skills = append(personMapped.Skills, domain.Skill{
			ID:          skill.ID,
			Name:        skill.Name,
			Description: skill.Description,
		})
	}

	return &personMapped
}
