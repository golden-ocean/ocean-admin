package global

import (
	"context"

	sqlxadapter "github.com/Blank-Xu/sqlx-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/storage/redis/v3"
	"github.com/jmoiron/sqlx"
)

var (
	DB             *sqlx.DB
	Storage        *redis.Storage
	CasbinEnforcer *casbin.Enforcer
	CasbinAdapter  *sqlxadapter.Adapter
	Ctx            context.Context
)
