package repository_adapters

type SkillModel struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
}

func (s *SkillModel) TableName() string {
	return "skills"
}
