package casbinx

import (
	"fmt"

	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"github.com/golden-ocean/ocean-admin/app/system/menu"
	"github.com/golden-ocean/ocean-admin/app/system/role"
	"github.com/golden-ocean/ocean-admin/app/system/role_menu"
	"github.com/golden-ocean/ocean-admin/app/system/staff"
	"github.com/golden-ocean/ocean-admin/app/system/staff_role"
	"github.com/golden-ocean/ocean-admin/pkg/common/constants"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Adapter struct {
	db                  *sqlx.DB
	staffRepository     *staff.Repository
	staffRoleRepository *staff_role.Repository
	roleRepository      *role.Repository
	roleMenuRepository  *role_menu.Repository
	menuRepository      *menu.Repository
}

func NewAdapter() *Adapter {
	return &Adapter{
		db:                  global.DB,
		staffRepository:     staff.NewRepository(),
		staffRoleRepository: staff_role.NewRepository(),
		roleRepository:      role.NewRepository(),
		roleMenuRepository:  role_menu.NewRepository(),
		menuRepository:      menu.NewRepository(),
	}
}

// LoadPolicy loads policy from database.
func (a *Adapter) LoadPolicy(model model.Model) error {
	err := a.loadRolePolicy(model)
	if err != nil {
		//a.logger.Zap.Errorf("Load casbin role policy error: %s", err.Error())
		return err
	}

	err = a.loadUserPolicy(model)
	if err != nil {
		//a.logger.Zap.Errorf("Load casbin user policy error: %s", err.Error())
		return err
	}

	return nil
}

// SavePolicy saves policy to database.
func (a *Adapter) SavePolicy(model model.Model) error {
	return nil
}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	//TODO implement me
	panic("implement me")
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	//TODO implement me
	panic("implement me")
}

// load role policy (p,role_id,path,method)
func (a *Adapter) loadRolePolicy(model model.Model) error {
	roles_menus, err := a.roleMenuRepository.QueryAll(&role_menu.WhereParams{}, a.db)
	if err != nil {
		return err
	}
	menu_ids := lo.Map(roles_menus, func(rm *role_menu.RoleMenu, _ int) string {
		return rm.MenuID
	})
	menus, err := a.menuRepository.QueryAllByIDs(menu_ids, a.db)
	if err != nil {
		return err
	}
	// 构建 p role_id path method
	for _, rm := range roles_menus {
		for _, m := range menus {
			if m.ID == rm.MenuID && m.Type == constants.Button {
				line := fmt.Sprintf("p,%s,%s,%s", rm.RoleID, m.Path, m.Method)
				err := persist.LoadPolicyLine(line, model)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (a *Adapter) loadUserPolicy(model model.Model) error {
	staffs_roles, err := a.staffRoleRepository.QueryAll(&staff_role.WhereParams{}, a.db)
	if err != nil {
		return err
	}
	// 构建 g staff_id role_id
	for _, staff_role := range staffs_roles {
		line := fmt.Sprintf("g,%s,%s", staff_role.StaffID, staff_role.RoleID)
		err := persist.LoadPolicyLine(line, model)
		if err != nil {
			return err
		}
	}
	return nil
}
