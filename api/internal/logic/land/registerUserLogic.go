package land

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterUserLogic) RegisterUser(req *types.RegisterUserReq) (
	resp *types.RegisterUserResp, err error) {

	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.Senderkey)
	userAddress := common.HexToAddress(req.UserAddress)
	tx, err := l.svcCtx.Conn.RegisterUser(auth, userAddress, req.UserName)
	if err != nil {
		log.Fatalf("Failed to call register user: %v", err)
	}
	log.Printf("Transaction hash: %s", tx.Hash().Hex())
	return &types.RegisterUserResp{
		Code: 0,
		Msg:  "Success",
	}, nil
}
