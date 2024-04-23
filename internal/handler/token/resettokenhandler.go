package token

import (
	"net/http"

	"accessToken_go_zero/internal/logic/token"
	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ResetTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := token.NewResetTokenLogic(r.Context(), svcCtx)
		resp, err := l.ResetToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
