package organization

import (
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(e *Organization, db *sqlx.DB) error {
	b := sq.Insert(Table).SetMap(sq.Eq{
		FieldID:        xid.New().String(),
		FieldName:      e.Name,
		FieldCode:      e.Code,
		FieldParentID:  e.ParentID,
		FieldStatus:    e.Status,
		FieldSort:      e.Sort,
		FieldRemark:    e.Remark,
		FieldCreatedAt: time.Now().Unix(),
		FieldUpdatedAt: time.Now().Unix(),
		FieldCreatedBy: e.CreatedBy,
	})
	sql, args, _ := b.ToSql()
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *Repository) Update(e *Organization, db *sqlx.DB) error {
	b := sq.Update(Table)
	if len(e.Name) > 0 {
		b = b.Set(FieldName, e.Name)
	}
	if len(e.Code) > 0 {
		b = b.Set(FieldCode, e.Code)
	}
	if len(e.ParentID) > 0 {
		b = b.Set(FieldParentID, e.ParentID)
	}
	if len(e.Status) > 0 {
		b = b.Set(FieldStatus, e.Status)
	}
	if e.Sort > 0 {
		b = b.Set(FieldSort, e.Sort)
	}
	b = b.Set(FieldRemark, e.Remark)
	b = b.SetMap(sq.Eq{
		FieldUpdatedAt: time.Now().Unix(),
		FieldUpdatedBy: e.UpdatedBy,
	}).Where(sq.Eq{FieldID: e.ID})
	sql, args, _ := b.ToSql()
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *Repository) Delete(e *Organization, db *sqlx.DB) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{FieldID: e.ID}).ToSql()
	result, err := db.Exec(sql, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}

func (r *Repository) QueryAll(w *WhereParams, db *sqlx.DB) ([]*Organization, error) {
	b := sq.Select(SelectFields...).From(Table)
	if len(w.Name) > 0 {
		b = b.Where(sq.Like{FieldName: fmt.Sprint("%", w.Name, "%")})
	}
	if len(w.Code) > 0 {
		b = b.Where(sq.Like{FieldCode: fmt.Sprint("%", w.Code, "%")})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	b = b.OrderBy(FieldSort)
	sql, args, _ := b.ToSql()
	es := []*Organization{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) Exist(e *Organization, db *sqlx.DB) (bool, error) {
	b := sq.Select(FieldID).Prefix("SELECT EXISTS (").From(Table)
	if len(e.ID) > 0 {
		b = b.Where(sq.Eq{FieldID: e.ID})
	}
	if len(e.ParentID) > 0 {
		b = b.Where(sq.Eq{FieldParentID: e.ParentID})
	}
	b = b.Suffix(")")
	sql, args, _ := b.ToSql()
	var exist bool
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&exist, args...)
	return exist, err
}

func (r *Repository) ValidationFields(c *Organization, db *sqlx.DB) (*Organization, error) {
	b := sq.Select(FieldID, FieldName, FieldCode).From(Table)
	var or []sq.Sqlizer
	if len(c.Name) > 0 {
		or = append(or, sq.Eq{FieldName: c.Name})
	}
	if len(c.Code) > 0 {
		or = append(or, sq.Eq{FieldCode: c.Code})
	}
	b = b.Where(sq.Or(or))
	if len(c.ParentID) > 0 {
		b = b.Where(sq.Eq{FieldParentID: c.ParentID})
	}
	if len(c.ID) > 0 {
		b = b.Where(sq.NotEq{FieldID: c.ID})
	}
	sql, args, _ := b.Limit(1).ToSql()
	e := &Organization{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(e, args...)
	return e, err
}
