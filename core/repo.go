package core

import (
	"fmt"
	"reflect"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/modl"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/wcamarao/devtools/config"
)

// Repo abstracts the repository pattern for database access
type Repo struct {
	db    *sqlx.DB
	dbmap *modl.DbMap
}

// OffsetOptions contain parameters common to multi-row select statements
type OffsetOptions struct {
	GroupBy string
	OrderBy string
	Limit   uint64
	Offset  uint64
}

// NewRepo creates a generic repository for the given entity
func NewRepo(entity interface{}) (*Repo, error) {
	db, dbmap, err := NewDB(config.GetConfig())
	if err != nil {
		return nil, err
	}
	if dbmap.TableFor(entity) == nil {
		dbmap.AddTable(entity).SetKeys(false, "id")
	}
	return &Repo{
		db:    db,
		dbmap: dbmap,
	}, nil
}

// Create executes an insert statement
func (r *Repo) Create(entity interface{}) error {
	return r.dbmap.Insert(entity)
}

// Update executes an update statement
func (r *Repo) Update(entity interface{}) error {
	count, err := r.dbmap.Update(entity)
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("%d rows affected", count)
	}
	return nil
}

// Get executes a select statement scanning the resulting row to dest
func (r *Repo) Get(dest interface{}, query string, args ...interface{}) error {
	return r.db.Get(dest, query, args...)
}

// GetInt64 executes a select statement returning an int64
func (r *Repo) GetInt64(query string) (int64, error) {
	var n int64
	err := r.db.Get(&n, query)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// QueryOffset executes a select statement for offset pagination
func (r *Repo) QueryOffset(sb sq.SelectBuilder, args map[string]interface{}, opts *OffsetOptions) (*sqlx.Rows, error) {
	if len(opts.GroupBy) > 0 {
		sb = sb.GroupBy(opts.GroupBy)
	}
	if len(opts.OrderBy) > 0 {
		sb = sb.OrderBy(opts.OrderBy)
	}
	query, _, err := sb.Limit(opts.Limit).Offset(opts.Offset).ToSql()
	if err != nil {
		return nil, err
	}
	for k, v := range args {
		kind := reflect.TypeOf(v).Kind()
		if kind == reflect.Array || kind == reflect.Slice {
			args[k] = pq.Array(v)
		}
	}
	return r.db.NamedQuery(query, args)
}
