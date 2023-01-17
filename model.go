package main

import (
	"time"
)

// OrderLine is a line on an order that consists of a SKU and
// a quantity.
// OrderLine is an immutable data class with no behavior.
type OrderLine struct {
	orderID string
	sku     string
	qty     int
}

// Batch is a collection of stock that the purchasing department orders.
// Each batch has a unique reference ID and a SKU, ETA, and available quantity.
type Batch struct {
	Ref          string
	Sku          string
	Eta          time.Time
	purchasedQty int
	allocations  []OrderLine
}

// NewBatch returns a batch.
func NewBatch(ref, sku string, qty int, eta time.Time, alloc []OrderLine) *Batch {
	return &Batch{
		Ref:          ref,
		Sku:          sku,
		Eta:          eta,
		purchasedQty: qty,
		allocations:  alloc,
	}
}

// allocate associates an OrderLine to a Batch.
func (b *Batch) allocate(line OrderLine) (*Batch, error) {
	// b.purchasedQty -= line.qty
	return b, nil
}

// deAllocate
// func (b *Batch) deAllocate(line OrderLine) error {
// 	for i, l := range b.allocations {
// 		if l == b.allocations[i] {

// 		}
// 	}
// }

// allocatedQty // getter method

// availableQty // getter method

// canAllocate verifies whether the batch has enough quantity to
// allocate the OrderLine.
func (b *Batch) canAllocate(line OrderLine) bool {
	return b.Sku == line.sku && b.purchasedQty >= line.qty
}
