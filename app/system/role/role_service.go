package role

import (
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/golden-ocean/ocean-admin/app/system/role_menu"
	"github.com/golden-ocean/ocean-admin/app/system/staff_role"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/golden-ocean/ocean-admin/platform/database"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Service struct {
	db             *sqlx.DB
	roleRepo       *Repository
	roleMenuRepo   *role_menu.Repository
	staffRoleRepo  *staff_role.Repository
	casbinEnforcer *casbin.Enforcer
}

func NewService() *Service {
	return &Service{
		db:             global.DB,
		roleRepo:       NewRepository(),
		roleMenuRepo:   role_menu.NewRepository(),
		staffRoleRepo:  staff_role.NewRepository(),
		casbinEnforcer: global.CasbinEnforcer,
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &Role{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.roleRepo.Create(e, s.db)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	e := &Role{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.roleRepo.Update(e, s.db)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	// 检查 staff_role 的关联
	if exist, err := s.staffRoleRepo.Exist(&staff_role.StaffRole{RoleID: r.ID}, s.db); err != nil {
		return err
	} else if exist {
		return errors.New(ErrorExistStaffs)
	}
	// 检查 role_menu 的关联
	if exist, err := s.roleMenuRepo.Exist(&role_menu.RoleMenu{RoleID: r.ID}, s.db); err != nil {
		return err
	} else if exist {
		return errors.New(ErrorExistMenus)
	}
	err := s.roleRepo.Delete(&Role{ID: r.ID}, s.db)
	// ***************************** casbin role_menu ********************************
	if _, err := s.casbinEnforcer.DeleteRole(r.ID); err != nil {
		return err
	}
	// *******************************************************************************
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*RoleOutput, uint64, error) {
	total, err := s.roleRepo.QueryCount(w, s.db)
	if err != nil {
		return nil, 0, err
	}
	es, err := s.roleRepo.QueryPage(w, s.db)
	output := make([]*RoleOutput, 0)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) QueryAll(w *WhereParams) ([]*RoleOutput, error) {
	es, err := s.roleRepo.QueryAll(w, s.db)
	output := make([]*RoleOutput, 0)
	_ = copier.Copy(&output, es)
	return output, err
}

func (s *Service) QueryMenus(w *RoleMenuInput) ([]string, error) {
	ids, err := s.roleMenuRepo.QueryMenuIDs(&role_menu.WhereParams{RoleID: w.RoleID}, s.db)
	return ids, err
}

func (s *Service) GrantMenus(r *RoleMenuInput) error {
	new_menu_ids := r.MenuIDs
	db_menu_ids, err := s.roleMenuRepo.QueryMenuIDs(&role_menu.WhereParams{RoleID: r.RoleID}, s.db)
	if err != nil {
		return err
	}
	remove_ids, add_ids := lo.Difference(db_menu_ids, new_menu_ids)
	err = database.WithTransaction(s.db, func(tx *sqlx.Tx) error {
		if len(remove_ids) > 0 {
			if err := s.roleMenuRepo.DeleteBatchWithTx(&role_menu.DeleteInput{RoleID: r.RoleID, MenuIDs: remove_ids}, tx); err != nil {
				return err
			}
		}
		if len(add_ids) > 0 {
			if err := s.roleMenuRepo.CreateBatchWithTx(&role_menu.CreateInput{RoleID: r.RoleID, MenuIDs: add_ids}, tx); err != nil {
				return err
			}
		}
		// ***************************** casbin role_menu ********************************
		rules := make([][]string, 0)
		for _, r := range r.Rules {
			if len(r.Method) > 0 {
				rules = append(rules, []string{r.Path, r.Method})
			}
		}
		if _, err := s.casbinEnforcer.DeletePermissionForUser(r.RoleID); err != nil {
			return err
		}

		if _, err := s.casbinEnforcer.AddPermissionsForUser(r.RoleID, rules...); err != nil {
			return err
		}
		// ********************************************************************************
		return nil
	})
	return err
}

func (s *Service) validationFields(c *Role) error {
	e, err := s.roleRepo.ValidationFields(c, s.db)
	if err != nil {
		return nil
	}
	if e.Name == c.Name {
		return errors.New(ErrorNameRepeat)
	}
	if e.Code == c.Code {
		return errors.New(ErrorCodeRepeat)
	}
	return nil
}
