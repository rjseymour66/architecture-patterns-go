package main

import (
	"errors"
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
func (b *Batch) allocate(line OrderLine) error {
	if b.canAllocate(line) {
		b.allocations = append(b.allocations, line)
		return nil
	}
	return errors.New("cannot allocate orderline")
}

// deAllocate
func (b *Batch) deAllocate(line OrderLine) error {
	for i, l := range b.allocations {
		if l == b.allocations[i] {
			b.allocations[i] = b.allocations[len(b.allocations)-1]
			b.allocations[len(b.allocations)-1] = OrderLine{}
			b.allocations = b.allocations[:len(b.allocations)-1]
			return nil
		}
	}
	return errors.New("OrderLine not found in batch")
}

// allocatedQty // getter method
func (b *Batch) allocatedQty() int {
	var sum int
	for _, val := range b.allocations {
		sum += val.qty
	}
	return sum
}

// availableQty // getter method
func (b *Batch) availableQty() int {
	return b.purchasedQty - b.allocatedQty()
}

// canAllocate verifies whether the batch has enough quantity to
// allocate the OrderLine.
func (b *Batch) canAllocate(line OrderLine) bool {
	return b.Sku == line.sku && b.purchasedQty >= line.qty
}
