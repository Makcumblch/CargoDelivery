package handler

import (
	"github.com/Makcumblch/CargoDelivery/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{
		projects := api.Group("/projects")
		{
			projects.POST("/", h.createProject)
			projects.GET("/", h.getAllProjects)

			project := projects.Group(":idProject", h.userAccessProject)
			{
				project.GET("/", h.getProjectById)
				projectOwner := project.Group("/", h.accessOwner)
				{
					projectOwner.PUT("/", h.updateProject)
					projectOwner.DELETE("/", h.deleteProject)
				}

				cars := project.Group("/cars")
				{
					cars.GET("/", h.getAllCars)
					cars.GET("/:id", h.getCarById)
					carsWrite := cars.Group("/", h.accessWrite)
					{
						carsWrite.POST("/", h.createCar)
						carsWrite.PUT("/:id", h.updateCar)
						carsWrite.DELETE("/:id", h.deleteCar)
					}
				}

				cargos := project.Group("/cargos")
				{
					cargos.GET("/", h.getAllCargos)
					cargos.GET("/:id", h.getCargoById)
					cargosWrite := cargos.Group("/", h.accessWrite)
					{
						cargosWrite.POST("/", h.createCargo)
						cargosWrite.PUT("/:id", h.updateCargo)
						cargosWrite.DELETE("/:id", h.deleteCargo)
					}
				}

				clients := project.Group("/clients")
				{
					clients.GET("/", h.getAllClients)
					clients.GET("/:id", h.getClientById)
					clientsWrite := clients.Group("/", h.accessWrite)
					{
						clientsWrite.POST("/", h.createClient)
						clientsWrite.PUT("/:id", h.updateClient)
						clientsWrite.DELETE("/:id", h.deleteClient)
					}

					clientsId := clients.Group(":id", h.userAccessClient)
					{
						orders := clientsId.Group("/orders")
						{
							orders.GET("/", h.getAllOrders)
							orders.GET("/:idOrder", h.getOrderById)
							ordersWrite := orders.Group("/", h.accessWrite)
							{
								ordersWrite.POST("/", h.createOrder)
								ordersWrite.PUT("/:idOrder", h.updateOrder)
								ordersWrite.DELETE("/:idOrder", h.deleteOrder)
							}
						}
					}
				}

				depo := project.Group("/depo")
				{
					depo.GET("/", h.getDepo)
					depoWrite := depo.Group("/", h.accessWrite)
					{
						depoWrite.POST("/", h.createRoute)
						depoWrite.PUT("/", h.updateDepo)
					}
				}

				routes := project.Group("/routes")
				{
					routes.GET("/")
					routes.GET("/:idRoute")
					routesWrite := routes.Group("/", h.accessWrite)
					{
						routesWrite.POST("/", h.createDepo)
						routesWrite.DELETE("/:idRoute")
					}
				}
			}
		}
	}

	return router
}
