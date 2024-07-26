package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransactionHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTransactionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransactionHistoryLogic {
	return &GetTransactionHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTransactionHistoryLogic) GetTransactionHistory(req *types.GetTransactionHistoryReq) (resp *types.GetTransactionHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
