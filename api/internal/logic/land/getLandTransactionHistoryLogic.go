package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLandTransactionHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLandTransactionHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLandTransactionHistoryLogic {
	return &GetLandTransactionHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLandTransactionHistoryLogic) GetLandTransactionHistory(req *types.GetLandTransactionHistoryReq) (resp *types.GetLandTransactionHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
