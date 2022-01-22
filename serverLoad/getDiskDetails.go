package serverLoad

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

func getDiskDetails() (*ResourceUsage, error) {

	diskStat, err := disk.IOCounters()
	if err != nil {
		return nil, err
	}

	var readBytes uint64 = 0
	var writeBytes uint64 = 0
	var iops uint64 = 0

	for _, diskStats := range diskStat {
		readBytes += diskStats.ReadBytes
		writeBytes += diskStats.WriteBytes
		iops += diskStats.IopsInProgress
	}

	return &ResourceUsage{
		ReadBytes:  convertMemoryToString(readBytes),
		WriteBytes: convertMemoryToString(writeBytes),
		IOps:       fmt.Sprintf("%d", iops),
	}, nil
}
