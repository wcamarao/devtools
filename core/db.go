package core

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	"github.com/wcamarao/devtools/config"
)

// NewDB creates a database connection
func NewDB(c *config.Config) (*sqlx.DB, *modl.DbMap, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DB.Host, c.DB.User, c.DB.Pass, c.DB.Name, c.DB.SSLMode,
	))
	if err != nil {
		return nil, nil, err
	}
	db.MapperFunc(strcase.ToSnake)
	modl.TableNameMapper = strcase.ToSnake
	sqlx.NameMapper = strcase.ToSnake
	return db, modl.NewDbMap(db.DB, modl.PostgresDialect{}), nil
}
