package receipt

import (
	"fmt"

	"github.com/ozonmp/omp-bot/internal/model/payment"
)

type ReceiptService interface {
	Describe(receiptID uint64) (*payment.Receipt, error)
	List(cursor uint64, limit uint64) ([]payment.Receipt, error)
	Create(newReceipt payment.Receipt) (uint64, error)
	Update(receiptID uint64, receipt payment.Receipt) error
	Remove(receiptID uint64) (bool, error)
}

type DummyReceiptService struct{}

func (d *DummyReceiptService) Describe(
	receiptID uint64) (*payment.Receipt, error) {
	for _, value := range payment.AllEntities {
		if value.ID == receiptID {
			return &value, nil
		}
	}
	return nil, fmt.Errorf("there is no receipt with such ID: %d", receiptID)
}

// Move all func in commander ?
func (d *DummyReceiptService) List(
	cursor uint64,
	limit uint64) ([]payment.Receipt, error) {
	end := cursor * limit
	start := end - limit
	lenght := uint64(len(payment.AllEntities))
	if end > lenght {
		if start > lenght {
			return payment.AllEntities[:], fmt.Errorf("there is no enought elements")
		} else {
			return payment.AllEntities[start:], nil
		}
	} else {
		return payment.AllEntities[start:end], nil
	}

}

func (d *DummyReceiptService) Create(
	newReceipt payment.Receipt) (uint64, error) {
	last := payment.AllEntities[len(payment.AllEntities)-1]
	if newReceipt.ID <= last.ID {
		return 0,
			fmt.Errorf("cannot create new receipt with such ID: %d", newReceipt.ID)
	}
	payment.AllEntities = append(payment.AllEntities, newReceipt)
	return newReceipt.ID, nil
}

func (d *DummyReceiptService) Update(
	receiptID uint64,
	receipt payment.Receipt) error {
	for i, valud := range payment.AllEntities {
		if valud.ID == receiptID {
			payment.AllEntities[i] = receipt
			return nil
		}
	}
	return fmt.Errorf("there is no receipt with such ID: %d", receiptID)
}

func (d *DummyReceiptService) Remove(receiptID uint64) (bool, error) {
	for i, valud := range payment.AllEntities {
		if valud.ID == receiptID {
			if i == len(payment.AllEntities)-1 {
				payment.AllEntities = payment.AllEntities[:i]
			} else {
				payment.AllEntities = append(payment.AllEntities[:i], payment.AllEntities[i+1:]...)
			}
			return true, nil
		}
	}
	return false, fmt.Errorf("there is no receipt with such ID: %d", receiptID)
}

func NewDummyReceiptService() *DummyReceiptService {
	return &DummyReceiptService{}
}
