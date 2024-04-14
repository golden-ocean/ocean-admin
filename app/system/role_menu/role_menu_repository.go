package role_menu

import (
	"errors"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) CreateBatchWithTx(i *CreateInput, tx *sqlx.Tx) error {
	now := time.Now().Unix()
	b := sq.Insert(Table).
		Columns(FieldRoleID, FieldMenuID, FieldCreatedAt, FieldUpdatedAt, FieldCreatedBy)
	for _, menu_id := range i.MenuIDs {
		b = b.Values(i.RoleID, menu_id, now, now, i.CreatedBy)
	}
	sql, args, _ := b.ToSql()
	result, err := tx.Exec(sql, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(CreatedFail)
	}
	return err
}

func (r *Repository) DeleteBatchWithTx(i *DeleteInput, tx *sqlx.Tx) error {
	sql, args, _ := sq.Delete(Table).
		Where(sq.Eq{FieldRoleID: i.RoleID}).
		Where(sq.Eq{FieldMenuID: i.MenuIDs}).
		ToSql()
	result, err := tx.Exec(sql, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(DeletedFail)
	}
	return err
}
func (r *Repository) QueryAll(w *WhereParams, db *sqlx.DB) ([]*RoleMenu, error) {
	b := sq.Select(FieldRoleID, FieldMenuID).From(Table)
	if len(w.RoleIDs) > 0 {
		b = b.Where(sq.Eq{FieldRoleID: w.RoleIDs})
	}
	if len(w.RoleID) > 0 {
		b = b.Where(sq.Eq{FieldRoleID: w.RoleID})
	}
	sql, args, _ := b.OrderBy(FieldCreatedAt).ToSql()
	es := []*RoleMenu{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) QueryMenuIDs(w *WhereParams, db *sqlx.DB) ([]string, error) {
	b := sq.Select(FieldMenuID).From(Table)
	if len(w.RoleIDs) > 0 {
		b = b.Where(sq.Eq{FieldRoleID: w.RoleIDs})
	}
	if len(w.RoleID) > 0 {
		b = b.Where(sq.Eq{FieldRoleID: w.RoleID})
	}
	sql, args, _ := b.OrderBy(FieldCreatedAt).ToSql()
	menu_ids := []string{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&menu_ids, args...)
	return menu_ids, err
}

func (r *Repository) Exist(e *RoleMenu, db *sqlx.DB) (bool, error) {
	b := sq.Select(FieldRoleID, FieldMenuID).Prefix("SELECT EXISTS (").From(Table)
	if len(e.RoleID) > 0 {
		b = b.Where(sq.Eq{FieldRoleID: e.RoleID})
	}
	if len(e.MenuID) > 0 {
		b = b.Where(sq.Eq{FieldMenuID: e.MenuID})
	}
	b = b.Suffix(")")
	sql, args, _ := b.ToSql()
	var exist bool
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&exist, args...)
	return exist, err
}
