package main

import . "people_crud/infrastructure/driven_adapters/repository_adapters"

func main() {
	conn := DBConnection{}
	conn.Connect()
	conn.DB.AutoMigrate(&PersonModel{}, &CompanyModel{}, &SkillModel{}, &PersonSkillsModel{}, &PayrollModel{})
}
