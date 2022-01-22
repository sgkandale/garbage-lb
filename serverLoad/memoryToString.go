package serverLoad

import "fmt"

func convertMemoryToString(memory uint64) string {
	// memory is in bytes

	if memory <= 0 {
		return "0"
	} else if memory < 1024 {
		return fmt.Sprintf("%.2f B", float64(memory))
	} else if memory < 1024*1024 {
		return fmt.Sprintf("%.2f K", float64(memory)/1024)
	} else if memory < 1024*1024*1024 {
		return fmt.Sprintf("%.2f M", float64(memory)/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f G", float64(memory)/(1024*1024*1024))
	}
}

func convertMemoryToPercentage(memory uint64, totalMemory uint64) string {
	// memory is in bytes
	// totalMemory is in bytes

	if totalMemory <= 0 {
		return "0"
	} else if memory <= 0 {
		return "0"
	} else {
		return fmt.Sprintf(
			"%.2f%%",
			float64(memory)*100/float64(totalMemory))
	}
}
