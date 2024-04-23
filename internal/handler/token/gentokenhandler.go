package token

import (
	"accessToken_go_zero/internal/model"
	"net/http"

	"accessToken_go_zero/internal/logic/token"
	"accessToken_go_zero/internal/svc"
	"accessToken_go_zero/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GenTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		mapper := model.NewTokenDetail(svcCtx.DB)
		l := token.NewGenTokenLogic(r.Context(), svcCtx, mapper)
		resp, err := l.GenToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
