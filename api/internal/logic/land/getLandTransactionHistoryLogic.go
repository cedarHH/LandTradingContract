package land

import (
	"context"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (l *GetLandTransactionHistoryLogic) GetLandTransactionHistory(
	req *types.GetLandTransactionHistoryReq) (resp *types.GetLandTransactionHistoryResp, err error) {

	history, err := l.svcCtx.Conn.GetLandTransactionHistory(&bind.CallOpts{}, req.LandId)
	if err != nil {
		return &types.GetLandTransactionHistoryResp{
			Code: 503,
			Data: []int64{},
			Msg:  "error",
		}, nil
	}

	int64s := make([]int64, len(history))
	for i, bigInt := range history {
		int64s[i] = bigInt.Int64()
	}
	return &types.GetLandTransactionHistoryResp{
		Code: 0,
		Data: int64s,
		Msg:  "",
	}, nil
}
