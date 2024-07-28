package land

import (
	"context"

	"github.com/cedarHH/LandTradingContract/api/internal/svc"
	"github.com/cedarHH/LandTradingContract/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSurveyorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSurveyorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSurveyorLogic {
	return &AddSurveyorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSurveyorLogic) AddSurveyor(req *types.AddSurveyorReq) (resp *types.AddSurveyorResp, err error) {
	// todo: add your logic here and delete this line

	return
}
