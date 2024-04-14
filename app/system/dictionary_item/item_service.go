package dictionary_item

import (
	"errors"

	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db       *sqlx.DB
	itemRepo *Repository
}

func NewService() *Service {
	return &Service{
		db:       global.DB,
		itemRepo: NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &DictionaryItem{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.itemRepo.Create(e, s.db)
	return err
}
func (s *Service) Update(r *UpdateInput) error {
	e := &DictionaryItem{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.itemRepo.Update(e, s.db)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	err := s.itemRepo.Delete(&DictionaryItem{ID: r.ID}, s.db)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*DictionaryItemOutput, uint64, error) {
	total, err := s.itemRepo.QueryCount(w, s.db)
	if err != nil {
		return nil, 0, err
	}
	es, err := s.itemRepo.QueryPage(w, s.db)
	output := make([]*DictionaryItemOutput, 0)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) validationFields(c *DictionaryItem) error {
	e, err := s.itemRepo.ValidationFields(c, s.db)
	if err != nil {
		return nil
	}
	if e.Label == c.Label {
		return errors.New(ErrorLabelRepeat)
	}
	if e.Value == c.Value {
		return errors.New(ErrorValueRepeat)
	}
	return nil
}
