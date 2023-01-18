package model

import (
	"sort"
)

// allocationService figures out which batches can allocate the line, sorts the batches
// by date, allocates the line, then returns the batch reference.
func allocationService(line OrderLine, batches []Batch) string {

	var canAlloc []Batch
	// if canAllocate
	for _, b := range batches {
		if b.canAllocate(line) {
			canAlloc = append(canAlloc, b)
		}
	}
	// sort all that canAllocate
	sort.Sort(ByDate(canAlloc))
	// allocate with the first in the b
	canAlloc[0].allocate(line)
	// return ref
	return canAlloc[0].Ref
}

// ByDate implements Sort for allocationService
type ByDate []Batch

// Len returns the length of ByDate
func (bd ByDate) Len() int {
	return len(bd)
}

// Swap switches values
func (bd ByDate) Swap(i, j int) {
	bd[i], bd[j] = bd[j], bd[i]
}

func (bd ByDate) Less(i, j int) bool {
	return bd[i].Eta.Before(bd[j].Eta)
}
