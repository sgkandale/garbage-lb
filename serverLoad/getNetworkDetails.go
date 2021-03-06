package serverLoad

import (
	"github.com/shirou/gopsutil/v3/net"
)

func getNetworkDetails() (*ResourceUsage, error) {

	netStat, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	for _, netStats := range netStat {
		if netStats.Name == "all" {
			return &ResourceUsage{
				SentData:     convertMemoryToString(netStats.BytesSent),
				ReceivedData: convertMemoryToString(netStats.BytesRecv),
				TotalData:    convertMemoryToString(netStats.BytesSent + netStats.BytesRecv),
			}, nil
		}
	}

	return nil, nil
}
