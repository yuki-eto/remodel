package dao

import (
	"example/entity"
	"time"

	"github.com/juju/errors"
	"go.knocknote.io/rapidash"
)

type User interface {
	Save(e *entity.User) error
	Delete(e *entity.User) error
	FindByID(k0 uint64) (*entity.User, error)
	FindByIDs(k0 []uint64) (entity.Users, error)
	FindByUuid(k0 string) (*entity.User, error)
	FindByUuids(k0 []string) (entity.Users, error)
	FindByOutsideUserID(k0 string) (*entity.User, error)
	FindByOutsideUserIDs(k0 []string) (entity.Users, error)
}

type UserImpl struct {
	tableName string
	txGetter  func() (*rapidash.Tx, error)
	qb        func() *rapidash.QueryBuilder
}

func NewUser(txGetter func(string) (*rapidash.Tx, error)) User {
	return &UserImpl{
		qb: func() *rapidash.QueryBuilder {
			return rapidash.NewQueryBuilder("users")
		},
		tableName: "users",
		txGetter: func() (*rapidash.Tx, error) {
			return txGetter("users")
		},
	}
}

func (d *UserImpl) Save(e *entity.User) error {
	tx, err := d.txGetter()
	if err != nil {
		return errors.Trace(err)
	}
	now := time.Now()
	e.UpdatedAt = &now
	if e.ID == 0 {
		e.CreatedAt = &now
		id, err := tx.CreateByTable(d.tableName, e)
		if err != nil {
			return errors.Trace(err)
		}
		e.ID = uint64(id)
		return nil
	}
	b := d.qb().Eq("id", e.ID)
	m := map[string]interface{}{
		"access_token":    e.AccessToken,
		"name":            e.Name,
		"outside_user_id": e.OutsideUserID,
		"updated_at":      e.UpdatedAt,
		"uuid":            e.Uuid,
	}
	if err := tx.UpdateByQueryBuilder(b, m); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (d *UserImpl) Delete(e *entity.User) error {
	tx, err := d.txGetter()
	if err != nil {
		return errors.Trace(err)
	}
	if e.ID == 0 {
		return errors.New("cannot delete without identifier")
	}
	b := d.qb().Eq("id", e.ID)
	if err := tx.DeleteByQueryBuilder(b); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (d *UserImpl) FindByID(k0 uint64) (*entity.User, error) {
	tx, err := d.txGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	b := d.qb().Eq("id", k0)
	e := &entity.User{}
	if err := tx.FindByQueryBuilder(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	if e.ID == 0 {
		return nil, nil
	}
	return e, nil
}

func (d *UserImpl) FindByIDs(k0 []uint64) (entity.Users, error) {
	tx, err := d.txGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	b := d.qb().In("id", k0)
	e := &entity.Users{}
	if err := tx.FindByQueryBuilder(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	return *e, nil
}

func (d *UserImpl) FindByUuid(k0 string) (*entity.User, error) {
	tx, err := d.txGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	b := d.qb().Eq("uuid", k0)
	e := &entity.User{}
	if err := tx.FindByQueryBuilder(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	if e.ID == 0 {
		return nil, nil
	}
	return e, nil
}

func (d *UserImpl) FindByUuids(k0 []string) (entity.Users, error) {
	tx, err := d.txGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	b := d.qb().In("uuid", k0)
	e := &entity.Users{}
	if err := tx.FindByQueryBuilder(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	return *e, nil
}

func (d *UserImpl) FindByOutsideUserID(k0 string) (*entity.User, error) {
	tx, err := d.txGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	b := d.qb().Eq("outside_user_id", k0)
	e := &entity.User{}
	if err := tx.FindByQueryBuilder(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	if e.ID == 0 {
		return nil, nil
	}
	return e, nil
}

func (d *UserImpl) FindByOutsideUserIDs(k0 []string) (entity.Users, error) {
	tx, err := d.txGetter()
	if err != nil {
		return nil, errors.Trace(err)
	}
	b := d.qb().In("outside_user_id", k0)
	e := &entity.Users{}
	if err := tx.FindByQueryBuilder(b, e); err != nil {
		return nil, errors.Trace(err)
	}
	return *e, nil
}
