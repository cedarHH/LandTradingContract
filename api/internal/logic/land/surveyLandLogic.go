package land

import (
	"context"

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

func (l *SurveyLandLogic) SurveyLand(req *types.SurveyLandReq) (resp *types.SurveyLandResp, err error) {
	// todo: add your logic here and delete this line

	return
}
