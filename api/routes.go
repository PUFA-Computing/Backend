package api

import (
	"Backend/api/handlers/event"
	"Backend/api/handlers/files"
	"Backend/api/handlers/merch"
	"Backend/api/handlers/news"
	"Backend/api/handlers/permission"
	"Backend/api/handlers/role"
	"Backend/api/handlers/user"
	"Backend/api/middleware"
	"Backend/internal/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.Static("/public", "./public")

	userService := services.NewUserService()
	eventService := services.NewEventService()
	newsService := services.NewNewsService()
	roleService := services.NewRoleService()
	permissionService := services.NewPermissionService()
	merchService := services.NewMerchService()
	filesService := services.NewFilesService()

	userHandlers := user.NewUserHandlers(userService, permissionService)
	eventHandlers := event.NewEventHandlers(eventService, permissionService)
	newsHandlers := news.NewNewsHandler(newsService, permissionService)
	roleHandlers := role.NewRoleHandler(roleService, userService, permissionService)
	permissionHandlers := permission.NewPermissionHandler(permissionService)
	merchHandlers := merch.NewMerchHandler(merchService, permissionService)
	filesHandlers := files.NewFilesHandler(filesService, permissionService)

	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", userHandlers.RegisterUser)
		authRoutes.POST("/login", userHandlers.Login)
		authRoutes.POST("/logout", userHandlers.Logout)
		authRoutes.POST("/refresh-token", middleware.TokenMiddleware(), userHandlers.RefreshToken)
	}

	userRoutes := r.Group("/user")
	{
		userRoutes.Use(middleware.TokenMiddleware())
		userRoutes.GET("/:userID", userHandlers.GetUserByID)
		userRoutes.PUT("/edit", userHandlers.EditUser)
		userRoutes.DELETE("/delete", userHandlers.DeleteUser)
		userRoutes.GET("/list", userHandlers.ListUsers)
	}

	eventRoutes := r.Group("/event")
	{
		eventRoutes.GET("/:eventID", eventHandlers.GetEventByID)
		eventRoutes.GET("/", eventHandlers.ListEvents)
		eventRoutes.Use(middleware.TokenMiddleware())
		eventRoutes.POST("/create", eventHandlers.CreateEvent)
		eventRoutes.PUT("/:eventID/edit", eventHandlers.EditEvent)
		eventRoutes.DELETE("/:eventID/delete", eventHandlers.DeleteEvent)
		eventRoutes.POST("/:eventID/register", eventHandlers.RegisterForEvent)
		eventRoutes.GET("/:eventID/registered-users", eventHandlers.ListRegisteredUsers)
	}

	newsRoutes := r.Group("/news")
	{
		newsRoutes.GET("/", newsHandlers.ListNews)
		newsRoutes.GET("/:newsID", newsHandlers.GetNewsByID)
		newsRoutes.Use(middleware.TokenMiddleware())
		newsRoutes.POST("/create", newsHandlers.CreateNews)
		newsRoutes.PUT("/:newsID/edit", newsHandlers.EditNews)
		newsRoutes.DELETE("/:newsID/delete", newsHandlers.DeleteNews)
		newsRoutes.POST("/:newsID/like", newsHandlers.LikeNews)
	}

	roleRoutes := r.Group("/roles")
	{
		roleRoutes.Use(middleware.TokenMiddleware())
		roleRoutes.GET("/", roleHandlers.ListRoles)
		roleRoutes.POST("/create", roleHandlers.CreateRole)
		roleRoutes.GET("/:roleID", roleHandlers.GetRoleByID)
		roleRoutes.PUT("/:roleID/edit", roleHandlers.EditRole)
		roleRoutes.DELETE("/:roleID/delete", roleHandlers.DeleteRole)
		roleRoutes.POST("/:roleID/assign/:userID", roleHandlers.AssignRoleToUser)
	}
	permissionRoutes := r.Group("/permissions")
	{
		permissionRoutes.Use(middleware.TokenMiddleware())
		permissionRoutes.GET("/list", permissionHandlers.ListPermissions)
		permissionRoutes.POST("/assign/:roleID", permissionHandlers.AssignPermissionToRole)

	}
	merchRoutes := r.Group("/merch")
	{
		// Products
		merchRoutes.GET("/", merchHandlers.ListProducts)
		merchRoutes.GET("/products/:productID", merchHandlers.GetProductByID)
		merchRoutes.GET("/products/:productID/sizes", merchHandlers.GetSizeProduct)
		merchRoutes.GET("/products/:productID/colors", merchHandlers.GetColorProduct)
		merchRoutes.GET("/products/:productID/price", merchHandlers.GetProductPrice)

		// Categories
		merchRoutes.GET("/categories", merchHandlers.ListCategories)
		merchRoutes.GET("/categories/:categoryID", merchHandlers.GetCategoryByID)

		// Coupons
		//merchRoutes.GET("/coupons", merchHandlers.ListCoupons)
		//merchRoutes.GET("/coupons/:code", merchHandlers.GetCouponByCode)
		//merchRoutes.POST("/transactions/:transactionID/apply-coupon/:code", merchHandlers.ApplyCoupon)

		// Restricted Routes
		merchRoutes.Use(middleware.TokenMiddleware())

		// Products
		merchRoutes.POST("/products/create", merchHandlers.CreateProduct)
		merchRoutes.PUT("/products/:productID", merchHandlers.UpdateProduct)
		merchRoutes.DELETE("/products/:productID", merchHandlers.DeleteProduct)

		// Categories
		merchRoutes.POST("/categories/create", merchHandlers.CreateCategory)
		merchRoutes.PUT("/categories/:categoryID/edit", merchHandlers.UpdateCategory)
		merchRoutes.DELETE("/categories/:categoryID/delete", merchHandlers.DeleteCategory)

		// Size
		//merchRoutes.POST("/sizes/create", merchHandlers.CreateSize)
		//merchRoutes.PUT("/sizes/:sizeID/edit", merchHandlers.UpdateSize)
		//merchRoutes.DELETE("/sizes/:sizeID/delete", merchHandlers.DeleteSize)

		// Color
		//merchRoutes.POST("/colors/create", merchHandlers.CreateColor)
		//merchRoutes.PUT("/colors/:colorID/edit", merchHandlers.UpdateColor)
		//merchRoutes.DELETE("/colors/:colorID/delete", merchHandlers.DeleteColor)

		// Transactions
		//merchRoutes.GET("/transactions", merchHandlers.ListTransactions)
		//merchRoutes.POST("/transactions/create", merchHandlers.CreateTransaction)
		//merchRoutes.GET("/transactions/:transactionID", merchHandlers.GetTransaction)

		// Price
		//merchRoutes.POST("/prices/create", merchHandlers.CreatePrice)
		//merchRoutes.PUT("/prices/:priceID/edit", merchHandlers.UpdatePrice)
		//merchRoutes.DELETE("/prices/:priceID/delete", merchHandlers.DeletePrice)

		// Image
		//merchRoutes.POST("/products/:product_id/images", merchHandlers.UploadImageProduct)

		// Coupons
		//merchRoutes.POST("/coupons/create", merchHandlers.CreateCoupon)
		//merchRoutes.PUT("/coupons/:couponID/edit", merchHandlers.UpdateCoupon)
		//merchRoutes.DELETE("/coupons/:couponID/delete", merchHandlers.DeleteCoupon)
	}

	filesRoutes := r.Group("/files")
	{
		filesRoutes.PUT("/", filesHandlers.UploadFile)
	}
	return r
}
