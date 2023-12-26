package main

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Client представляет собой клиента, который использует сервис.
type Client struct {
}

// GetProfile вызывает метод GetProfile у переданного сервиса.
func (c *Client) GetProfile(service Service) {
	service.GetProfile()
}

// Service определяет интерфейс для сервиса.
type Service interface {
	GetProfile()
}

// ClientService представляет собой конкретную реализацию сервиса.
type ClientService struct {
}

// GetProfile реализует метод интерфейса Service для ClientService.
func (s *ClientService) GetProfile() {
	fmt.Println("some result")
}

func task21() {
	client := &Client{}
	service := &ClientService{}

	// выполняет пример использования паттерна "адаптер".
	client.GetProfile(service)
}
