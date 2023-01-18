package main

import (
	"testing"
	"time"
)

func TestCanAllocate(t *testing.T) {
	batch := NewBatch("batch-001", "", time.Now(), 0, nil)
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

// func TestIdempotency(t *testing.T) {
// 	batch := NewBatch("batch-001", "", time.Now(), 0, nil)
// 	line := OrderLine{
// 		orderID: "order-123",
// 		sku:     "",
// 		qty:     0,
// 	}
// 	testCases := []struct {
// 		bSku string
// 		lSku string
// 		bQty int
// 		lQty int
// 		exp  bool
// 	}{
// 		{
// 			bSku: "ELEGANT-LAMP",
// 			lSku: "ELEGANT-LAMP",
// 			bQty: 20,
// 			lQty: 2,
// 			exp:  true,
// 		},
// 		{
// 			bSku: "ELEGANT-LAMP",
// 			lSku: "ELEGANT-LAMP",
// 			bQty: 20,
// 			lQty: 2,
// 			exp:  false,
// 		},
// 	}

// 	for _, tc := range testCases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			batch.Sku = tc.bSku
// 			batch.purchasedQty = tc.bQty
// 			line.sku = tc.lSku
// 			line.qty = tc.lQty

// 			got := batch.canAllocate(line)

// 			if got != tc.exp {
// 				t.Errorf("Expected %t, got %t", tc.exp, got)
// 			}
// 		})
// 	}
// }

func TestAllocatedQty(t *testing.T) {
	batch := NewBatch("batch-001", "TEST1", time.Now(), 1, nil)

	line1 := OrderLine{
		orderID: "one",
		sku:     "TEST1",
		qty:     1,
	}

	batch.allocate(line1)

	exp := 1
	got := batch.allocations
	if exp != len(got) {
		t.Errorf("Expected %v, got %v", exp, got)
	}
}

func TestDeallocate(t *testing.T) {
	batch := NewBatch("batch-001", "TEST1", time.Now(), 1, nil)

	line1 := OrderLine{
		orderID: "one",
		sku:     "TEST1",
		qty:     1,
	}

	batch.allocate(line1)

	if len(batch.allocations) != 1 {
		t.Errorf("Did not allocate")
	}

	batch.deallocate(line1)

	expLength := 0
	gotLength := len(batch.allocations)

	if expLength != gotLength {
		t.Errorf("Expected %v allocations, got %v", expLength, gotLength)
	}

	for _, line := range batch.allocations {
		if line == line1 {
			t.Errorf("line2 was not deallocated")
		}
	}
}
