// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	blog "blog_backend/internal/handler/blog"
	"blog_backend/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/list/blog",
				Handler: blog.ListBlogHandler(serverCtx),
			},
		},
	)
}