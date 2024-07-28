package land

import (
	"context"
	"fmt"
	"github.com/cedarHH/LandTradingContract/api/internal/tool"
	"log"
	"math/big"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SurveyLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSurveyLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SurveyLandLogic {
	return &SurveyLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SurveyLandLogic) SurveyLand(
	req *types.SurveyLandReq) (resp *types.SurveyLandResp, err error) {

	download, err := l.svcCtx.LandBucket.Download(l.ctx, req.Report)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}

	fileHash, err := tool.ComputeFileHash(download)

	auth := l.svcCtx.AccountAuth.GetAccountAuth(l.svcCtx.OracleKey)
	tx, err := l.svcCtx.Conn.LandSurveyingArea(auth, req.LandId, big.NewInt(req.Area), fileHash)
	if err != nil {
		return nil, fmt.Errorf("failed to call landSurveyingArea: %v", err)
	}
	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	return &types.SurveyLandResp{
		Code: 0,
		Msg:  "Success",
	}, nil
}
