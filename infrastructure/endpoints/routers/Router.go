package routers

import (
	"github.com/gin-gonic/gin"
	"people_crud/infrastructure/config"
	"people_crud/infrastructure/driven_adapters/repository_adapters"
	"people_crud/infrastructure/endpoints/controllers"
)

type Router struct {
	injector         *config.Injector
	route            *gin.Engine
	personController *controllers.PersonController
}

func (router *Router) Run() {
	router.injector = &config.Injector{
		DBConn: &repository_adapters.PostgreSQLConnection{},
	}

	router.route = gin.Default()
	router.injector.InitializeDBConnection()
	router.initializeControllers()
	router.initializeRoutes()

	router.route.Run(":8080")
}

func (router *Router) initializeControllers() {
	router.personController = &controllers.PersonController{
		RegisterPersonUseCase: router.injector.RegisterPerson(),
		FindPersonUseCase:     router.injector.FindPersonById(),
		UpdatePersonUseCase:   router.injector.UpdatePerson(),
	}
}

func (router *Router) initializeRoutes() {
	router.route.POST("/api/person", router.personController.CreatePerson)
	router.route.GET("/api/person/:document", router.personController.FindPersonById)
	router.route.PUT("/api/person", router.personController.UpdatePerson)
}
