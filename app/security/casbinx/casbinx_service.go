package casbinx

import (
	"github.com/casbin/casbin/v2"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db       *sqlx.DB
	enforcer *casbin.Enforcer
}

func NewService() *Service {
	return &Service{
		db:       global.DB,
		enforcer: &casbin.Enforcer{},
	}
}

// func (s *Service) Delete(req *DeleteInput) error {
// 	// ("p", "eve", "data3", "read")
// 	// xx := [][]string{
// 	// 	[]string{"jack", "data4", "read"},
// 	// 	[]string{"katy", "data4", "write"},
// 	// 	[]string{"leyo", "data4", "read"},
// 	// 	[]string{"ham", "data4", "write"},
// 	// }
// 	rules := [][]string
// 	s.enforcer.RemovePolicy(req.Username, req.Path, req.Method)
// }
