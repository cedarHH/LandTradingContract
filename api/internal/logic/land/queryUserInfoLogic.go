package land

import (
	"context"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryUserInfoLogic {
	return &QueryUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryUserInfoLogic) QueryUserInfo(
	req *types.QueryUserInfoReq) (resp *types.QueryUserInfoResp, err error) {

	callOpts := &bind.CallOpts{
		From: common.HexToAddress(req.Address),
	}
	info, err := l.svcCtx.Conn.QueryUserInfo(callOpts)
	if err != nil {
		log.Fatalf("Failed to call queryUserInfo: %v", err)
	}

	int64s := make([]int64, len(info.TransactionIdList))
	for i, bigInt := range info.TransactionIdList {
		int64s[i] = bigInt.Int64()
	}

	return &types.QueryUserInfoResp{
		Code: 0,
		Data: types.UserDetail{
			Username:          info.Username,
			LandIdList:        info.LandIdList,
			TransactionIdList: int64s,
		},
		Msg: "",
	}, nil
}
