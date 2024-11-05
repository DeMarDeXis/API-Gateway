package httphandler

import (
	"ApiGateway/internal/clients/grpc"
	"ApiGateway/internal/service"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type Handler struct {
	logg       *slog.Logger
	services   *service.Service
	grpcClient *grpc.Client
}

func NewHandler(services *service.Service, logg *slog.Logger, grpc *grpc.Client) *Handler {
	return &Handler{
		logg:       logg,
		services:   services,
		grpcClient: grpc,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	courses := router.Group("/courses")
	{
		courses.POST("/create", h.createCourse)
		courses.GET("/", h.getAllCourses)
		courses.GET("/:id", h.getCourseByID)
		courses.PUT("/:id", h.updateCourse)
		courses.DELETE("/:id", h.deleteCourse)
		//Teacher
	}

	return router
}
