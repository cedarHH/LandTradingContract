package land

import (
	"context"
	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/tool"
	"github.com/cedarHH/LandTradingContract/api/internal/types"
	"github.com/cedarHH/LandTradingContract/model/database"
	"log"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLandLogic {
	return &RegisterLandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLandLogic) RegisterLand(
	req *types.RegisterLandReq) (resp *types.RegisterLandResp, err error) {

	download, err := l.svcCtx.LandBucket.Download(l.ctx, req.Details)
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}

	fileHash, err := tool.ComputeFileHash(download)
	if err != nil {
		log.Fatalf("Failed to compute file hash: %v", err)
	}

	auth := l.svcCtx.AccountAuth.GetAccountAuth(req.Senderkey)
	tx, err := l.svcCtx.Conn.RegisterLand(
		auth, req.LandId, req.Location, fileHash)
	if err != nil {
		log.Fatalf("Failed to call register land: %v", err)
	}

	log.Printf("Transaction hash: %s", tx.Hash().Hex())

	land := &database.Land{
		LandId:   req.LandId,
		Detail:   req.Details,
		Report:   "",
		Document: "",
	}
	err = l.svcCtx.LandModel.Insert(l.ctx, land)
	if err != nil {
		log.Fatalf("Failed to create land entry: %v", err)
	}

	return &types.RegisterLandResp{
		Code: 0,
		Msg:  "😥😥😥",
	}, nil
}
