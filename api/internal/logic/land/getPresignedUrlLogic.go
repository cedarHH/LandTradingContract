package land

import (
	"context"
	"log"

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

func (l *GetPresignedUrlLogic) GetPresignedUrl(
	req *types.GetPresignedUrlReq) (resp *types.GetPresignedUrlResp, err error) {

	url, err := l.svcCtx.LandBucket.GetPresignedUploadURL(l.ctx, req.FileName, 3600)
	if err != nil {
		log.Fatalf("Failed to get presigned url: %v", err)
	}

	return &types.GetPresignedUrlResp{
		Code: 0,
		Url:  url,
		Msg:  "🤮🤮🤮🤮",
	}, nil
}
