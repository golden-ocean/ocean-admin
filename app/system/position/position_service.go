package position

import (
	"errors"

	"github.com/golden-ocean/ocean-admin/app/system/staff_position"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db                *sqlx.DB
	positionRepo      *Repository
	staffPositionRepo *staff_position.Repository
}

func NewService() *Service {
	return &Service{
		db:                global.DB,
		positionRepo:      NewRepository(),
		staffPositionRepo: staff_position.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &Position{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.positionRepo.Create(e, s.db)
	return err
}
func (s *Service) Update(r *UpdateInput) error {
	e := &Position{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.positionRepo.Update(e, s.db)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	// 检查 staff_position 关联
	if exist, err := s.staffPositionRepo.Exist(&staff_position.StaffPosition{PositionID: r.ID}, s.db); err != nil {
		return err
	} else if exist {
		return errors.New(ErrorExistStaffs)
	}
	err := s.positionRepo.Delete(&Position{ID: r.ID}, s.db)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*PositionOutput, uint64, error) {
	total, err := s.positionRepo.QueryCount(w, s.db)
	if err != nil {
		return nil, 0, err
	}
	es, err := s.positionRepo.QueryPage(w, s.db)
	output := make([]*PositionOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) QueryAll(w *WhereParams) ([]*PositionOutput, error) {
	es, err := s.positionRepo.QueryAll(w, s.db)
	output := make([]*PositionOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, err
}

func (s *Service) validationFields(c *Position) error {
	e, err := s.positionRepo.ValidationFields(c, s.db)
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
