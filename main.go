package main

import (
	"assignment-2/app/controller"
	"assignment-2/app/model"
	"assignment-2/app/repository"
	"assignment-2/app/service"
	"assignment-2/config"

	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Initialize Database
	config.InitDB()

	InitDB()

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

// const (
// 	host     = "localhost"
// 	port     = "5432"
// 	user     = "jefrys"
// 	password = "postgres"
// 	dbname   = "assignment-2"
// )

var db *gorm.DB

func InitDB() {
	// config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)
	dsn := "host=localhost dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// AutoMigrate models
	db.Debug().AutoMigrate(&model.Order{}, &model.Item{})
}
