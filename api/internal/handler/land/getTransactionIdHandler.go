package land

import (
	"net/http"

	"github.com/cedarHH/LandTradingContract/api/internal/logic/land"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetTransactionIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := land.NewGetTransactionIdLogic(r.Context(), svcCtx)
		resp, err := l.GetTransactionId()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
