package organization

import (
	"errors"

	"github.com/golden-ocean/ocean-admin/app/system/staff"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/jinzhu/copier"
	"github.com/jmoiron/sqlx"
	"github.com/samber/lo"
)

type Service struct {
	db               *sqlx.DB
	organizationRepo *Repository
	staffRepo        *staff.Repository
}

func NewService() *Service {
	return &Service{
		db:               global.DB,
		organizationRepo: NewRepository(),
		staffRepo:        staff.NewRepository(),
	}
}

func (s *Service) Create(r *CreateInput) error {
	e := &Organization{}
	_ = copier.Copy(e, r)
	if err := s.validationFields(e); err != nil {
		return err
	}
	err := s.organizationRepo.Create(e, s.db)
	return err
}

func (s *Service) Update(r *UpdateInput) error {
	all, err := s.organizationRepo.QueryAll(&WhereParams{}, s.db)
	if err != nil {
		return err
	}
	ok := s.isValidParentID(r.ID, r.ParentID, all)
	if !ok {
		return errors.New(ErrorPidCantEqSelfAndChildId)
	}
	e := &Organization{}
	_ = copier.Copy(e, r)
	err = s.organizationRepo.Update(e, s.db)
	return err
}

func (s *Service) Delete(r *DeleteInput) error {
	// 检查 organization parent_id 关联
	if exist, err := s.organizationRepo.Exist(&Organization{ParentID: r.ID}, s.db); err != nil {
		return err
	} else if exist {
		return errors.New(ErrorExistChildren)
	}
	// 检查 staff organization 关联
	if exist, err := s.staffRepo.Exist(&staff.Staff{OrganizationID: r.ID}, s.db); err != nil {
		return err
	} else if exist {
		return errors.New(ErrorExistStaff)
	}
	err := s.organizationRepo.Delete(&Organization{ID: r.ID}, s.db)
	return err
}

func (s *Service) QueryTree(w *WhereParams) ([]*OrganizationOutput, error) {
	es, err := s.organizationRepo.QueryAll(w, s.db)
	if err != nil {
		return nil, err
	}
	filter := lo.Filter(es, func(item *Organization, index int) bool {
		return item.ID != w.ID
	})
	output := make([]*OrganizationOutput, 0)
	_ = copier.Copy(&output, filter)
	tree := s.buildTree(output)
	return tree, err
}

func (s *Service) buildTree(organizations []*OrganizationOutput) []*OrganizationOutput {
	orgMap := make(map[string]*OrganizationOutput)
	// 将组织映射到一个字典中
	for _, org := range organizations {
		orgMap[org.ID] = org
	}
	// 将子组织连接到父组织
	for _, org := range organizations {
		if parent, ok := orgMap[org.ParentID]; ok {
			parent.Children = append(parent.Children, org)
		}
	}
	// 寻找根组织并返回
	var roots []*OrganizationOutput
	for _, org := range organizations {
		if len(org.ParentID) == 0 {
			roots = append(roots, org)
		}
	}
	return roots
}

func (s *Service) isValidParentID(id, parentID string, organizations []*Organization) bool {
	if id == parentID {
		return false
	}
	// 创建一个集合，用于存储组织ID的所有祖先
	ancestors := make(map[string]struct{})
	// 将所有组织以ID为键，组织对象为值存储到一个映射中，以便快速访问
	orgMap := make(map[string]*Organization)
	for _, org := range organizations {
		orgMap[org.ID] = org
	}
	// 递归地查找组织的所有祖先
	for len(parentID) > 0 {
		// 检查组织的父级ID是否在祖先集合中
		if _, ok := ancestors[parentID]; ok {
			// 如果父级ID已经在祖先集合中，说明存在循环引用，返回 false
			return false
		}
		// 将父级ID添加到祖先集合中
		ancestors[parentID] = struct{}{}
		// 获取父级组织
		parentOrg, ok := orgMap[parentID]
		if !ok {
			// 如果找不到父级组织，则说明 parentID 是无效的
			return false
		}
		// 检查父级组织是否与给定ID相同，如果相同，说明 ParentID 是自己的子孙
		if parentOrg.ID == id {
			return false
		}
		// 继续向上查找父级组织
		parentID = parentOrg.ParentID
	}
	// 如果以上条件都不满足，则说明 ParentID 是有效的
	return true
}

func (s *Service) validationFields(c *Organization) error {
	// name 唯一  code 唯一
	e, err := s.organizationRepo.ValidationFields(c, s.db)
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
