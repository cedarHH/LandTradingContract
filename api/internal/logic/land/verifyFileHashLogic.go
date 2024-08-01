package land

import (
	"context"
	"fmt"
	"github.com/cedarHH/LandTradingContract/api/internal/tool"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"

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

func (l *VerifyFileHashLogic) VerifyFileHash(
	req *types.VerifyFileHashReq) (resp *types.VerifyFileHashResp, err error) {

	land, err := l.svcCtx.Conn.QueryLand(&bind.CallOpts{}, req.LandId)
	if err != nil {
		return nil, fmt.Errorf("failed to queryLand: %v", err)
	}

	landEntry, err := l.svcCtx.LandModel.FindOne(l.ctx, req.LandId)
	if err != nil {
		log.Fatalf("Failed to find land entry: %v", err)
	}

	download, err := l.svcCtx.LandBucket.Download(l.ctx, landEntry.Detail)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	detail, err := tool.VerifyFileIntegrity(download, land.DetailsHash)
	if err != nil {
	}

	download, err = l.svcCtx.LandBucket.Download(l.ctx, landEntry.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	report, err := tool.VerifyFileIntegrity(download, land.ReportHash)
	if err != nil {
	}

	download, err = l.svcCtx.LandBucket.Download(l.ctx, landEntry.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}
	document, err := tool.VerifyFileIntegrity(download, land.DocumentsHash)
	if err != nil {
	}

	return &types.VerifyFileHashResp{
		Code: 0,
		Data: types.VerifyData{
			IsTampered: !(report && detail && document),
		},
		Msg: "",
	}, nil
}
