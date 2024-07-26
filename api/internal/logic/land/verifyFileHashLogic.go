package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyFileHashLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyFileHashLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyFileHashLogic {
	return &VerifyFileHashLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyFileHashLogic) VerifyFileHash(req *types.VerifyFileHashReq) (resp *types.VerifyFileHashResp, err error) {
	// todo: add your logic here and delete this line

	return
}
