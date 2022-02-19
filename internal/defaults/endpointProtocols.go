package defaults

import "strings"

func GetEndpointProtocol(protocol string) string {

	if strings.EqualFold(protocol, "http") {
		return "http"
	}
	if strings.EqualFold(protocol, "https") {
		return "https"
	}
	if strings.EqualFold(protocol, "tcp") {
		return "tcp"
	}
	if strings.EqualFold(protocol, "tcp4") {
		return "tcp4"
	}
	if strings.EqualFold(protocol, "tcp6") {
		return "tcp6"
	}

	return ""
}
