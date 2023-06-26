package blog

import (
	"blog_server/common/respx"
	logic "blog_server/internal/logic/blog"
	"blog_server/internal/svc"
	"blog_server/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ThumbsupBlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ThumbsUpBlogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewThumbsUpBlogLogic(r.Context(), svcCtx)
		resp, err, msg := l.ThumbsUpBlog(&req)
		respx.Response(w, resp, err, msg)
		//if err != nil {
		//	httpx.ErrorCtx(r.Context(), w, err)
		//} else {
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}
	}
}
