package model

import (
	"testing"
	"time"
)

/*
	type Batch struct {
		Ref          string
		Sku          string
		Eta          time.Time
		purchasedQty int
		allocations  []OrderLine
	}
*/
func TestAllocationService(t *testing.T) {
	testCases := []struct {
		name    string
		batches []Batch
		line    OrderLine
		exp     string
	}{
		{
			name: "Prefers current batches to shipments",
			batches: []Batch{
				{"in-stock-batch", "RETRO-CLOCK", time.Time{}, 100, nil},
				{"shipment-batch", "RETRO-CLOCK", time.Now().Add(time.Hour * 24), 100, nil},
				{"dummy-batch", "RETRO-CLOCK", time.Now().Add(time.Hour * 48), 100, nil},
			},
			line: OrderLine{"oref", "RETRO-CLOCK", 10},
			exp:  "in-stock-batch",
		},
		{
			name: "Prefers earlier batches",
			batches: []Batch{
				{"speedy-batch", "MINIMALIST-SPOON", time.Now(), 100, nil},
				{"normal-batch", "MINIMALIST-SPOON", time.Now().Add(time.Hour * 24), 100, nil},
				{"slow-batch", "MINIMALIST-SPOON", time.Now().Add(time.Hour * 48), 100, nil},
			},
			line: OrderLine{"order1", "MINIMALIST-SPOON", 10},
			exp:  "speedy-batch",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := allocationService(tc.line, tc.batches)

			if got != tc.exp {
				t.Errorf("Expected %v, got %v", tc.exp, got)
			}
		})
	}
}
