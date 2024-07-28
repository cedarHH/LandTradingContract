package land

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLandLogic {
	return &QueryLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLandLogic) QueryLand(req *types.QueryLandReq) (
	resp *types.QueryLandResp, err error) {

	// auth := l.svcCtx.AccountAuth.GetAccountAuth(req.Senderkey)
	land, err := l.svcCtx.Conn.QueryLand(&bind.CallOpts{}, req.LandId)
	if err != nil {
		return nil, fmt.Errorf("failed to queryLand: %v", err)
	}

	landEntry, err := l.svcCtx.LandModel.FindOne(l.ctx, req.LandId)
	if err != nil {
		return nil, fmt.Errorf("failed to find land entry: %v", err)
	}

	detail, err := l.svcCtx.LandBucket.GetPresignedDownloadURL(l.ctx, landEntry.Detail, 3600)
	if err != nil {
		return nil, fmt.Errorf("s3 error: %v", err)
	}

	report, err := l.svcCtx.LandBucket.GetPresignedDownloadURL(l.ctx, landEntry.Detail, 3600)
	if err != nil {
		return nil, fmt.Errorf("s3 error: %v", err)
	}

	document, err := l.svcCtx.LandBucket.GetPresignedDownloadURL(l.ctx, landEntry.Detail, 3600)
	if err != nil {
		return nil, fmt.Errorf("s3 error: %v", err)
	}

	return &types.QueryLandResp{
		Code: 0,
		Data: types.LandDetail{
			Owner:      land.Owner.String(),
			Location:   land.Location,
			Area:       land.Area.Int64(),
			IsVerified: land.IsVerified,
			Details:    detail,
			Report:     report,
			Document:   document,
		},
		Msg: "",
	}, nil
}
