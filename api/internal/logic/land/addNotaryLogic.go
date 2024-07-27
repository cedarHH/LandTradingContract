package land

import (
	"context"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type AddNotaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddNotaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNotaryLogic {
	return &AddNotaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddNotaryLogic) AddNotary(req *types.AddNotaryReq) (
	resp *types.AddNotaryResp, err error) {

	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.NotaryAddress)
	tx, err := l.svcCtx.Conn.AddNotary(auth, common.HexToAddress(req.NotaryAddress))
	if err != nil {
		log.Fatalf("Failed to call AddNotary: %v", err)
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())
	return &types.AddNotaryResp{
		Code: 0,
		Msg:  "add notary success",
	}, nil
}
