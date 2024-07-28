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

	return
}
