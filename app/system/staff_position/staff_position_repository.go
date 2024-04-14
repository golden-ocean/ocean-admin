package staff_position

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
		Columns(FieldStaffID, FieldPositionID, FieldCreatedAt, FieldUpdatedAt, FieldCreatedBy)
	for _, position_id := range i.PositionIDs {
		b = b.Values(i.StaffID, position_id, now, now, i.CreatedBy)
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
		Where(sq.Eq{FieldPositionID: i.PositionIDs}).
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

func (r *Repository) QueryAll(w *WhereParams, db *sqlx.DB) ([]*StaffPosition, error) {
	b := sq.Select(FieldStaffID, FieldPositionID).From(Table)
	if len(w.StaffIDs) > 0 {
		b = b.Where(sq.Eq{FieldStaffID: w.StaffIDs})
	}
	if len(w.StaffID) > 0 {
		b = b.Where(sq.Eq{FieldStaffID: w.StaffID})
	}
	sql, args, _ := b.OrderBy(FieldCreatedAt).ToSql()
	entities := []*StaffPosition{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&entities, args...)
	return entities, err
}

func (r *Repository) Exist(e *StaffPosition, db *sqlx.DB) (bool, error) {
	b := sq.Select(FieldStaffID, FieldPositionID).Prefix("SELECT EXISTS (").From(Table)
	if len(e.StaffID) > 0 {
		b = b.Where(sq.Eq{FieldStaffID: e.StaffID})
	}
	if len(e.PositionID) > 0 {
		b = b.Where(sq.Eq{FieldPositionID: e.PositionID})
	}
	b = b.Suffix(")")
	sql, args, _ := b.ToSql()
	var exist bool
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&exist, args...)
	return exist, err
}
