package staff

import (
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/golden-ocean/ocean-admin/app/system/staff_position"
	"github.com/golden-ocean/ocean-admin/app/system/staff_role"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/golden-ocean/ocean-admin/platform/database"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Service struct {
	db                *sqlx.DB
	casbinEnforcer    *casbin.Enforcer
	staffRepo         *Repository
	staffRoleRepo     *staff_role.Repository
	staffPositionRepo *staff_position.Repository
}

func NewService() *Service {
	return &Service{
		db:                global.DB,
		casbinEnforcer:    global.CasbinEnforcer,
		staffRepo:         NewRepository(),
		staffRoleRepo:     staff_role.NewRepository(),
		staffPositionRepo: staff_position.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &Staff{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	position_ids := r.PositionIDs
	role_ids := r.RoleIDs

	err := database.WithTransaction(s.db, func(tx *sqlx.Tx) error {
		id, err := s.staffRepo.CreateWithTx(e, tx)
		if err != nil {
			return err
		}
		// 创建staff_role 关联
		if err = s.staffRoleRepo.CreateBatchWithTx(&staff_role.CreateInput{StaffID: id, RoleIDs: role_ids}, tx); err != nil {
			return err
		}
		// 创建staff_position 关联
		if err = s.staffPositionRepo.CreateBatchWithTx(&staff_position.CreateInput{StaffID: id, PositionIDs: position_ids}, tx); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	staff_id := r.ID
	new_position_ids := r.PositionIDs
	new_role_ids := r.RoleIDs
	staff_position_list, err := s.staffPositionRepo.QueryAll(&staff_position.WhereParams{StaffID: staff_id}, s.db)
	if err != nil {
		return err
	}
	db_position_ids := lo.Map(staff_position_list, func(item *staff_position.StaffPosition, _ int) string { return item.PositionID })
	staff_role_list, err := s.staffRoleRepo.QueryAll(&staff_role.WhereParams{StaffID: staff_id}, s.db)
	if err != nil {
		return err
	}
	db_role_ids := lo.Map(staff_role_list, func(item *staff_role.StaffRole, _ int) string { return item.RoleID })

	remove_p_ids, add_p_ids := lo.Difference(db_position_ids, new_position_ids)
	remove_r_ids, add_r_ids := lo.Difference(db_role_ids, new_role_ids)

	err = database.WithTransaction(s.db, func(tx *sqlx.Tx) error {
		if err := s.staffRepo.UpdateWithTx(r, tx); err != nil {
			return err
		}
		if len(remove_p_ids) > 0 {
			if err := s.staffPositionRepo.DeleteBatchWithTx(&staff_position.DeleteInput{StaffID: r.ID, PositionIDs: remove_p_ids}, tx); err != nil {
				return err
			}
		}
		if len(add_p_ids) > 0 {
			if err := s.staffPositionRepo.CreateBatchWithTx(&staff_position.CreateInput{StaffID: r.ID, PositionIDs: add_p_ids}, tx); err != nil {
				return err
			}
		}
		if len(remove_r_ids) > 0 {
			if err := s.staffRoleRepo.DeleteBatchWithTx(&staff_role.DeleteInput{StaffID: r.ID, RoleIDs: remove_r_ids}, tx); err != nil {
				return err
			}
		}
		if len(add_r_ids) > 0 {
			if err := s.staffRoleRepo.CreateBatchWithTx(&staff_role.CreateInput{StaffID: r.ID, RoleIDs: add_r_ids}, tx); err != nil {
				return err
			}

		}
		// // *************** casbin 用户角色关系 ***************************
		// 构建 p role_id menu_path menu_method
		if _, err := s.casbinEnforcer.DeleteRolesForUser(r.ID); err != nil {
			return err
		}
		if _, err := s.casbinEnforcer.AddRolesForUser(r.ID, new_role_ids); err != nil {
			return err
		}
		// // ***************⬇ casbin 用户角色关系 ⬇***************************
		return nil
	})
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	// 事务
	err := database.WithTransaction(s.db, func(tx *sqlx.Tx) error {
		// 删除 staff_role 关联
		if err := s.staffRoleRepo.DeleteBatchWithTx(&staff_role.DeleteInput{StaffID: r.ID, RoleIDs: r.RoleIDs}, tx); err != nil {
			return err
		}
		// 删除 staff_position 关联
		if err := s.staffPositionRepo.DeleteBatchWithTx(&staff_position.DeleteInput{StaffID: r.ID, PositionIDs: r.PositionIDs}, tx); err != nil {
			return err
		}
		if err := s.staffRepo.DeleteWithTx(&Staff{ID: r.ID}, tx); err != nil {
			return err
		}
		return nil
	})
	return err
}

func (s *Service) QueryByUniqueField(w *WhereParams) (*Staff, error) {
	e, err := s.staffRepo.QueryByUniqueField(w, s.db)
	return e, err
}

func (s *Service) QueryPage(w *WhereParams) ([]*StaffOutput, uint64, error) {
	total, err := s.staffRepo.QueryCount(w, s.db)
	if err != nil {
		return nil, 0, err
	}
	es, err := s.staffRepo.QueryPage(w, s.db)
	if err != nil {
		return nil, 0, err
	}
	staff_ids := lo.Map(es, func(e *Staff, _ int) string {
		return e.ID
	})
	staff_position_list, err := s.staffPositionRepo.QueryAll(&staff_position.WhereParams{StaffIDs: staff_ids}, s.db)
	if err != nil {
		return nil, 0, err
	}
	staff_role_list, err := s.staffRoleRepo.QueryAll(&staff_role.WhereParams{StaffIDs: staff_ids}, s.db)
	if err != nil {
		return nil, 0, err
	}
	output := make([]*StaffOutput, 0)
	_ = copier.Copy(&output, es)

	for _, entity := range output {
		position_ids := lo.FilterMap(staff_position_list, func(item *staff_position.StaffPosition, _ int) (string, bool) {
			if entity.ID == item.StaffID {
				return item.PositionID, true
			}
			return "", false
		})
		role_ids := lo.FilterMap(staff_role_list, func(item *staff_role.StaffRole, _ int) (string, bool) {
			if entity.ID == item.StaffID {
				return item.RoleID, true
			}
			return "", false
		})
		entity.PositionIDs = position_ids
		entity.RoleIDs = role_ids
	}

	return output, total, err
}

func (s *Service) validationFields(c *Staff) error {
	e, err := s.staffRepo.ValidationFields(c, s.db)
	if err != nil {
		return nil
	}
	if e.Username == c.Username {
		return errors.New(ErrorUsernameRepeat)
	}
	if e.Email == c.Email {
		return errors.New(ErrorEmailRepeat)
	}
	if e.Mobile == c.Mobile {
		return errors.New(ErrorMobileRepeat)
	}
	return nil
}
