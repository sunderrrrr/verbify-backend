package handler

import (
	"WhyAi/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(frontendUrl string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.RedirectTrailingSlash = false
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendUrl},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Baggage", "Sentry-Trace"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("auth") // Мутки с авторизацией
			{

				auth.POST("/sign-up", h.signUp)
				auth.POST("/sign-in", h.signIn)
				auth.POST("/forgot", h.SendResetRequest)
				auth.POST("/reset", h.UpdatePassword)
			}

			user := v1.Group("/user", h.userIdentity) // Инфа пользователя, подписки
			{
				user.GET("/info", h.GetUserInfo)
				user.PUT("/update", h.UpdateUserInfo)
			}

			theory := v1.Group("/theory", h.userIdentity)
			{
				theory.GET("/:id", h.SendTheory)           //Получение теории
				theory.POST("/:id/chat", h.SendMessage)    //Сообщение по заданию
				theory.GET("/:id/chat", h.GetOrCreateChat) // Получить чат
				theory.DELETE("/:id/chat", h.ClearContext) //Стереть контекст
			}

			essay := v1.Group("/essay", h.userIdentity)
			{
				essay.GET("/themes", h.GetEssayTasks)
				essay.POST("/", h.SendEssay)
			}

			fact := v1.Group("/fact")
			fact.Use(cors.New(cors.Config{
				AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Baggage", "Sentry-Trace"},
				ExposeHeaders:    []string{"Content-Length", "Authorization"},
				AllowCredentials: true,
				AllowOrigins:     []string{frontendUrl},
				MaxAge:           12 * time.Hour,
			}))
			{
				fact.GET("", h.GetFact) //Получить случайный лайфхак, ну или массив лайфхаков, чтобы уменьшить нагрузку
			}
			admin := v1.Group("/admin", h.userIdentity, h.adminIdentity) //Админка
			{
				admin.GET("/")
			}
		}

	}

	return router
}
