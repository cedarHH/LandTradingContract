package land

import (
	"context"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/ethereum/go-ethereum/common"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type TransferContractOwnershipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTransferContractOwnershipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TransferContractOwnershipLogic {
	return &TransferContractOwnershipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TransferContractOwnershipLogic) TransferContractOwnership(
	req *types.TransferContractOwnershipReq) (resp *types.TransferContractOwnershipResp, err error) {

	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.Senderkey)
	newOwner := common.HexToAddress(req.NewOwner)
	tx, err := l.svcCtx.Conn.TransferOwner(auth, newOwner)
	if err != nil {
		log.Fatalf("Failed to call add notary: %v", err)
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())
	return &types.TransferContractOwnershipResp{
		Code: 0,
		Msg:  "😋😋😋",
	}, nil
}
