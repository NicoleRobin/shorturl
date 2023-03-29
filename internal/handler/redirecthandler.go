package handler

import (
	"net/http"

	"github.com/nicolerobin/shorturl/internal/logic"
	"github.com/nicolerobin/shorturl/internal/svc"
	"github.com/nicolerobin/shorturl/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func redirectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RedirectReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRedirectLogic(r.Context(), svcCtx)
		resp, err := l.Redirect(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			// httpx.OkJsonCtx(r.Context(), w, resp)
			w.Header().Set("Location", resp.Url)
			httpx.WriteJsonCtx(r.Context(), w, http.StatusMovedPermanently, resp)
		}
	}
}
