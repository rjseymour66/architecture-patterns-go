package main

import (
	"testing"
	"time"
)

func TestAllocate(t *testing.T) {
	batch := NewBatch("batch-001", "SMALL-TABLE", 20, time.Now(), nil)
	line := OrderLine{
		orderID: "batch-001",
		sku:     "SMALL-TABLE",
		qty:     2,
	}

	err := batch.allocate(line)
	if err != nil {
		t.Fatal(err)
	}

	exp := 18

	if batch.purchasedQty != exp {
		t.Fatalf("Expected %d, got %d instead", exp, batch.purchasedQty)
	}

}

func TestCanAllocate(t *testing.T) {
	batch := NewBatch("batch-001", "", 0, time.Now(), nil)
	line := OrderLine{
		orderID: "order-123",
		sku:     "",
		qty:     0,
	}
	testCases := []struct {
		name string
		bSku string
		lSku string
		bQty int
		lQty int
		exp  bool
	}{
		{
			name: "Available greater than required",
			bSku: "ELEGANT-LAMP",
			lSku: "ELEGANT-LAMP",
			bQty: 20,
			lQty: 2,
			exp:  true,
		},
		{
			name: "Available smaller than required",
			bSku: "ELEGANT-LAMP",
			lSku: "ELEGANT-LAMP",
			bQty: 2,
			lQty: 20,
			exp:  false,
		},
		{
			name: "Available equal to required",
			bSku: "ELEGANT-LAMP",
			lSku: "ELEGANT-LAMP",
			bQty: 2,
			lQty: 2,
			exp:  true,
		},
		{
			name: "SKUs do not match",
			bSku: "UNCOMFORTABLE-CHAIR",
			lSku: "EXPENSIVE-TOASTER",
			bQty: 100,
			lQty: 10,
			exp:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			batch.Sku = tc.bSku
			batch.purchasedQty = tc.bQty
			line.sku = tc.lSku
			line.qty = tc.lQty

			got := batch.canAllocate(line)

			if got != tc.exp {
				t.Errorf("Expected %t, got %t", tc.exp, got)
			}
		})
	}
}

// func TestDeallocate(t *testing.T) {

// }
