package serverLoad

import (
	"github.com/shirou/gopsutil/v3/disk"
)

func getStorageDetails() (*ResourceUsage, error) {

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	return &ResourceUsage{
		Total:           convertMemoryToString(diskStat.Total),
		InUse:           convertMemoryToString(diskStat.Used),
		Available:       convertMemoryToString(diskStat.Free),
		InUsePercentage: convertMemoryToPercentage(diskStat.Used, diskStat.Total),
	}, nil
}
