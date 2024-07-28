package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateProofOfOwnershipLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGenerateProofOfOwnershipLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateProofOfOwnershipLogic {
	return &GenerateProofOfOwnershipLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateProofOfOwnershipLogic) GenerateProofOfOwnership(
	req *types.GenerateProofOfOwnershipReq) (resp *types.GenerateProofOfOwnershipResp, err error) {

	return
}
