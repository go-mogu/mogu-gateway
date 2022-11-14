package config

import (
	"github.com/go-mogu/mogu-gateway/internal/consts"
	"github.com/go-mogu/mogu-gateway/internal/model"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func init() {
	routesVar := g.Cfg().MustGet(gctx.New(), consts.RoutesKey)
	if routesVar == nil {
		panic(gerror.New("not found routes !"))
	}
	err := routesVar.Scan(&model.Routes)
	if err != nil {
		panic(err)
	}
}
