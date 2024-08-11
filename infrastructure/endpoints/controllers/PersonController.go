package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"people_crud/application/ports"
	"people_crud/domain"
	"people_crud/infrastructure/endpoints/dtos"
)

type PersonController struct {
	RegisterPersonUseCase ports.InputPort[domain.Person, domain.Person]
	FindPersonUseCase     ports.InputPort[string, domain.Person]
	UpdatePersonUseCase   ports.InputPort[domain.Person, domain.Person]
}

func (controller *PersonController) CreatePerson(c *gin.Context) {
	var personDTO dtos.PersonDTO

	if err := c.ShouldBindJSON(&personDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personCreated, err := controller.RegisterPersonUseCase.Execute(dtos.MapDtoToPerson(&personDTO))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, personCreated)
}

func (controller *PersonController) FindPersonById(c *gin.Context) {
	document := c.Param("document")
	person, err := controller.FindPersonUseCase.Execute(&document)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.MapPersonToDto(person))
}

func (controller *PersonController) UpdatePerson(c *gin.Context) {
	var personDTO dtos.PersonDTO

	if err := c.ShouldBindJSON(&personDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	personUpdated, err := controller.UpdatePersonUseCase.Execute(dtos.MapDtoToPerson(&personDTO))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.MapPersonToDto(personUpdated))
}
