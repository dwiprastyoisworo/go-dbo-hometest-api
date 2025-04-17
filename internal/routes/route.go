package routes

import (
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/controllers"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/models"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/repositories"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/internal/services"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/config"
	"github.com/dwiprastyoisworo/go-dbo-hometest-api/lib/middlewares"
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
	r.customerRouteInit()
	r.orderRouteInit()
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

func (r *Route) customerRouteInit() {
	group := r.ctx.Group("/customer", middlewares.JWTMiddleware(r.cfg.App.JwtSecret))
	genericCustomerRepo := repositories.NewRepository[models.Customer]()
	customerService := services.NewCustomerService(genericCustomerRepo, r.validator, r.db, r.cfg)
	customerController := controllers.NewCustomerController(customerService)
	group.GET("/", customerController.GetAll)
	group.GET("/:id", customerController.GetDetail)
	group.POST("/", customerController.Create)
	group.PUT("/:id", customerController.Update)
	group.DELETE("/:id", customerController.Delete)
	group.GET("/search", customerController.Search)
}

func (r *Route) orderRouteInit() {
	group := r.ctx.Group("/order", middlewares.JWTMiddleware(r.cfg.App.JwtSecret))
	genericOrderRepo := repositories.NewRepository[models.Order]()
	genericOrderItemRepo := repositories.NewRepository[models.OrderItem]()
	genericCustomerRepo := repositories.NewRepository[models.Customer]()
	orderRepo := repositories.NewOrderRepository()
	orderService := services.NewOrderService(genericOrderRepo, genericOrderItemRepo, genericCustomerRepo, orderRepo, r.validator, r.db, r.cfg)
	orderController := controllers.NewOrderController(orderService)
	group.GET("/", orderController.GetAll)
	group.GET("/:id", orderController.GetDetail)
	group.POST("/", orderController.Create)
	group.PUT("/:id", orderController.Update)
	group.DELETE("/:id", orderController.Delete)

}
