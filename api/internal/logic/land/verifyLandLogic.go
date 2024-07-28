package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLandLogic {
	return &VerifyLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLandLogic) VerifyLand(req *types.VerifyLandReq) (resp *types.VerifyLandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
