package serverLoad

type ResourceUsage struct {
	Total           string `json:"total,omitempty"`
	InUse           string `json:"inUse,omitempty"`
	Available       string `json:"available,omitempty"`
	InUsePercentage string `json:"inUsePercentage,omitempty"`
	SentData        string `json:"sentData,omitempty"`
	ReceivedData    string `json:"receivedData,omitempty"`
	ReadBytes       string `json:"readBytes,omitempty"`
	WriteBytes      string `json:"writeBytes,omitempty"`
	IOps            string `json:"iops,omitempty"`
}

type ServerLoad struct {
	CPU     ResourceUsage `json:"cpu,omitempty"`
	Memory  ResourceUsage `json:"memory,omitempty"`
	Disk    ResourceUsage `json:"disk,omitempty"`
	Network ResourceUsage `json:"network,omitempty"`
	Storage ResourceUsage `json:"storage,omitempty"`
	Runtime ResourceUsage `json:"runtime,omitempty"`
}
