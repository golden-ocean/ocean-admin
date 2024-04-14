package dictionary_item

import (
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"

	sq "github.com/Masterminds/squirrel"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(e *DictionaryItem, db *sqlx.DB) error {
	b := sq.Insert(Table).SetMap(sq.Eq{
		FieldID:           xid.New().String(),
		FieldLabel:        e.Label,
		FieldValue:        e.Value,
		FieldColor:        e.Color,
		FieldDictionaryID: e.DictionaryID,
		FieldStatus:       e.Status,
		FieldSort:         e.Sort,
		FieldRemark:       e.Remark,
		FieldCreatedAt:    time.Now().Unix(),
		FieldUpdatedAt:    time.Now().Unix(),
		FieldCreatedBy:    e.CreatedBy,
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

func (r *Repository) Update(e *DictionaryItem, db *sqlx.DB) error {
	b := sq.Update(Table)
	if len(e.Label) > 0 {
		b = b.Set(FieldLabel, e.Label)
	}
	if len(e.Value) > 0 {
		b = b.Set(FieldValue, e.Value)
	}
	if len(e.Color) > 0 {
		b = b.Set(FieldColor, e.Color)
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

func (r *Repository) Delete(e *DictionaryItem, db *sqlx.DB) error {
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

func (r *Repository) QueryCount(w *WhereParams, db *sqlx.DB) (uint64, error) {
	b := sq.Select("COUNT('id')").From(Table)
	if len(w.DictionaryID) > 0 {
		b = b.Where(sq.Eq{FieldDictionaryID: w.DictionaryID})
	}
	if len(w.Label) > 0 {
		b = b.Where(sq.Like{FieldLabel: fmt.Sprint("%", w.Label, "%")})
	}
	if len(w.Value) > 0 {
		b = b.Where(sq.Like{FieldValue: fmt.Sprint("%", w.Value, "%")})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	sql, args, _ := b.ToSql()
	var count uint64
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&count, args...)
	return count, err
}

func (r *Repository) QueryPage(w *WhereParams, db *sqlx.DB) ([]*DictionaryItem, error) {
	b := sq.Select(SelectFields...).From(Table)
	if len(w.DictionaryID) > 0 {
		b = b.Where(sq.Eq{FieldDictionaryID: w.DictionaryID})
	}
	if len(w.Label) > 0 {
		b = b.Where(sq.Like{FieldLabel: fmt.Sprint("%", w.Label, "%")})
	}
	if len(w.Value) > 0 {
		b = b.Where(sq.Like{FieldValue: fmt.Sprint("%", w.Value, "%")})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	b = b.Limit(w.PageSize).Offset((w.Current - 1) * w.PageSize).OrderBy(FieldSort)
	sql, args, _ := b.ToSql()
	es := []*DictionaryItem{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) Exist(e *DictionaryItem, db *sqlx.DB) (bool, error) {
	b := sq.Select(FieldID).Prefix("SELECT EXISTS (").From(Table)
	if len(e.DictionaryID) > 0 {
		b = b.Where(sq.Eq{FieldDictionaryID: e.DictionaryID})
	}
	b = b.Suffix(")")
	sql, args, _ := b.ToSql()
	var exist bool
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&exist, args...)
	return exist, err
}

func (r *Repository) ValidationFields(c *DictionaryItem, db *sqlx.DB) (*DictionaryItem, error) {
	label := c.Label
	value := c.Value
	dictionaryId := c.DictionaryID
	id := c.ID
	// 当 dictionary_id 相等的时候 name 和 code 相同 则返回结果
	// dictionary_id name 联合唯一
	// dictionary_id code 联合唯一
	b := sq.Select(FieldID, FieldLabel, FieldValue, FieldDictionaryID).From(Table)
	var or []sq.Sqlizer
	if len(label) > 0 {
		or = append(or, sq.Eq{FieldLabel: label})
	}
	if len(value) > 0 {
		or = append(or, sq.Eq{FieldValue: value})
	}
	b = b.Where(sq.Eq{FieldDictionaryID: dictionaryId})
	b = b.Where(sq.Or(or))
	if len(id) > 0 {
		b = b.Where(sq.NotEq{FieldID: id})
	}
	sql, args, _ := b.Limit(1).ToSql()
	e := &DictionaryItem{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(e, args...)
	return e, err
}
