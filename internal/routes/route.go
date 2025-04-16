package routes

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/controllers"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/repositories"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/services"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Route struct {
	db        *gorm.DB
	ctx       *gin.Engine
	cfg       config.AppConfig
	validator *validator.Validate
}

func NewRoute(db *gorm.DB, ctx *gin.Engine, cfg config.AppConfig, validator *validator.Validate) *Route {
	return &Route{db: db, ctx: ctx, cfg: cfg, validator: validator}
}

func (r *Route) RouteInit() {
	r.userRouteInit()
}

func (r *Route) userRouteInit() {
	group := r.ctx.Group("/user")
	genericUserRepo := repositories.NewRepository[models.User]()
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(genericUserRepo, userRepo, r.validator, r.db, r.cfg)
	userController := controllers.NewUserController(userService)
	group.POST("/register", userController.Register)
	group.POST("/login", userController.Login)
}
