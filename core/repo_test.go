package core_test

import (
	"testing"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/stretchr/testify/suite"
	"github.com/wcamarao/devtools/config"
	"github.com/wcamarao/devtools/core"
)

var db, _, _ = core.NewDB(config.GetConfig())
var r, _ = core.NewRepo(EntitySample{})

type RepoTestSuite struct {
	suite.Suite
}

func (s *RepoTestSuite) SetupTest() {
	db.MustExec("truncate table entity_sample")
}

func TestRepo(t *testing.T) {
	suite.Run(t, new(RepoTestSuite))
}

func (s *RepoTestSuite) TestCreate() {
	t1 := time.Now().UnixNano()

	err := r.Create(&EntitySample{
		ID: "foo",
	})
	s.Nil(err)

	t2 := time.Now().UnixNano()

	entity := &EntitySample{}
	err = db.Get(entity, "select * from entity_sample where id = $1", "foo")
	s.Nil(err)

	s.Equal("foo", entity.ID)
	s.Equal(false, entity.IsActive)
	s.True(entity.CreatedAt.UnixNano() >= t1)
	s.True(entity.CreatedAt.UnixNano() <= t2)
	s.Equal(entity.CreatedAt, entity.UpdatedAt)
}

func (s *RepoTestSuite) TestUpdate() {
	err := r.Create(&EntitySample{
		ID: "foo",
	})
	s.Nil(err)

	entity := &EntitySample{}
	err = db.Get(entity, "select * from entity_sample where id = $1", "foo")
	s.Nil(err)

	t1 := time.Now().UnixNano()

	entity.IsActive = true
	err = r.Update(entity)
	s.Nil(err)

	t2 := time.Now().UnixNano()

	entity = &EntitySample{}
	err = db.Get(entity, "select * from entity_sample where id = $1", "foo")
	s.Nil(err)

	s.Equal("foo", entity.ID)
	s.Equal(true, entity.IsActive)
	s.True(entity.UpdatedAt.UnixNano() >= t1)
	s.True(entity.UpdatedAt.UnixNano() <= t2)
	s.NotEqual(entity.CreatedAt, entity.UpdatedAt)
}

func (s *RepoTestSuite) TestGet() {
	err := r.Create(&EntitySample{
		ID: "foo",
	})
	s.Nil(err)

	entity := &EntitySample{}
	err = r.Get(entity, "select * from entity_sample where id = $1", "foo")
	s.Nil(err)

	s.Equal("foo", entity.ID)
}

func (s *RepoTestSuite) TestGetInt64() {
	n, err := r.GetInt64("select count(id) from entity_sample")
	s.Nil(err)
	s.Equal(int64(0), n)

	n, err = r.GetInt64("select 42")
	s.Nil(err)
	s.Equal(int64(42), n)
}

func (s *RepoTestSuite) TestQueryOffsetPagination() {
	expectedIds := []string{}
	actualIds := []string{}

	for i := 0; i < 10; i++ {
		id, err := core.NewID()
		s.Nil(err)
		expectedIds = append(expectedIds, id)
		err = r.Create(&EntitySample{
			ID: id,
		})
		s.Nil(err)
	}

	sb := sq.Select("*").From("entity_sample")
	rows, err := r.QueryOffset(sb, map[string]interface{}{}, &core.OffsetOptions{
		OrderBy: "id",
		Limit:   5,
		Offset:  0,
	})
	s.Nil(err)

	for rows.Next() {
		entity := &EntitySample{}
		rows.StructScan(entity)
		actualIds = append(actualIds, entity.ID)
	}

	s.Equal(expectedIds[:5], actualIds)

	rows, err = r.QueryOffset(sb, map[string]interface{}{}, &core.OffsetOptions{
		OrderBy: "id",
		Limit:   5,
		Offset:  5,
	})
	s.Nil(err)

	for rows.Next() {
		entity := &EntitySample{}
		rows.StructScan(entity)
		actualIds = append(actualIds, entity.ID)
	}

	s.Equal(expectedIds, actualIds)

	rows, err = r.QueryOffset(sb, map[string]interface{}{}, &core.OffsetOptions{
		OrderBy: "id",
		Limit:   5,
		Offset:  10,
	})
	s.Nil(err)
	s.False(rows.Next())
}

func (s *RepoTestSuite) TestQueryOffsetArrayArg() {
	allIds := []string{"foo", "bar", "zip", "zoo"}
	expectedIds := []string{"foo", "zip"}
	actualIds := []string{}

	for _, id := range allIds {
		err := r.Create(&EntitySample{
			ID: id,
		})
		s.Nil(err)
	}

	sb := sq.Select("*").From("entity_sample").Where("id = ANY (:ids)")
	args := map[string]interface{}{"ids": expectedIds}
	rows, err := r.QueryOffset(sb, args, &core.OffsetOptions{
		OrderBy: "id",
		Limit:   10,
	})
	s.Nil(err)

	for rows.Next() {
		entity := &EntitySample{}
		rows.StructScan(entity)
		actualIds = append(actualIds, entity.ID)
	}

	s.Equal(expectedIds, actualIds)
}
