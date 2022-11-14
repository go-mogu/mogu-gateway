package controller

import (
	"context"
	v1 "github.com/go-mogu/mogu-gateway/api/v1"
	"github.com/go-mogu/mogu-gateway/internal/consts"
	"github.com/go-mogu/mogu-gateway/internal/model"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Swagger = cSwagger{}
	apiUrl  = "/v3/api-docs"
)

type cSwagger struct{}

// SwaggerResources resources
func (c *cSwagger) SwaggerResources(ctx context.Context, req *v1.SwaggerResourcesReq) (res v1.SwaggerResourcesRes, err error) {
	r := g.RequestFromCtx(ctx)
	res = make(v1.SwaggerResourcesRes, 0)
	for _, route := range model.Routes {
		uriInfo, _ := gurl.ParseURL(route.Uri, -1)
		name := uriInfo["host"]
		url := consts.PathSeparator + name + apiUrl
		res = append(res, &model.SwaggerResources{
			Name:           name,
			URL:            url,
			SwaggerVersion: "3.0.0",
			Location:       url,
		})

	}
	r.Response.WriteJson(res)
	return
}

// SwaggerUi ui
func (c *cSwagger) SwaggerUi(ctx context.Context, req *v1.SwaggerUiReq) (res v1.SwaggerUiRes, err error) {
	r := g.RequestFromCtx(ctx)
	res = `{"deepLinking":true,"displayOperationId":false,"defaultModelsExpandDepth":1,"defaultModelExpandDepth":1,"defaultModelRendering":"example","displayRequestDuration":false,"docExpansion":"none","filter":false,"operationsSorter":"alpha","showExtensions":false,"showCommonExtensions":false,"tagsSorter":"alpha","validatorUrl":"","supportedSubmitMethods":["get","put","post","delete","options","head","patch","trace"],"swaggerBaseUiUrl":""}`
	r.Response.Write(res)
	return
}

// SwaggerSecurity ui
func (c *cSwagger) SwaggerSecurity(ctx context.Context, req *v1.SwaggerSecurityReq) (res v1.SwaggerSecurityRes, err error) {
	r := g.RequestFromCtx(ctx)
	res = `{}`
	r.Response.Write(res)
	return
}
