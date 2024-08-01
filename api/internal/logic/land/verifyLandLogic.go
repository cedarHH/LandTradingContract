package land

import (
	"context"
	"fmt"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/tool"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type VerifyLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerifyLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerifyLandLogic {
	return &VerifyLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerifyLandLogic) VerifyLand(
	req *types.VerifyLandReq) (resp *types.VerifyLandResp, err error) {

	download, err := l.svcCtx.LandBucket.Download(l.ctx, req.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %v", err)
	}

	fileHash, err := tool.ComputeFileHash(download)

	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.Senderkey)
	tx, err := l.svcCtx.Conn.VerifyLand(auth, req.LandId, fileHash, req.IsVerify)
	if err != nil {
		return nil, fmt.Errorf("failed to call landSurveyingArea: %v", err)
	}
	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	updates := map[string]interface{}{}
	updates["document"] = req.Document
	err = l.svcCtx.LandModel.UpdateAttributes(l.ctx, req.LandId, updates)
	if err != nil {
		return nil, fmt.Errorf("failed to update attributes: %v", err)
	}

	return &types.VerifyLandResp{
		Code: 0,
		Msg:  "Success",
	}, nil
}
