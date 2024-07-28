package land

import (
	"context"

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

func (l *TransferOwnershipLogic) TransferOwnership(req *types.TransferOwnershipReq) (resp *types.TransferOwnershipResp, err error) {
	// todo: add your logic here and delete this line
	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.NotaryAddress)
    
    // 调用智能合约的 transferOwnership 方法并传入授权信息和交易信息
    tx, err := l.svcCtx.Conn.TransferOwnership(auth, big.NewInt(int64(req.LandID)), common.HexToAddress(req.ToAddress), big.NewInt(int64(req.TransactionID)))
    if err != nil {
        log.Fatalf("Failed to call TransferOwnership: %v", err)
    }

    log.Printf("Transaction hash: %s", tx.Hash().Hex())
    return &types.TransferOwnershipResp{
        Code: 0,
        Msg:  "Ownership transferred successfully",
    }, nil
}
