package handler

import (
	"WhyAi/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("auth")
			{

				auth.POST("/sign-up", h.signUp)
				auth.POST("/sign-in", h.signIn)
				//auth.POST("/reset-request", h.passwordResetRequest)
				//auth.POST("/reset-confirm/", h.passwordResetConfirm)
			}
			user := v1.Group("/user", h.userIdentity)
			{
				user.GET("/info", h.GetUserInfo)
				user.PUT("/update", h.UpdateUserInfo)
			}
			theory := v1.Group("/theory")
			{
				theory.GET("/:id", h.SendTheory) //Получение теории
				theory.POST("/chat/:id")         //Чат по заданию
			}
			fact := v1.Group("/fact")
			{
				fact.GET("/") //Получить случайный лайфхак, ну или массив лайфхаков, чтобы уменьшить нагрузку
			}
		}

	}

	return router
}
