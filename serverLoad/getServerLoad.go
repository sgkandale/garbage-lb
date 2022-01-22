package serverLoad

import (
	"runtime"
)

func GetServerLoad() (*ServerLoad, error) {
	serverLoad := &ServerLoad{}

	// Program Stats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	serverLoad.Runtime.InUse = convertMemoryToString(m.Sys)

	// Memory Stats
	memoryDetails, err := getMemoryDetails()
	if err != nil {
		return nil, err
	}
	serverLoad.Memory = *memoryDetails

	// Storage Stats
	storageDetails, err := getStorageDetails()
	if err != nil {
		return nil, err
	}
	serverLoad.Storage = *storageDetails

	// Disk Stats
	// something is wrong with the disk.IOCounters() function
	// diskDetails, err := getDiskDetails()
	// if err != nil {
	// 	return nil, err
	// }
	// serverLoad.Disk = *diskDetails

	// Network Stats
	networkDetails, err := getNetworkDetails()
	if err != nil {
		return nil, err
	}
	serverLoad.Network = *networkDetails

	// Runtime Stats
	runtimeDetails, err := getRuntimeDetails()
	if err != nil {
		return nil, err
	}
	serverLoad.Runtime = *runtimeDetails

	return serverLoad, nil
}
