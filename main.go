package main

import (
	"ddd_example/app/controller"
	"ddd_example/app/repository"
	"ddd_example/app/service"
	"ddd_example/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	config.InitDB()

	// Initialize Repository
	orderRepo := repository.NewOrderRepository(config.GetDB())

	// Initialize Service
	orderService := service.NewOrderService(orderRepo)

	// Initialize Controller
	orderController := controller.NewOrderController(orderService)

	// Initialize Router
	router := gin.Default()

	// Define Routes
	router.POST("/orders", orderController.CreateOrder)
	router.GET("/orders", orderController.GetAllOrders)
	router.PUT("/orders/:id", orderController.UpdateOrder)
	router.DELETE("/orders/:id", orderController.DeleteOrder)

	// Start Server
	router.Run(":8080")
}

func InitDB() {
	dsn := "host=localhost user=youruser password=yourpass dbname=yourdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// AutoMigrate models
	DB.AutoMigrate(&model.Order{}, &model.Item{})
}
