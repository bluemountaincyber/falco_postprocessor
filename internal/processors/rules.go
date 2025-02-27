package main

import (
	"encoding/base64"
	"fmt"
)

func retrieveDNSQueryHost(data string) (string, error) {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	var domain []byte
	byteCounter := 13
	if len(payload) < byteCounter {
		return "", fmt.Errorf("invalid payload")
	}
	wordLen := int(payload[12])

	for {
		domain = append(domain, payload[byteCounter:byteCounter+wordLen]...)
		if payload[byteCounter+wordLen] == 0 {
			break
		} else {
			domain = append(domain, '.')
			byteCounter += wordLen + 1
			wordLen = int(payload[byteCounter-1])
		}
	}

	return string(domain), nil
}
