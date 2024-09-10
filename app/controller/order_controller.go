package controller

import (
	"assignment-2/app/model"
	"assignment-2/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.orderService.CreateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"data": order})
}

func (c *OrderController) GetAllOrders(ctx *gin.Context) {
	orders, err := c.orderService.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}

func (c *OrderController) UpdateOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.orderService.UpdateOrder(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Update data success", "success": true, "data": order})
}

func (c *OrderController) DeleteOrder(ctx *gin.Context) {
	orderID := ctx.Param("id") // Get the order ID from the URL

	// Convert string ID to uint
	id, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	// Check if the order exists
	order, err := c.orderService.GetOrderById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// If the order exists, proceed to delete it
	if err := c.orderService.DeleteOrder(order.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Data Success", "success": true})
}

// func (c *OrderController) DeleteOrder(ctx *gin.Context) {
// 	orderID := ctx.Param("id")

// 	id, err := strconv.ParseUint(orderID, 10, 32)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
// 		return
// 	}

// 	order, err := c.orderService.GetOrderById(uint(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
// 		return
// 	}

// 	if err := c.orderService.DeleteOrder(order.ID); err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{"message": "Delete Data Success", "success": true})
// }

func (c *OrderController) GetOrderById(ctx *gin.Context) {
	orderID := ctx.Param("id")

	id, err := strconv.ParseUint(orderID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.orderService.GetOrderById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}
