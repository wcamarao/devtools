package core_test

import (
	"testing"

	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/wcamarao/devtools/config"
	"github.com/wcamarao/devtools/core"
)

func TestNewDB(t *testing.T) {
	db, dbmap, err := core.NewDB(config.GetConfig())
	assert.Nil(t, err)

	assert.Equal(t, "postgres", db.DriverName())
	assert.Equal(t, modl.PostgresDialect{}, dbmap.Dialect)
	assert.Equal(t, "entity_name", modl.TableNameMapper("EntityName"))
	assert.Equal(t, "entity_name", sqlx.NameMapper("EntityName"))

	var n int
	err = db.Get(&n, "select 42")
	assert.Nil(t, err)
	assert.Equal(t, 42, n)
}
