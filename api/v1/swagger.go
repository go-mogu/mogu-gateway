package v1

import (
	"github.com/go-mogu/mogu-gateway/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

// SwaggerConfigReq Req
type SwaggerConfigReq struct {
	g.Meta `path:"/v3/api-docs/swagger-config" tags:"swagger" method:"get" summary:"resources"`
}

// SwaggerConfigRes Res
type SwaggerConfigRes []*model.SwaggerResources

// SwaggerResourcesReq Req
type SwaggerResourcesReq struct {
	g.Meta `path:"/" tags:"swagger" method:"get" summary:"resources"`
}

// SwaggerResourcesRes Res
type SwaggerResourcesRes []*model.SwaggerResources

// SwaggerUiReq ui Req
type SwaggerUiReq struct {
	g.Meta `path:"/configuration/ui" tags:"swagger" method:"get" summary:"ui"`
}

// SwaggerUiRes Res
type SwaggerUiRes string

// SwaggerSecurityReq security Req
type SwaggerSecurityReq struct {
	g.Meta `path:"/configuration/security" tags:"swagger" method:"get" summary:"security"`
}

// SwaggerSecurityRes Res
type SwaggerSecurityRes string
