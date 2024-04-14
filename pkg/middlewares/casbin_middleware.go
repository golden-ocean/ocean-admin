package middlewares

import (
	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	casbinMiddleware "github.com/gofiber/contrib/casbin"
	"github.com/gofiber/fiber/v2"
	"github.com/golden-ocean/ocean-admin/pkg/common/global"
	"github.com/golden-ocean/ocean-admin/pkg/utils"
	"github.com/jmoiron/sqlx"
)

func CasbinProtected() func(*fiber.Ctx) error {
	config := casbinMiddleware.Config{
		// ModelFilePath: "./pkg/configs/rbac_model.conf",
		// PolicyAdapter: casbinAdapter(global.DB),
		// Enforcer:     Enforcer(global.DB),
		Enforcer:     Enforcer(global.DB),
		Lookup:       lookup,
		Unauthorized: unauthorized,
		Forbidden:    forbidden,
	}
	// 按路由规则匹配权限
	return casbinMiddleware.New(config).RoutePermission()
}

var modelConf = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && keyMatch2(r.obj, p.obj) && r.act == p.act || r.sub == "co57jisvg7l30a15aje0"
`

func Enforcer(db *sqlx.DB) *casbin.Enforcer {
	a := casbinAdapter(db)
	// a := casbinx.NewAdapter()
	m, err := model.NewModelFromString(modelConf)
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer(m, a)
	if err != nil {
		panic(err)
	}
	global.CasbinEnforcer = e
	return e
}

func casbinAdapter(db *sqlx.DB) *sqlxadapter.Adapter {
	a, err := sqlxadapter.NewAdapter(db, "casbin_rule")
	if err != nil {
		panic(err)
	}
	return a
}

func lookup(c *fiber.Ctx) string {
	claims, _ := utils.ExtractTokenMetadata(c)
	return claims.ID
}

func unauthorized(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusUnauthorized, "没有权限！")
}

func forbidden(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusForbidden, "禁止访问！")
}
