package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"
	"github.com/gin-gonic/gin"
)

type FieldRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IFieldRoute interface {
	Run()
}

func NewFieldRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IFieldRoute {
	return &FieldRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (f *FieldRoute) Run() {
	group := f.group.Group("/field")
	group.GET("", middlewares.AuthenticateWithoutToken(), f.controller.GetField().GetAllWithoutPagination)
	group.GET("/:uuid", middlewares.AuthenticateWithoutToken(), f.controller.GetField().GetByUUID)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client),
		f.controller.GetField().GetAllWithPagination)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetField().Create)
	group.PUT("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetField().Update)
	group.DELETE("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetField().Delete)
}
