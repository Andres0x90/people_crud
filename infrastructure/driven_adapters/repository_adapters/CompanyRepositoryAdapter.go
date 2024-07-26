package repository_adapters

type CompanyModel struct {
	NIT         string `gorm:"primaryKey"`
	Name        string
	Description string
}

func (c *CompanyModel) TableName() string {
	return "companies"
}
