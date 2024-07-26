package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLandLogic {
	return &RegisterLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLandLogic) RegisterLand(req *types.RegisterLandReq) (resp *types.RegisterLandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
