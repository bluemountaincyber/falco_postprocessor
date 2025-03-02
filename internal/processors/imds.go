package processors

import (
	"encoding/base64"
	"fmt"
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
	if data == "" {
		return "", fmt.Errorf("data is empty")
	}

	payload, err := base64.StdEncoding.DecodeString(data)

	if err != nil {
		return "", fmt.Errorf("error decoding base64: %v", err)
	}

	firstLine := strings.Split(string(payload), "\n")[0]
	path := strings.Split(firstLine, " ")[1]

	return path, nil
}
