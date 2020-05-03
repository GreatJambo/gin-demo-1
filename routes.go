package main

import (
	"github.com/gin-gonic/gin"
	"jkdev.cn/api/controller"
	"jkdev.cn/api/middleware"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(),middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id",categoryController.Update)//替换
	categoryRoutes.GET("/:id",categoryController.Show)
	categoryRoutes.DELETE("/:id",categoryController.Delete)

	return r
}
