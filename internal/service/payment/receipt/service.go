package receipt

import (
	"errors"
	"sort"

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

var (
	ErrNotEnougthElem = errors.New("not enought elements")
	ErrWrongID        = errors.New("there is no receipt with such ID")
	ErrIdIsUsed       = errors.New("this ID is already used")
	ErrWrongIdOrUsed  = errors.New("there is no receipt with such ID: %d or ID alredy exist")
)

func (d *DummyReceiptService) Describe(
	receiptID uint64) (*payment.Receipt, error) {
	for _, value := range payment.AllEntities {
		if value.ID == receiptID {
			return &value, nil
		}
	}
	return nil, ErrWrongID
}

func (d *DummyReceiptService) List(
	cursor uint64,
	limit uint64) ([]payment.Receipt, error) {
	end := cursor * limit
	start := end - limit
	lenght := d.Len()
	if end > lenght {
		if start >= lenght {
			return payment.AllEntities[:], ErrNotEnougthElem
		} else {
			return payment.AllEntities[start:], nil
		}
	} else {
		return payment.AllEntities[start:end], nil
	}

}

func (d *DummyReceiptService) Create(
	newReceipt payment.Receipt) (uint64, error) {
	if d.Contains(newReceipt.ID) {
		return 0, ErrWrongID
	}
	payment.AllEntities = append(payment.AllEntities, newReceipt)
	return newReceipt.ID, nil
}

func (d *DummyReceiptService) Update(
	receiptID uint64,
	receipt payment.Receipt) error {
	for i, value := range payment.AllEntities {
		if value.ID == receiptID &&
			(value.ID == receipt.ID || !d.Contains(receipt.ID)) {
			payment.AllEntities[i] = receipt
			return nil
		}
	}
	return ErrWrongIdOrUsed
}

func (d *DummyReceiptService) Remove(receiptID uint64) (bool, error) {
	for i, value := range payment.AllEntities {
		if value.ID == receiptID {
			payment.AllEntities = append(payment.AllEntities[:i], payment.AllEntities[i+1:]...)
			return true, nil
		}
	}
	return false, ErrWrongID
}

func NewDummyReceiptService() *DummyReceiptService {
	return &DummyReceiptService{}
}

func (d *DummyReceiptService) Len() uint64 {
	return uint64(len(payment.AllEntities))
}

func (d *DummyReceiptService) AvailIndex() []uint64 {
	var indexes []uint64
	for _, value := range payment.AllEntities {
		indexes = append(indexes, value.ID)
	}
	sort.Slice(indexes, func(i, j int) bool { return indexes[i] < indexes[j] })
	return indexes
}

func (d *DummyReceiptService) Contains(idx uint64) bool {
	for _, value := range payment.AllEntities {
		if value.ID == idx {
			return true
		}
	}
	return false
}
