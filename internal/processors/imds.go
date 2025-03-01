package processors

import (
	"encoding/base64"
	"strings"
)

// RetrieveMetadataAccessPath retrieves the DNS query host from the data.
//
// The input to this function is a base64 encoded string.
//
// An expected usage might be:
//
//	path, err := RetrieveMetadataAccessPath(data)
func RetrieveMetadataAccessPath(data string) (string, error) {
	payload, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return "", err
	}

	firstLine := strings.Split(string(payload), "\n")[0]
	path := strings.Split(firstLine, " ")[1]

	return path, nil
}
