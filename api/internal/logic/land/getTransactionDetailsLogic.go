package land

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"

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

func (l *GetTransactionDetailsLogic) GetTransactionDetails(
	req *types.GetTransactionDetailsReq) (resp *types.GetTransactionDetailsResp, err error) {

	detail, err := l.svcCtx.Conn.GetTransactionDetails(&bind.CallOpts{}, big.NewInt(int64(req.TransactionId)))
	if err != nil {
		return &types.GetTransactionDetailsResp{
			Code: 503,
			Data: types.TransactionDetail{},
			Msg:  "error",
		}, nil
	}

	return &types.GetTransactionDetailsResp{
		Code: 0,
		Data: types.TransactionDetail{
			LandId:    detail.LandId,
			From:      detail.From.String(),
			To:        detail.To.String(),
			Timestamp: detail.Timestamp.Uint64(),
		},
		Msg: "",
	}, nil
}
