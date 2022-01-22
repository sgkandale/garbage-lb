package serverLoad

import (
	"runtime"
)

func getRuntimeDetails() (*ResourceUsage, error) {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &ResourceUsage{
		Total:           convertMemoryToString(m.Sys),
		InUse:           convertMemoryToString(m.Alloc),
		InUsePercentage: convertMemoryToPercentage(m.Alloc, m.Sys),
	}, nil
}
