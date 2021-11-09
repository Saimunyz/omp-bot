package payment

import (
	"fmt"
	"strings"
)

var AllEntities = []Receipt{
	{ID: 1, Descr: "First purchase", Goods: map[string]uint64{"1-tool": 200, "2-tool": 150}},
	{ID: 2, Descr: "Second purchase", Goods: map[string]uint64{"1-tool": 1000}},
	{ID: 3, Descr: "Third purchase", Goods: map[string]uint64{"1-tool": 50, "2-tool": 111}},
	{ID: 4, Descr: "Fourth purchase", Goods: map[string]uint64{"1-tool": 200}},
	{ID: 5, Descr: "Fifth purchase", Goods: map[string]uint64{"1-tool": 547, "2-tool": 286}},
	{ID: 6, Descr: "Sixth purchase", Goods: map[string]uint64{"1-tool": 2317, "2-tool": 23286}},
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

	res := strings.Builder{}
	res.WriteString(fmt.Sprintf("%9sYour receipt\n", " "))
	res.WriteString(fmt.Sprintf("%15sId:%1d\n\n", " ", r.ID))

	for key, value := range r.Goods {
		res.WriteString(fmt.Sprintf("%10s:%10d\n", key, value))
		totalSum += value
	}
	res.WriteString(fmt.Sprintf("\n%5sTotal:%10d\n", " ", totalSum))
	res.WriteString(fmt.Sprintf("\n%20s\n", r.Descr))

	return res.String()
}
