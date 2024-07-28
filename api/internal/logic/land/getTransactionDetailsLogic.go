package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransactionDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTransactionDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransactionDetailsLogic {
	return &GetTransactionDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTransactionDetailsLogic) GetTransactionDetails(req *types.GetTransactionDetailsReq) (resp *types.GetTransactionDetailsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
