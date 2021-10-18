package receipt

import (
	"github.com/ozonmp/omp-bot/internal/model/payment"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []payment.Receipt {
	return payment.AllEntities
}

func (s *Service) Get(idx int) (*payment.Receipt, error) {
	return &payment.AllEntities[idx], nil
}
