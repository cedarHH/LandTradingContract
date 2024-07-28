package land

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/cedarHH/LandTradingContract/api/internal/oracle"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	OracleService *oracle.OracleService
}

func NewRegisterLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLandLogic {
	return &RegisterLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		oracleService: svcCtx.OracleService, 
	}
}

func (l *RegisterLandLogic) RegisterLand(req *types.RegisterLandReq) (resp *types.RegisterLandResp, err error) {
	// todo: add your logic here and delete this line
    auth := l.svcCtx.AccountAuth.GetAccountAuth(req.UserAddress)
    
    tx, err := l.svcCtx.Conn.registerLand(auth, big.NewInt(int64(req.LandID)), req.Location)
    if err != nil {
        log.Fatalf("Failed to call RegisterLand: %v", err)
    }

    log.Printf("Transaction hash: %s", tx.Hash().Hex())
    return &types.RegisterLandResp{
        Code: 0,
        Msg:  "Land registered successfully",
    }, nil
    
}
