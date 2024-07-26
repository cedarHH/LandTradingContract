package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyProofOfOwnershipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyProofOfOwnershipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyProofOfOwnershipLogic {
	return &VerifyProofOfOwnershipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyProofOfOwnershipLogic) VerifyProofOfOwnership(req *types.VerifyProofOfOwnershipReq) (resp *types.VerifyProofOfOwnershipResp, err error) {
	// todo: add your logic here and delete this line

	return
}
