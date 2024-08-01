package land

import (
	"context"
	"github.com/google/uuid"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTransactionIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTransactionIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTransactionIdLogic {
	return &GetTransactionIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTransactionIdLogic) GetTransactionId() (resp *types.GetTransactionIdResp, err error) {

	return &types.GetTransactionIdResp{
		Code: 0,
		Data: int64(uuid.New().ID()),
		Msg:  "",
	}, nil
}
