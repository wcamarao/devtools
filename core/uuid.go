package core

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// NewID generates a random UUID V4 prefixed with a sortable unix nano timestamp
func NewID() (string, error) {
	t := fmt.Sprintf("%x", time.Now().UnixNano())
	v4, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", t, v4.String()), nil
}
