package service

import (
	"assignment-2/app/model"
	"assignment-2/app/repository"
)

type OrderService interface {
	CreateOrder(order *model.Order) error
	GetAllOrders() ([]model.Order, error)
	UpdateOrder(order *model.Order) error
	DeleteOrder(orderID uint) error
	GetOrderById(id uint) (*model.Order, error)
}

type orderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) OrderService {
	return &orderService{orderRepo}
}

func (s *orderService) CreateOrder(order *model.Order) error {
	return s.orderRepo.Create(order)
}

func (s *orderService) GetAllOrders() ([]model.Order, error) {
	return s.orderRepo.FindAll()
}

func (s *orderService) UpdateOrder(order *model.Order) error {
	return s.orderRepo.Update(order)
}

func (s *orderService) DeleteOrder(orderID uint) error {
	return s.orderRepo.Delete(orderID)
}

func (s *orderService) GetOrderById(id uint) (*model.Order, error) {
	return s.orderRepo.GetOrderById(id)
}
