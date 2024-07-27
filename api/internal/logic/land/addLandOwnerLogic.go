package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLandOwnerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLandOwnerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLandOwnerLogic {
	return &AddLandOwnerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLandOwnerLogic) AddLandOwner(req *types.AddLandOwnerReq) (resp *types.AddLandOwnerResp, err error) {
	// todo: add your logic here and delete this line
	return
}
