package dictionary

import (
	"errors"

	"github.com/golden-ocean/ocean-admin/app/system/dictionary_item"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db                 *sqlx.DB
	dictionaryRepo     *Repository
	dictionaryItemRepo *dictionary_item.Repository
}

func NewService() *Service {
	return &Service{
		db:                 global.DB,
		dictionaryRepo:     NewRepository(),
		dictionaryItemRepo: dictionary_item.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &Dictionary{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.dictionaryRepo.Create(e, s.db)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	e := &Dictionary{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.dictionaryRepo.Update(e, s.db)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	if exist, err := s.dictionaryItemRepo.Exist(&dictionary_item.DictionaryItem{DictionaryID: r.ID}, s.db); err != nil {
		return err
	} else if exist {
		return errors.New(ErrorExistChildren)
	}
	err := s.dictionaryRepo.Delete(&Dictionary{ID: r.ID}, s.db)
	return err
}

func (s *Service) QueryPage(w *WhereParams) ([]*DictionaryOutput, uint64, error) {
	total, err := s.dictionaryRepo.QueryCount(w, s.db)
	if err != nil {
		return nil, 0, err
	}
	es, err := s.dictionaryRepo.QueryPage(w, s.db)
	output := make([]*DictionaryOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, total, err
}

func (s *Service) validationFields(c *Dictionary) error {
	e, err := s.dictionaryRepo.ValidationFields(c, s.db)
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

func (s *Service) QueryByCode(code string) ([]*dictionary_item.DictionaryItemOutput, error) {
	es, err := s.dictionaryRepo.QueryByCode(code, s.db)
	output := make([]*dictionary_item.DictionaryItemOutput, 0, 10)
	_ = copier.Copy(&output, es)
	return output, err
}

// func (s *Service) QueryByCode(code string) ([]*dictionary_item.DictionaryItemOutput, error) {
// 	d, err := s.dictionaryRepo.QueryByUniqueField(&WhereParams{Code: code}, s.db)
// 	if err != nil {
// 		return nil, err
// 	}
// 	es, err := s.dictionaryItemRepo.QueryAll(&dictionary_item.WhereParams{DictionaryID: d.ID}, s.db)
// 	if err != nil {
// 		return nil, err
// 	}
// 	output := make([]*dictionary_item.DictionaryItemOutput, 0)
// 	_ = copier.Copy(&output, es)
// 	return output, err
// }
