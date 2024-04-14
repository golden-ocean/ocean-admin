package staff

import (
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/golden-ocean/ocean-admin/pkg/utils"
	"github.com/jmoiron/sqlx"
	"github.com/rs/xid"
)

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

// func (r *Repository) Create(e *Staff, db *sqlx.DB) error {
// 	b := sq.Insert(Table).SetMap(sq.Eq{
// 		FieldID:             xid.New().String(),
// 		FieldUsername:       e.Username,
// 		FieldPassword:       utils.GeneratePassword(e.Password),
// 		FieldName:           e.Name,
// 		FieldEmail:          e.Email,
// 		FieldMobile:         e.Mobile,
// 		FieldAvatar:         e.Avatar,
// 		FieldGender:         e.Gender,
// 		FieldOrganizationID: e.OrganizationID,
// 		FieldWorkStatus:     e.WorkStatus,
// 		FieldStatus:         e.Status,
// 		FieldSort:           e.Sort,
// 		FieldRemark:         e.Remark,
// 		FieldCreatedAt:      time.Now().Unix(),
// 		FieldUpdatedAt:      time.Now().Unix(),
// 		FieldCreatedBy:      e.CreatedBy,
// 	})
// 	sql, args, _ := b.ToSql()
// 	result, err := db.Exec(sql, args...)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New(CreatedFail)
// 	}
// 	return err
// }

// func (r *Repository) Update(e *UpdateInput, db *sqlx.DB) error {
// 	b := sq.Update(Table)
// 	if len(e.Username) > 0 {
// 		b = b.Set(FieldUsername, e.Username)
// 	}
// 	if len(e.Name) > 0 {
// 		b = b.Set(FieldName, e.Name)
// 	}
// 	if len(e.Email) > 0 {
// 		b = b.Set(FieldEmail, e.Email)
// 	}
// 	if len(e.Mobile) > 0 {
// 		b = b.Set(FieldMobile, e.Mobile)
// 	}
// 	if len(e.Avatar) > 0 {
// 		b = b.Set(FieldAvatar, e.Avatar)
// 	}
// 	if len(e.Gender) > 0 {
// 		b = b.Set(FieldGender, e.Gender)
// 	}
// 	if len(e.WorkStatus) > 0 {
// 		b = b.Set(FieldWorkStatus, e.WorkStatus)
// 	}
// 	if len(e.Status) > 0 {
// 		b = b.Set(FieldStatus, e.Status)
// 	}
// 	if e.Sort > 0 {
// 		b = b.Set(FieldSort, e.Sort)
// 	}
// 	if len(e.OrganizationID) > 0 {
// 		b = b.Set(FieldOrganizationID, e.OrganizationID)
// 	}
// 	b = b.Set(FieldRemark, e.Remark)
// 	b = b.SetMap(sq.Eq{
// 		FieldUpdatedAt: time.Now().Unix(),
// 		FieldUpdatedBy: e.UpdatedBy,
// 	}).Where(sq.Eq{FieldID: e.ID})
// 	sql, args, _ := b.ToSql()
// 	result, err := db.Exec(sql, args...)
// 	if err != nil {
// 		return err
// 	}
// 	rowsAffected, err := result.RowsAffected()
// 	if rowsAffected == 0 {
// 		return errors.New(UpdatedFail)
// 	}
// 	return err
// }

func (r *Repository) QueryCount(w *WhereParams, db *sqlx.DB) (uint64, error) {
	b := sq.Select("COUNT('id')").From(Table)
	if len(w.Username) > 0 {
		b = b.Where(sq.Like{FieldUsername: fmt.Sprint("%", w.Username, "%")})
	}
	if len(w.Name) > 0 {
		b = b.Where(sq.Like{FieldName: fmt.Sprint("%", w.Name, "%")})
	}
	if len(w.Email) > 0 {
		b = b.Where(sq.Like{FieldEmail: fmt.Sprint("%", w.Email, "%")})
	}
	if len(w.Mobile) > 0 {
		b = b.Where(sq.Like{FieldMobile: fmt.Sprint("%", w.Mobile, "%")})
	}
	if len(w.Gender) > 0 {
		b = b.Where(sq.Eq{FieldGender: w.Gender})
	}
	if len(w.WorkStatus) > 0 {
		b = b.Where(sq.Eq{FieldWorkStatus: w.WorkStatus})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	if len(w.OrganizationID) > 0 {
		b = b.Where(sq.Eq{FieldOrganizationID: w.OrganizationID})
	}
	sql, args, _ := b.ToSql()
	var count uint64
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&count, args...)
	return count, err
}

func (r *Repository) QueryPage(w *WhereParams, db *sqlx.DB) ([]*Staff, error) {
	b := sq.Select(SelectFields...).From(Table)
	if len(w.Username) > 0 {
		b = b.Where(sq.Like{FieldUsername: fmt.Sprint("%", w.Username, "%")})
	}
	if len(w.Name) > 0 {
		b = b.Where(sq.Like{FieldName: fmt.Sprint("%", w.Name, "%")})
	}
	if len(w.Email) > 0 {
		b = b.Where(sq.Like{FieldEmail: fmt.Sprint("%", w.Email, "%")})
	}
	if len(w.Mobile) > 0 {
		b = b.Where(sq.Like{FieldMobile: fmt.Sprint("%", w.Mobile, "%")})
	}
	if len(w.Gender) > 0 {
		b = b.Where(sq.Eq{FieldGender: w.Gender})
	}
	if len(w.WorkStatus) > 0 {
		b = b.Where(sq.Eq{FieldWorkStatus: w.WorkStatus})
	}
	if len(w.Remark) > 0 {
		b = b.Where(sq.Like{FieldRemark: fmt.Sprint("%", w.Remark, "%")})
	}
	if len(w.Status) > 0 {
		b = b.Where(sq.Eq{FieldStatus: w.Status})
	}
	if len(w.OrganizationID) > 0 {
		b = b.Where(sq.Eq{FieldOrganizationID: w.OrganizationID})
	}
	b = b.Limit(w.PageSize).Offset((w.Current - 1) * w.PageSize).OrderBy(FieldSort)
	sql, args, _ := b.ToSql()
	es := []*Staff{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Select(&es, args...)
	return es, err
}

func (r *Repository) QueryByUniqueField(w *WhereParams, db *sqlx.DB) (*Staff, error) {
	b := sq.Select(FieldID, FieldUsername, FieldPassword, FieldEmail, FieldMobile, FieldStatus).From(Table)
	if len(w.ID) > 0 {
		b = b.Where(sq.Eq{FieldID: w.ID})
	}
	if len(w.Username) > 0 {
		b = b.Where(sq.Eq{FieldUsername: w.Username})
	}
	if len(w.Email) > 0 {
		b = b.Where(sq.Eq{FieldEmail: w.Email})
	}
	if len(w.Mobile) > 0 {
		b = b.Where(sq.Eq{FieldMobile: w.Mobile})
	}
	sql, args, _ := b.Limit(1).ToSql()
	e := &Staff{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(e, args...)
	return e, err
}

func (r *Repository) Exist(e *Staff, db *sqlx.DB) (bool, error) {
	b := sq.Select(FieldID).Prefix("SELECT EXISTS (").From(Table)
	if len(e.OrganizationID) > 0 {
		b = b.Where(sq.Eq{FieldOrganizationID: e.OrganizationID})
	}
	b = b.Suffix(")")
	sql, args, _ := b.ToSql()
	var exist bool
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(&exist, args...)
	return exist, err
}

func (r *Repository) ValidationFields(c *Staff, db *sqlx.DB) (*Staff, error) {
	b := sq.Select(FieldID, FieldUsername, FieldEmail, FieldMobile).From(Table)
	var or []sq.Sqlizer
	if len(c.Name) > 0 {
		or = append(or, sq.Eq{FieldUsername: c.Username})
	}
	if len(c.Email) > 0 {
		or = append(or, sq.Eq{FieldEmail: c.Email})
	}
	if len(c.Mobile) > 0 {
		or = append(or, sq.Eq{FieldMobile: c.Mobile})
	}
	b = b.Where(sq.Or(or))
	if len(c.ID) > 0 {
		b = b.Where(sq.NotEq{FieldID: c.ID})
	}
	sql, args, _ := b.Limit(1).ToSql()
	e := &Staff{}
	stmt, _ := db.Preparex(sql)
	err := stmt.Get(e, args...)
	return e, err
}

func (r *Repository) CreateWithTx(e *Staff, tx *sqlx.Tx) (string, error) {
	id := xid.New().String()
	b := sq.Insert(Table).SetMap(sq.Eq{
		FieldID:             id,
		FieldUsername:       e.Username,
		FieldPassword:       utils.GeneratePassword(e.Password),
		FieldName:           e.Name,
		FieldEmail:          e.Email,
		FieldMobile:         e.Mobile,
		FieldAvatar:         e.Avatar,
		FieldGender:         e.Gender,
		FieldOrganizationID: e.OrganizationID,
		FieldWorkStatus:     e.WorkStatus,
		FieldStatus:         e.Status,
		FieldSort:           e.Sort,
		FieldRemark:         e.Remark,
		FieldCreatedAt:      time.Now().Unix(),
		FieldUpdatedAt:      time.Now().Unix(),
		FieldCreatedBy:      e.CreatedBy,
	})
	sql, args, _ := b.ToSql()
	result, err := tx.Exec(sql, args...)
	if err != nil {
		return "", err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return "", errors.New(CreatedFail)
	}
	return id, err
}

func (r *Repository) UpdateWithTx(e *UpdateInput, tx *sqlx.Tx) error {
	b := sq.Update(Table)
	if len(e.Username) > 0 {
		b = b.Set(FieldUsername, e.Username)
	}
	if len(e.Name) > 0 {
		b = b.Set(FieldName, e.Name)
	}
	if len(e.Email) > 0 {
		b = b.Set(FieldEmail, e.Email)
	}
	if len(e.Mobile) > 0 {
		b = b.Set(FieldMobile, e.Mobile)
	}
	if len(e.Avatar) > 0 {
		b = b.Set(FieldAvatar, e.Avatar)
	}
	if len(e.Gender) > 0 {
		b = b.Set(FieldGender, e.Gender)
	}
	if len(e.WorkStatus) > 0 {
		b = b.Set(FieldWorkStatus, e.WorkStatus)
	}
	if len(e.Status) > 0 {
		b = b.Set(FieldStatus, e.Status)
	}
	if e.Sort > 0 {
		b = b.Set(FieldSort, e.Sort)
	}
	if len(e.OrganizationID) > 0 {
		b = b.Set(FieldOrganizationID, e.OrganizationID)
	}
	b = b.Set(FieldRemark, e.Remark)
	b = b.SetMap(sq.Eq{
		FieldUpdatedAt: time.Now().Unix(),
		FieldUpdatedBy: e.UpdatedBy,
	}).Where(sq.Eq{FieldID: e.ID})
	sql, args, _ := b.ToSql()
	result, err := tx.Exec(sql, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New(UpdatedFail)
	}
	return err
}

func (r *Repository) DeleteWithTx(e *Staff, tx *sqlx.Tx) error {
	sql, args, _ := sq.Delete(Table).Where(sq.Eq{FieldID: e.ID}).ToSql()
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
