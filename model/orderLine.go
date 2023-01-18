package model

// OrderLine is a line on an order that consists of a SKU and
// a quantity.
// OrderLine is an immutable data class with no behavior.
type OrderLine struct {
	orderID string
	sku     string
	qty     int
}
