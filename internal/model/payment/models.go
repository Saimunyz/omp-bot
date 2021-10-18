package payment

import "fmt"

var AllEntities = []Receipt{
	{ID: 1, Descr: "First purchase", Goods: map[string]uint64{"1-tool": 200, "2-tool": 150}},
	{ID: 2, Descr: "Second purchase", Goods: map[string]uint64{"1-tool": 1000}},
	{ID: 3, Descr: "Third purchase", Goods: map[string]uint64{"1-tool": 50, "2-tool": 111}},
	{ID: 4, Descr: "Fourth purchase", Goods: map[string]uint64{"1-tool": 200}},
	{ID: 5, Descr: "Fifth purchase", Goods: map[string]uint64{"1-tool": 547, "2-tool": 286}},
}

type Receipt struct {
	ID    uint64
	Descr string
	Goods map[string]uint64
}

func NewReceipt(ID uint64, Descr string, Goods map[string]uint64) *Receipt {
	return &Receipt{
		ID:    ID,
		Descr: Descr,
		Goods: Goods,
	}
}

func (r *Receipt) String() string {
	var totalSum uint64

	res := fmt.Sprintf("%9sYour receipt\n", " ")
	res += fmt.Sprintf("%15sId:%1d\n\n", " ", r.ID)

	for key, value := range r.Goods {
		res += fmt.Sprintf("%10s:%10d\n", key, value)
		totalSum += value
	}
	res += "\n"
	res += fmt.Sprintf("%5sTotal:%10d\n", " ", totalSum)
	res += "\n"
	res += fmt.Sprintf("%20s\n", r.Descr)

	return res
}
