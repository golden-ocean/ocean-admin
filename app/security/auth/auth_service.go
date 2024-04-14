package auth

import (
	"errors"

	"github.com/golden-ocean/ocean-admin/app/system/menu"
	"github.com/golden-ocean/ocean-admin/app/system/staff"
	"github.com/golden-ocean/ocean-admin/pkg/common/constants"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/golden-ocean/ocean-admin/pkg/utils"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db           *sqlx.DB
	staffService *staff.Service
	menuService  *menu.Service
}

func NewService() *Service {
	return &Service{
		db:           global.DB,
		staffService: staff.NewService(),
		menuService:  menu.NewService(),
	}
}

func (s *Service) Login(req *LoginInput) (*staff.Staff, error) {
	e, err := s.staffService.QueryByUniqueField(&staff.WhereParams{Username: req.Username})
	if err != nil {
		return nil, errors.New(ErrorUsernameOrPassword)
	}
	if e.Status != constants.ENABLE {
		return nil, errors.New(ErrorDisableStatus)
	}
	match := utils.ComparePasswords(e.Password, req.Password)
	if !match {
		return nil, errors.New(ErrorUsernameOrPassword)
	}
	return e, err
}

func (s *Service) Logout(id string) error {
	return nil
}

func (s *Service) QueryInfo(id string) (*InfoOutput, error) {
	// staff
	staff, err := s.staffService.QueryByUniqueField(&staff.WhereParams{ID: id})
	if err != nil {
		return nil, err
	}
	menus, permissions, err := s.menuService.QueryByStaffID(staff.ID)
	if err != nil {
		return nil, err
	}
	return &InfoOutput{
		Staff:       staff,
		Menus:       menus,
		Permissions: permissions,
	}, err
}

func (s *Service) Refresh(id string) (*staff.Staff, error) {
	e, err := s.staffService.QueryByUniqueField(&staff.WhereParams{ID: id})
	return e, err
}
