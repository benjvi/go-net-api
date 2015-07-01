package netAPI

import (
	"fmt"
	"testing"
)

func checkStringAttributePresent(name, value string, t *testing.T) {
	if value == "" {
		t.Error(fmt.Sprintf("Resource attribute %s not found. Check raw response value and unmarshalling.", value))
	}
}

func checkStringAttributeValue(name, value, expected string, t *testing.T) {
	if value == "" {
		t.Error(fmt.Sprintf("Resource attribute %s value found: %s does not match expected value: %s", name, value, expected))
	}
}
