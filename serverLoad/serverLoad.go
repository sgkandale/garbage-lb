package serverLoad

type ResourceDetails struct {
	Total           string
	InUse           string
	Available       string
	InUsePercentage string
	UpStreamUse     string
	DownStreamUse   string
}

type ServerLoad struct {
	CPU     ResourceDetails
	Memory  ResourceDetails
	Disk    ResourceDetails
	Network ResourceDetails
	Storage ResourceDetails
}
