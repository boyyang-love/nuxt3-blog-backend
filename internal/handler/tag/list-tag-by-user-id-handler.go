package tag

import (
	"net/http"

	"blog_backend/internal/logic/tag"
	"blog_backend/internal/svc"
	"blog_backend/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListTagByUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListTagUserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tag.NewListTagByUserIdLogic(r.Context(), svcCtx)
		resp, err := l.ListTagByUserId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
