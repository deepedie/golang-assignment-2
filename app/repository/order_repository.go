package repository

import (
	"assignment-2/app/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *model.Order) error
	FindAll() ([]model.Order, error)
	Update(order *model.Order) error
	Delete(orderID uint) error
	GetOrderById(id uint) (*model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Preload("Items").Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Update(order *model.Order) error {
	return r.db.Save(order).Error
}

func (r *orderRepository) Delete(orderID uint) error {
	return r.db.Where("id = ?", orderID).Delete(&model.Order{}).Error
}

func (r *orderRepository) GetOrderById(id uint) (*model.Order, error) {
	var order model.Order
	if err := r.db.First(&order, id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
