package staff_role

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
		Columns(FieldStaffID, FieldRoleID, FieldCreatedAt, FieldUpdatedAt, FieldCreatedBy)
	for _, role_id := range i.RoleIDs {
		b = b.Values(i.StaffID, role_id, now, now, i.CreatedBy)
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
		Where(sq.Eq{FieldStaffID: i.StaffID}).
		Where(sq.Eq{FieldRoleID: i.RoleIDs}).
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

func (r *Repository) QueryAll(w *WhereParams, db *sqlx.DB) ([]*StaffRole, error) {
	b := sq.Select(FieldStaffID, FieldRoleID).From(Table)
	if len(w.StaffIDs) > 0 {
		b = b.Where(sq.Eq{FieldStaffID: w.StaffIDs})
	}
	if len(w.StaffID) > 0 {
		b = b.Where(sq.Eq{FieldStaffID: w.StaffID})
	}
	sql, args, _ := b.OrderBy(FieldCreatedAt).ToSql()
	es := []*StaffRole{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) Exist(e *StaffRole, db *sqlx.DB) (bool, error) {
	b := sq.Select(FieldStaffID, FieldRoleID).Prefix("SELECT EXISTS (").From(Table)
	if len(e.StaffID) > 0 {
		b = b.Where(sq.Eq{FieldStaffID: e.StaffID})
	}
	if len(e.RoleID) > 0 {
		b = b.Where(sq.Eq{FieldRoleID: e.RoleID})
	}
	b = b.Suffix(")")
	sql, args, _ := b.ToSql()
	var exist bool
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&exist, args...)
	return exist, err
}
