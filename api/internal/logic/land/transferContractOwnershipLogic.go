package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

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

func (l *TransferContractOwnershipLogic) TransferContractOwnership(req *types.TransferContractOwnershipReq) (resp *types.TransferContractOwnershipResp, err error) {
	// todo: add your logic here and delete this line

	return
}
