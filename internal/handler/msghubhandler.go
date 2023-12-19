package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"msghub/internal/logic"
	"msghub/internal/svc"
	"msghub/internal/types"
)

func MsghubHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewMsghubLogic(r.Context(), svcCtx)
		resp, err := l.Msghub(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
