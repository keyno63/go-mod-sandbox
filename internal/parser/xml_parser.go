package parser

import (
	"encoding/xml"
	"fmt"
)

func isParsableXml(value string) bool {
	if err := xml.Unmarshal([]byte(value), &struct{}{}); err != nil {
		fmt.Printf("failed to paese xml, value = [%s], reason = [%s]\n", value, err.Error())
		return false
	}
	return true
}
