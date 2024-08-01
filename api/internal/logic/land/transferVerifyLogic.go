package land

import (
	"context"
	"fmt"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferVerifyLogic {
	return &TransferVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferVerifyLogic) TransferVerify(
	req *types.TransferVerifyReq) (resp *types.TransferVerifyResp, err error) {

	auth := l.svcCtx.AccountAuth.GetAccountAuth(l.svcCtx.OracleKey)
	tx, err := l.svcCtx.Conn.TransferVerify(
		auth, req.LandId, common.HexToAddress(req.Address_from), common.HexToAddress(req.Address_to), big.NewInt(int64(int(req.TransactionId))))
	if err != nil {
		return nil, fmt.Errorf("failed to call landSurveyingArea: %v", err)
	}
	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	return &types.TransferVerifyResp{
		Code: 0,
		Msg:  "Success",
	}, nil
}
