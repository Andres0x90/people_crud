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
}

func (controller *PersonController) CreatePerson(c *gin.Context) {
	var personDTO dtos.PersonDTO

	if err := c.ShouldBindJSON(&personDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	personCreated, err := controller.RegisterPersonUseCase.Execute(dtos.MapDtoToPerson(&personDTO))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, personCreated)
}
