package core

import (
	"encoding/json"
	"fmt"
)

type allowedAttributes struct {
	attributes map[string]bool
}

// AllowedAttributes stores a set of allowed attributes
func AllowedAttributes(attributes ...string) *allowedAttributes {
	a := &allowedAttributes{}
	for _, attr := range attributes {
		a.attributes[attr] = true
	}
	return a
}

// Validate restricts JSON payload to allowed attributes only
func (a *allowedAttributes) Validate(jsonPayload string) error {
	payload := map[string]interface{}{}
	err := json.Unmarshal([]byte(jsonPayload), payload)
	if err != nil {
		return fmt.Errorf("Failed parsing payload")
	}

	for key, _ := range payload {
		if !a.attributes[key] {
			return fmt.Errorf("Attribute not allowed: %s", key)
		}
	}

	return nil
}
