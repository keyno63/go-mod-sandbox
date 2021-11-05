package parser

import (
	"encoding/json"
	"fmt"
)

func isParsableJSON(value string) bool {
	if err := json.Unmarshal([]byte(value), &struct{}{}); err != nil {
		fmt.Printf("failed to paese json, value = [%s], reason = [%s]\n", value, err.Error())
		return false
	}
	return true
}
