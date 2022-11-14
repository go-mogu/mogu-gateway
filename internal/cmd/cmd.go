package cmd

import (
	"context"
	"github.com/go-mogu/mogu-gateway/internal/consts"
	"github.com/go-mogu/mogu-gateway/internal/controller"
	"github.com/go-mogu/mogu-gateway/internal/core/config"
	"github.com/go-mogu/mogu-gateway/internal/core/lb"
	"github.com/go-mogu/mogu-gateway/internal/model"
	utils "github.com/go-mogu/mogu-gateway/utility"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/text/gstr"
	"net/http"
	"net/http/httputil"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	appName = g.Cfg().MustGet(gctx.New(), consts.AppNameKey).String()
	Main    = gcmd.Command{
		Name:  appName,
		Usage: appName,
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server(appName)
			s.Use(ghttp.MiddlewareCORS)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(controller.Actuator)
				s.Group("/swagger-resources", func(group *ghttp.RouterGroup) {
					group.Bind(controller.Swagger)
				})
				for _, route := range model.Routes {
					RegisterRoute(ctx, route, s)
				}
			})

			err = config.RegisterInstance(ctx, s)
			if err != nil {
				return err
			}
			s.Run()
			return nil
		},
	}
)

func RegisterRoute(ctx context.Context, route *model.Route, s *ghttp.Server) {
	uriInfo, err := gurl.ParseURL(route.Uri, -1)
	utils.ErrIsNil(ctx, err, "uri解析失败")
	serviceName := uriInfo["host"]
	if uriInfo["scheme"] == "lb" {
		s.BindHandler("/"+serviceName+"/*any", func(r *ghttp.Request) {
			ReverseProxy(serviceName, r)
		})
	}

}

func ReverseProxy(serviceName string, r *ghttp.Request) {
	g.Log().Infof(r.Context(), "请求地址：%v", r.URL)
	serviceHost := lb.GetServiceUrl(r.Context(), serviceName)
	var target = r.URL
	target.Host = serviceHost
	target.Scheme = "http"
	director := func(req *http.Request) {
		pathList := gstr.Split(target.Path, consts.PathSeparator)
		target.Path = consts.PathSeparator + gstr.Join(pathList[2:], consts.PathSeparator)
		req.URL = target
		req.Host = target.Host
		g.Log().Infof(r.Context(), "代理转发地址：%v", req.URL)
	}
	proxy := &httputil.ReverseProxy{Director: director}
	proxy.ServeHTTP(r.Response.Writer, r.Request)
}
