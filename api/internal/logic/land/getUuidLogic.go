package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUuidLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUuidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUuidLogic {
	return &GetUuidLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUuidLogic) GetUuid() (resp *types.GetUuidResp, err error) {
	// todo: add your logic here and delete this line

	return
}
