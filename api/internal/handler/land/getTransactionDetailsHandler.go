package land

import (
	"net/http"

	"github.com/cedarHH/LandTradingContract/api/internal/logic/land"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTransactionDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTransactionDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := land.NewGetTransactionDetailsLogic(r.Context(), svcCtx)
		resp, err := l.GetTransactionDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
