package serverLoad

import (
	"github.com/shirou/gopsutil/v3/mem"
)

func getMemoryDetails() (*ResourceUsage, error) {

	memStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &ResourceUsage{
		Total:           convertMemoryToString(memStat.Total),
		InUse:           convertMemoryToString(memStat.Used),
		Available:       convertMemoryToString(memStat.Available),
		InUsePercentage: convertMemoryToPercentage(memStat.Used, memStat.Total),
	}, nil
}
