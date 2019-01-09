package core_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wcamarao/devtools/core"
)

func TestNewID(t *testing.T) {
	t1 := time.Now().UnixNano()

	id, err := core.NewID()
	assert.Nil(t, err)

	t2 := time.Now().UnixNano()

	nano, err := strconv.ParseInt(id[:16], 16, 64)
	assert.Nil(t, err)

	assert.Equal(t, 53, len(id))
	assert.True(t, nano >= t1)
	assert.True(t, nano <= t2)

	_, err = uuid.Parse(id[17:])
	assert.Nil(t, err)
}
