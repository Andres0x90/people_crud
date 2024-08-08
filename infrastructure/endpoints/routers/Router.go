package routers

import (
	"github.com/gin-gonic/gin"
	"people_crud/infrastructure/config"
	"people_crud/infrastructure/driven_adapters/repository_adapters"
	"people_crud/infrastructure/endpoints/controllers"
)

type Router struct {
	injector *config.Injector
	route    *gin.Engine
}

func (router *Router) Run() {
	router.injector = &config.Injector{
		DBConn: &repository_adapters.PostgreSQLConnection{},
	}
	router.route = gin.Default()
	router.injector.InitializeDBConnection()

	PersonController := controllers.PersonController{
		RegisterPersonUseCase: router.injector.RegisterPerson(),
		FindPersonUseCase:     router.injector.FindPersonById(),
	}

	router.route.POST("/api/person", PersonController.CreatePerson)
	router.route.GET("/api/person/:document", PersonController.FindPersonById)

	router.route.Run(":8080")
}
