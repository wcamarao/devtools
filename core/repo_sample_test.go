package core_test

import (
	"time"

	"github.com/jmoiron/modl"
)

// EntitySample represents table defined in db/schema.sql
type EntitySample struct {
	ID        string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// PreInsert runs before insert statements
func (c *EntitySample) PreInsert(s modl.SqlExecutor) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	return nil
}

// PreUpdate runs before update statements
func (c *EntitySample) PreUpdate(s modl.SqlExecutor) error {
	c.UpdatedAt = time.Now()
	return nil
}
