package token

import (
	"net/http"

	"accessToken_go_zero/internal/logic/token"
	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VerifyTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := token.NewVerifyTokenLogic(r.Context(), svcCtx)
		resp, err := l.VerifyToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
