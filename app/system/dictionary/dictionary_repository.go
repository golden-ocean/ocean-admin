package dictionary

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/golden-ocean/ocean-admin/app/system/dictionary_item"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(e *Dictionary, db *sqlx.DB) error {
	b := sq.Insert(Table).SetMap(sq.Eq{
		FieldID:        xid.New().String(),
		FieldName:      e.Name,
		FieldCode:      e.Code,
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

func (r *Repository) Update(e *Dictionary, db *sqlx.DB) error {
	b := sq.Update(Table)
	if len(e.Name) > 0 {
		b = b.Set(FieldName, e.Name)
	}
	if len(e.Code) > 0 {
		b = b.Set(FieldCode, e.Code)
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

func (r *Repository) Delete(e *Dictionary, db *sqlx.DB) error {
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
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(sq.Like{FieldName: fmt.Sprint("%", w.Name, "%")})
	}
	if len(strings.TrimSpace(w.Code)) > 0 {
		b = b.Where(sq.Like{FieldCode: fmt.Sprint("%", w.Code, "%")})
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

func (r *Repository) QueryPage(w *WhereParams, db *sqlx.DB) ([]*Dictionary, error) {
	b := sq.Select(SelectFields...).From(Table)
	if len(strings.TrimSpace(w.Name)) > 0 {
		b = b.Where(sq.Like{FieldName: fmt.Sprint("%", w.Name, "%")})
	}
	if len(strings.TrimSpace(w.Code)) > 0 {
		b = b.Where(sq.Like{FieldCode: fmt.Sprint("%", w.Code, "%")})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	b = b.Limit(w.PageSize).Offset((w.Current - 1) * w.PageSize).OrderBy(FieldSort)
	sql, args, _ := b.ToSql()
	es := []*Dictionary{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) QueryByCode(c string, db *sqlx.DB) ([]*dictionary_item.DictionaryItem, error) {
	b := sq.Select("i.id", "i.label", "i.value", "i.color", "i.status", "i.sort").
		From("system_dictionary_items AS i").
		Join("system_dictionaries AS d ON i.dictionary_id = d.id").
		Where(sq.Eq{"d.code": c}).
		OrderBy("i.sort")
	sql, args, _ := b.ToSql()
	es := []*dictionary_item.DictionaryItem{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

// func (r *Repository) QueryByCode(c string, db *sqlx.DB) ([]*dictionary_item.DictionaryItem, error) {
// 	main := sq.Select(dictionary_item.SelectFields...).From(dictionary_item.Table)
// 	sub := sq.Select(FieldID).From(Table).Where(sq.Eq{FieldCode: c})
// 	main = main.Where(sq.Expr("dictionary_id IN (?)", sub)).OrderBy(dictionary_item.FieldSort)
// 	sql, args, _ := main.ToSql()
// 	es := []*dictionary_item.DictionaryItem{}
// 	err := db.Select(&es, sql, args...)
// 	return es, err
// }

func (r *Repository) ValidationFields(e *Dictionary, db *sqlx.DB) (*Dictionary, error) {
	name := e.Name
	code := e.Code
	id := e.ID
	b := sq.Select(FieldID, FieldName, FieldCode).From(Table)
	var or []sq.Sqlizer
	if len(name) > 0 {
		or = append(or, sq.Eq{FieldName: name})
	}
	if len(code) > 0 {
		or = append(or, sq.Eq{FieldCode: code})
	}
	b = b.Where(sq.Or(or))
	if len(id) > 0 {
		b = b.Where(sq.NotEq{FieldID: id})
	}
	sql, args, _ := b.Limit(1).ToSql()
	re := &Dictionary{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(e, args...)
	return re, err
}
