package handler

import (
	"net/http"

	"github.com/nicolerobin/shorturl/internal/logic"
	"github.com/nicolerobin/shorturl/internal/svc"
	"github.com/nicolerobin/shorturl/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func generateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGenerateLogic(r.Context(), svcCtx)
		resp, err := l.Generate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
