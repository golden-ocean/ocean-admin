package menu

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

func (r *Repository) Create(e *Menu, db *sqlx.DB) error {
	b := sq.Insert(Table).SetMap(sq.Eq{
		FieldID:        xid.New().String(),
		FieldName:      e.Name,
		FieldParentID:  e.ParentID,
		FieldIcon:      e.Icon,
		FieldPath:      e.Path,
		FieldType:      e.Type,
		FieldMethod:    e.Method,
		FieldComponent: e.Component,
		FieldVisible:   e.Visible,
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

func (r *Repository) Update(e *Menu, db *sqlx.DB) error {
	b := sq.Update(Table)
	if len(e.Name) > 0 {
		b = b.Set(FieldName, e.Name)
	}
	if len(e.ParentID) > 0 {
		b = b.Set(FieldParentID, e.ParentID)
	}
	if len(e.Icon) > 0 {
		b = b.Set(FieldIcon, e.Icon)
	}
	if len(e.Path) > 0 {
		b = b.Set(FieldPath, e.Path)
	}
	if len(e.Type) > 0 {
		b = b.Set(FieldType, e.Type)
	}
	if len(e.Method) > 0 {
		b = b.Set(FieldMethod, e.Method)
	}
	if len(e.Component) > 0 {
		b = b.Set(FieldComponent, e.Component)
	}
	if len(e.Status) > 0 {
		b = b.Set(FieldStatus, e.Status)
	}
	if e.Sort > 0 {
		b = b.Set(FieldSort, e.Sort)
	}
	b = b.Set(FieldRemark, e.Remark)
	b = b.Set(FieldVisible, e.Visible)
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

func (r *Repository) Delete(e *Menu, db *sqlx.DB) error {
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

func (r *Repository) QueryAll(w *WhereParams, db *sqlx.DB) ([]*Menu, error) {
	b := sq.Select(SelectFields...).From(Table)
	if len(w.Name) > 0 {
		b = b.Where(sq.Like{FieldName: fmt.Sprint("%", w.Name, "%")})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	if len(w.Visible) > 0 {
		b = b.Where(sq.Eq{FieldVisible: w.Visible})
	}
	b = b.OrderBy(FieldSort)
	sql, args, _ := b.ToSql()
	es := []*Menu{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}
func (r *Repository) QueryAllByIDs(ids []string, db *sqlx.DB) ([]*Menu, error) {
	b := sq.Select(SelectFields...).From(Table)
	if len(ids) > 0 {
		b = b.Where(sq.Eq{FieldID: ids})
	}
	b = b.OrderBy(FieldSort)
	sql, args, _ := b.ToSql()
	es := []*Menu{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) Exist(e *Menu, db *sqlx.DB) (bool, error) {
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

func (r *Repository) QueryByStaffID(id string, db *sqlx.DB) ([]*Menu, error) {
	b := sq.Select("m.id, m.name, m.parent_id, m.icon, m.path, m.type, m.method, m.component, m.visible, m.status, m.sort").
		Distinct().
		From("system_menus AS m").
		Join("system_roles_menus AS rm ON m.id = rm.menu_id").
		Join("system_staffs_roles AS sr ON rm.role_id = sr.role_id").
		// Join("system_roles AS r ON sr.role_id = r.id").
		Where(sq.Eq{"sr.staff_id": id}).OrderBy("m.sort")
	sql, args, _ := b.ToSql()
	es := []*Menu{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}
