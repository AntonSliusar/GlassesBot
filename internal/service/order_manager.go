package service

import (
	"fmt"
	"glassesbot/internal/domain"
	"glassesbot/internal/repository"
	"sync"
	"time"
)



type OrderManager struct {
	Orders map[int64] *domain.Order
	repository *repository.OrderRepository
	mutex sync.RWMutex

}

func NewOrderManager(repo *repository.OrderRepository) *OrderManager {
	return &OrderManager{
		Orders: make(map[int64]*domain.Order),
		repository: repo,
	}
}

func (m *OrderManager) CreateOrder() int64 {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	randomID := int64(time.Now().Unix() / 10)
	
	order := domain.NewOrder()
	m.Orders[randomID] = order

	return randomID
}	

func (m *OrderManager) PauseOrder(id int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	order, exist := m.Orders[id]
	if !exist {
		return fmt.Errorf("Order with ID %d not found", id)
	}
	if order.Status != domain.STATUS_IN_WORK {
		return fmt.Errorf("Order with Id %d is not in work", id)
	}
	order.Pause()
	return nil
}

func (m *OrderManager) ResumeOrder(id int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	order, exist := m.Orders[id]
	if !exist {
		return fmt.Errorf("Order with ID %d not found", id)
	}
	if order.Status != domain.STATUS_PAUSE {
		return fmt.Errorf("Order with ID %d is not paused", id)
	}
	order.Resume()
	return nil
}

func (m *OrderManager) FinishOrder(id int64) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	order, exist := m.Orders[id]
	if !exist {
		fmt.Printf("Order with ID %d not found", id)
		
	}
	order.TotalWorkinTime()

	if err := m.repository.Save(order); err != nil {
		fmt.Printf("Failde to save order: &w", err)
		return err
	}
	delete(m.Orders, id)
	return nil
}

func (m *OrderManager) GetAllOrders() map[int64]*domain.Order {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	
	return m.Orders
}