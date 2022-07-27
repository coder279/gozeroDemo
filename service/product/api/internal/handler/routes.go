// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"gozeroDemo/service/product/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/product/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/update",
				Handler: UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/product/remove",
				Handler: RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/product/detail",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/product/list",
				Handler: ListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
