package defaults

import "strings"

func GetEndpointProtocol(protocol string) string {

	if strings.EqualFold(protocol, "http") {
		return "http"
	}
	if strings.EqualFold(protocol, "tcp") {
		return "tcp"
	}
	if strings.EqualFold(protocol, "udp") {
		return "udp"
	}

	return ""
}
