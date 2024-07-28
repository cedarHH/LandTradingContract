package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPresignedUrlLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPresignedUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPresignedUrlLogic {
	return &GetPresignedUrlLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPresignedUrlLogic) GetPresignedUrl(req *types.GetPresignedUrlReq) (resp *types.GetPresignedUrlResp, err error) {
	// todo: add your logic here and delete this line

	return
}
