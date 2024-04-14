package role

// import (
// 	"errors"

// 	"github.com/casbin/casbin/v2"
// 	"github.com/golden-ocean/ocean-admin/app/system/role_menu"
// 	"github.com/golden-ocean/ocean-admin/app/system/staff_role"
// 	"github.com/golden-ocean/ocean-admin/pkg/common/global"
// 	"github.com/golden-ocean/ocean-admin/platform/database"
// 	"github.com/jinzhu/copier"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/rs/xid"
// 	"github.com/samber/lo"
// )

// type Service struct {
// 	db             *sqlx.DB
// 	roleRepo       *Repository
// 	roleMenuRepo   *role_menu.Repository
// 	staffRoleRepo  *staff_role.Repository
// 	casbinEnforcer *casbin.Enforcer
// }

// func NewService() *Service {
// 	return &Service{
// 		db:             global.DB,
// 		roleRepo:       NewRepository(),
// 		roleMenuRepo:   role_menu.NewRepository(),
// 		staffRoleRepo:  staff_role.NewRepository(),
// 		casbinEnforcer: global.CasbinEnforcer,
// 	}
// }

// func (s *Service) Create(req *CreateInput) error {
// 	e := &Role{}
// 	_ = copier.Copy(e, req)
// 	if err := s.validationFields(e); err != nil {
// 		return err
// 	}
// 	err := s.roleRepo.Create(e, s.db)
// 	return err
// }

// func (s *Service) Update(req *UpdateInput) error {
// 	e := &Role{}
// 	_ = copier.Copy(e, req)
// 	if err := s.validationFields(e); err != nil {
// 		return err
// 	}
// 	err := s.roleRepo.Update(e, s.db)
// 	return err
// }

// func (s *Service) Delete(req *DeleteInput) error {
// 	// 检查 staff_role 的关联
// 	if exist, err := s.staffRoleRepo.Exist(&staff_role.StaffRole{RoleID: req.ID}, s.db); err != nil {
// 		return err
// 	} else if exist {
// 		return errors.New(ErrorExistStaffs)
// 	}
// 	// 检查 role_menu 的关联
// 	if exist, err := s.roleMenuRepo.Exist(&role_menu.RoleMenu{RoleID: req.ID}, s.db); err != nil {
// 		return err
// 	} else if exist {
// 		return errors.New(ErrorExistMenus)
// 	}
// 	err := s.roleRepo.Delete(&Role{ID: req.ID}, s.db)
// 	return err
// }

// func (s *Service) QueryPage(w *WhereParams) ([]*RoleOutput, uint64, error) {
// 	total, err := s.roleRepo.QueryCount(w, s.db)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	es, err := s.roleRepo.QueryPage(w, s.db)
// 	output := make([]*RoleOutput, 0)
// 	_ = copier.Copy(&output, es)
// 	return output, total, err
// }

// func (s *Service) QueryAll(w *WhereParams) ([]*RoleOutput, error) {
// 	es, err := s.roleRepo.QueryAll(w, s.db)
// 	output := make([]*RoleOutput, 0)
// 	_ = copier.Copy(&output, es)
// 	return output, err
// }

// func (s *Service) QueryMenus(w *RoleMenuInput) ([]string, error) {
// 	ids, err := s.roleMenuRepo.QueryMenuIDs(&role_menu.WhereParams{RoleID: w.RoleID}, s.db)
// 	// output := make([]string, 0)
// 	// _ = copier.Copy(&output, ids)
// 	return ids, err
// }

// func (s *Service) GrantMenus(req *RoleMenuInput) error {
// 	new_menu_ids := req.MenuIDs
// 	db_menu_ids, err := s.roleMenuRepo.QueryMenuIDs(&role_menu.WhereParams{RoleID: req.RoleID}, s.db)
// 	if err != nil {
// 		return err
// 	}
// 	remove_ids, add_ids := lo.Difference(db_menu_ids, new_menu_ids)
// 	err = database.WithTransaction(s.db, func(tx *sqlx.Tx) error {
// 		if len(remove_ids) > 0 {
// 			if err := s.roleMenuRepo.DeleteBatchWithTx(&role_menu.DeleteInput{RoleID: req.RoleID, MenuIDs: remove_ids}, tx); err != nil {
// 				return err
// 			}
// 		}
// 		if len(add_ids) > 0 {
// 			if err := s.roleMenuRepo.CreateBatchWithTx(&role_menu.CreateInput{RoleID: req.RoleID, MenuIDs: add_ids}, tx); err != nil {
// 				return err
// 			}
// 		}
// 		// // ***************************** casbin role_menu ********************************
// 		// menus, err := s.menuRepo.QueryAllByIDs(new_menu_ids, s.db)
// 		// if err != nil {
// 		// 	return err
// 		// }
// 		// // 构建 p role_id menu_path menu_method
// 		// rules := make([][]string, 0)
// 		// for _, m_id := range new_menu_ids {
// 		// 	for _, m := range menus {
// 		// 		if m.ID == m_id {
// 		// 			rules = append(rules, []string{req.RoleID.String(), m.Path, m.Method})
// 		// 		}
// 		// 	}
// 		// }
// 		// global.CasbinEnforcer.DeletePermissionsForUser(req.RoleID.String())
// 		// global.CasbinEnforcer.AddNamedPolicies("p", rules)
// 		// // ********************************************************************************
// 		return nil
// 	})
// 	return err
// }

// func (s *Service) validationFields(c *Role) error {
// 	// name 唯一  code 唯一
// 	e, err := s.roleRepo.ValidationFields(c, s.db)
// 	if err != nil {
// 		return nil
// 	}
// 	if e.Name == c.Name {
// 		return errors.New(ErrorNameRepeat)
// 	}
// 	if e.Code == c.Code {
// 		return errors.New(ErrorCodeRepeat)
// 	}
// 	return nil
// }
