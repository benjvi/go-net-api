package netAPI

import (
	"fmt"
	"testing"
)

func checkStringAttribute(name, attr string, t *testing.T) {
	if attr == "" {
		t.Error(fmt.Sprintf("Resource attribute %s not found. Check raw response value and unmarshalling.", name))
	}
}
