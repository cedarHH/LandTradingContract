package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNotaryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddNotaryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNotaryLogic {
	return &AddNotaryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddNotaryLogic) AddNotary(req *types.AddNotaryReq) (resp *types.AddNotaryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
