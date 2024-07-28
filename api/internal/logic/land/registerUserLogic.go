package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type registerUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewregisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *registerUserLogic {
	return &registerUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *registerUserLogic) registerUser(req *types.registerUserReq) (resp *types.registerUserResp, err error) {
	
	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.OwnerAddress)
    tx, err := l.svcCtx.Conn.RegisterUser(auth, common.HexToAddress(req.UserAddress), req.UserName)
    if err != nil {
        log.Fatalf("Failed to call RegisterUser: %v", err)
    }

    log.Printf("Transaction hash: %s", tx.Hash().Hex())
    return &types.RegisterUserResp{
        Code: 0,
        Msg:  "User registered successfully",
    }, nil
}
