package land

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferOwnershipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferOwnershipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferOwnershipLogic {
	return &TransferOwnershipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferOwnershipLogic) TransferOwnership(
	req *types.TransferOwnershipReq) (resp *types.TransferOwnershipResp, err error) {

	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.Senderkey)
	tx, err := l.svcCtx.Conn.TransferOwnership(
		auth, req.LandId, common.HexToAddress(req.Address_to))
	if err != nil {
		return nil, fmt.Errorf("failed to call transfer ownership: %v", err)
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	return &types.TransferOwnershipResp{
		Code: 0,
		Msg:  "😥😥😥",
	}, nil
}
